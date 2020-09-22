// (C) Copyright [2020] Anand S - <Anand.S.Main@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package datacommunicator ...
package datacommunicator

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/go-amqp"
)

// QpidPacket defines the Qpid Message Packet
// HandleConn - AMQP Connection Object
type QpidPacket struct {

	// Qpid Server Connection Object
	HandleConn *amqp.Client
	// Qpid Server name and Port holder
	Server string
}

// Connect - Defines the procedure to connect to Qpid server.
func (qp *QpidPacket) Connect() error {

	// In case where HandleConn is already created, we don't recreate the connection again.
	if qp.HandleConn == nil {
		conn, e := amqp.Dial("")
		if e != nil {
			er := fmt.Errorf("AMQP Connect Error: %#v, %s", e, "")
			return er
		}
		qp.HandleConn = conn
	}
	return nil
}

// Distribute would send a user defined message using Qpid server
func (qp *QpidPacket) Distribute(pipe string, d interface{}) error {
	var e error
	session, e := qp.HandleConn.NewSession()
	if e != nil {
		er := fmt.Errorf("AMQP Dispatch Session Error - %#v", e)
		return er
	}
	sender, e := session.NewSender(amqp.LinkTargetAddress(pipe))
	if e != nil {
		er := fmt.Errorf("Creating Sender Link - FAILED : %#v", e)
		return er
	}
	// Message Encoding done
	b, e := Encode(d)
	if e != nil {
		log.Printf("%v", e)
		return e
	}
	Ctx := context.Background()
	defer sender.Close(Ctx)
	if e = sender.Send(Ctx, amqp.NewMessage(b)); e != nil {
		er := fmt.Errorf("Message dispatch Error - %#v", e)
		return er
	}
	return nil
}

// Accept defines the Subscription / Consume procedure. The Message processing
// will be done in separate Go routine.
func (qp *QpidPacket) Accept(pipe string, fn MsgProcess) error {
	session, e := qp.HandleConn.NewSession()
	if e != nil {
		er := fmt.Errorf("AMQP Accept Session Error - %#v", e)
		return er
	}
	receiver, e := session.NewReceiver(
		amqp.LinkSourceAddress(pipe),
		amqp.LinkCredit(10),
	)
	if e != nil {
		er := fmt.Errorf("Creating Receiver Link - FAILED : %#v", e)
		return er
	}
	go qp.read(receiver, fn)
	return nil
}

// read - get the message using receiver call.
func (qp *QpidPacket) read(r *amqp.Receiver, fn MsgProcess) error {
	ctx := context.Background()
	var d interface{}

	defer r.Close(ctx)
	for {
		m, e := r.Receive(ctx)
		if e != nil {
			er := fmt.Errorf("Message Receive Error - %#v", e)
			return er
		}
		m.Accept(ctx)
		if e = Decode(m.GetData(), &d); e != nil {
			return e
		}
		fn(d)
	}
}

// Remove will close a specific session from the local stored map
func (qp *QpidPacket) Remove(pipe string) error {
	return nil
}

// Close will close AMQP connection. Repeated calls would result in
// unexpected error.
func (qp *QpidPacket) Close() {
	qp.HandleConn.Close()
}
