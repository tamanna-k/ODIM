//(C) Copyright [2020] Anand S <Anand.S.Main@gmail.com>
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
package datacommunicator

import (
	"reflect"
	"testing"

	"github.com/Azure/go-amqp"
)

func TestQpidConnect(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := reflect.New(MMap[QPID]).Interface().(MQBus)
			if err := qp.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() Qpid - error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQpidPacket_Distribute(t *testing.T) {
	type fields struct {
		d *amqp.Client
		s string
	}
	type args struct {
		pipe string
		d    interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QpidPacket{
				HandleConn: tt.fields.d,
				Server:     tt.fields.s,
			}
			qp.Distribute(tt.args.pipe, tt.args.d)
		})
	}
}

func TestQpidPacket_Accept(t *testing.T) {
	type fields struct {
		d *amqp.Client
		s string
	}
	type args struct {
		pipe string
		fn   MsgProcess
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QpidPacket{
				HandleConn: tt.fields.d,
				Server:     tt.fields.s,
			}
			if err := qp.Accept(tt.args.pipe, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Accept() Qpid - error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQpidPacket_Read(t *testing.T) {
	type fields struct {
		d *amqp.Client
		s string
	}
	type args struct {
		p  string
		fn MsgProcess
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QpidPacket{
				HandleConn: tt.fields.d,
				Server:     tt.fields.s,
			}
			if err := qp.read(tt.args.p, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("KafkaPacket.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestQpidPacket_Close(t *testing.T) {
	type fields struct {
		d *amqp.Client
		s string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QpidPacket{
				DialerConn: tt.fields.d,
				Server:     tt.fields.s,
			}
			qp.Close()
		})
	}
}
