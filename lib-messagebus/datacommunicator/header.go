//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

// Package datacommunicator ...
package datacommunicator

// -----------------------------------------------------------------------------
// IMPORT Section
// -----------------------------------------------------------------------------
import (
	"encoding/json"
	"log"
	"reflect"
)

// BrokerID defines the Backend Communication Broker / Router ID
type BrokerID int

// KAFKA, QPID possible Routers, that can be used for communication
const (
	KAFKA BrokerID = iota
	QPID
)

// MMap - defines Communication Medium map
var MMap = make(map[BrokerID]reflect.Type)

func init() {
	MMap[KAFKA] = reflect.TypeOf((*KafkaPacket)(nil)).Elem()
	MMap[QPID] = reflect.TypeOf((*QpidPacket)(nil)).Elem()
}

// MQBus Interface defines the Process interface function These functions
// are implemented as part of Packet struct.
// Connect - API to connect to specified Router / Broker / Medium
// Distribute - API to Publish Messages into specified Pipe (Topic / Subject)
// Accept - Consume the incoming message if subscribed by that component
// Get - Would initiate blocking call to remote process to get response
// Close - Would disconnect the connection with Middleware.
type MQBus interface {
	Connect() error
	Distribute(pipe string, data interface{}) error
	Accept(pipe string, fn MsgProcess) error
	Remove(pipe string) error
	Close()
}

// MsgProcess defines the functions for processing accepted messages. Any client
// who wants to accept and handle the events / notifications / messages, should
// implement this function as part of their procedure. That same function should
// be sent to MessageBus as callback for handling the incoming messages.
type MsgProcess func(d interface{})

// Communicator defines the Broker platform Middleware selection and corresponding
// communication object would be created to send / receive the messages. Broker
// type would be stored as part of Connection Object "Packet".
func Communicator(bt BrokerID) (MQBus, error) {

	p := reflect.New(MMap[bt]).Interface().(MQBus)
	if e := p.Connect(); e != nil {
		return nil, e
	}
	return p, nil
}

// Encode converts the interface into Byte stream (ENCODE).
func Encode(d interface{}) ([]byte, error) {

	data, err := json.Marshal(d)
	if err != nil {
		log.Println("error: Failed to encode the given event data: ", err)
		return nil, err
	}
	return data, nil
}

// Decode converts the byte stream into Data (DECODE).
///data will  be masked as Interface before sent to Consumer or Requester.
func Decode(d []byte, a interface{}) error {
	err := json.Unmarshal(d, &a)
	if err != nil {
		log.Println("error: Failed to decode the event data: ", err)
		return err
	}
	return nil
}
