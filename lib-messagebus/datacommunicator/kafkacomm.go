// (C) Copyright [2020] Hewlett Packard Enterprise Development LP
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

// -----------------------------------------------------------------------------
// IMPORT Section
// -----------------------------------------------------------------------------
import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaPacket would define the Kafka specific stream object maps and Connection
// handler. Mapping between streams and Pipe would allow user to create or close
// Pipe specific connections without winding down Kafka connection handle.
type KafkaPacket struct {

	// Readers - Dict between Pipename and Kafka Reader streams
	Readers map[string]*kafka.Reader

	// Writers - Dict between Kafka writers and Pipename passed to Connect
	Writers map[string]*kafka.Writer

	// DialerConn - Single Connection Handler towards Kafka
	DialerConn *kafka.Dialer

	// Server defines the KAFKA server with port
	Server string
}

// TLS creates the TLS Configuration object to used by any Broker for Auth and
// Encryption. The Certficate and Key files are created from Java Keytool
// generated JKS format files. Please look into README for more information
// In case of Kafka, we generate the Server certificate files in JKS format.
// We do the same for Clients as well. Then we convert those files into PEM
// format.
func TLS(cCert, cKey, caCert string) (*tls.Config, error) {

	tlsConfig := tls.Config{}

	// Load client cert
	cert, e1 := tls.LoadX509KeyPair(cCert, cKey)
	if e1 != nil {
		return &tlsConfig, e1
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCertR, e2 := ioutil.ReadFile(caCert)
	if e2 != nil {
		return &tlsConfig, e2
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCertR)
	tlsConfig.RootCAs = caCertPool

	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, e2
}

// Connect - Initiates the connection towards Kafka with the specified TLS
// object. This would initialize the Reader and Writer stream for handling data.
func (kp *KafkaPacket) Connect() error {

	// Using MQF details, connecting to the KAFKA Server.
	kp.Server = mq.KServer + ":" + strconv.Itoa(mq.KLport)

	// Creation of TLS Config and Dialer
	tls, e := TLS(mq.KAFKACertFile, mq.KAFKAKeyFile, mq.KAFKACAFile)
	if e != nil {
		log.Printf("%v", e)
		return e
	}

	kp.DialerConn = &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       tls,
	}

	// Initialize the connection map for both Reader and Writer for KAFKA
	if kp.Readers == nil {
		kp.Readers = make(map[string]*kafka.Reader)
	}
	if kp.Writers == nil {
		kp.Writers = make(map[string]*kafka.Writer)
	}

	return nil
}

// Distribute - would define the message publishing task. If the Writer stream
// for this specific stream not available, same would be created.
func (kp *KafkaPacket) Distribute(pipe string, d interface{}) error {

	// Check for existing Writers. If not existing for this specific Pipe,
	// then we would create this Writer object for sending the message.
	if _, a := kp.Writers[pipe]; a == false {

		kp.Writers[pipe] = kafka.NewWriter(kafka.WriterConfig{
			Brokers:       []string{kp.Server},
			Topic:         pipe,
			Balancer:      &kafka.LeastBytes{},
			BatchSize:     1,
			QueueCapacity: 1,
			Async:         true,
			Dialer:        kp.DialerConn,
		})
	}
	// Encode the message before appending into KAFKA Message struct
	b, e := Encode(d)
	if e != nil {
		log.Printf("%v", e)
		return e
	}
	// Place the byte stream into Kafka.Message
	km := kafka.Message{
		Key:   []byte(pipe),
		Value: b,
	}
	// Write the messgae in the specified Pipe.
	if e = kp.Writers[pipe].WriteMessages(context.Background(), km); e != nil {
		log.Printf("%v", e)
		return e
	}
	return nil
}

// Accept function defines the Consumer or Subscriber functionality for KAFKA.
// If Reader object for the specified Pipe is not available, New Reader Object
// would be created.
func (kp *KafkaPacket) Accept(pipe string, fn MsgProcess) error {

	// If for the Reader Object for pipe and create one if required.
	if _, a := kp.Readers[pipe]; a == false {

		kp.Readers[pipe] = kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{kp.Server},
			GroupID:        pipe,
			Topic:          pipe,
			MinBytes:       10e1,
			MaxBytes:       10e6,
			CommitInterval: 1 * time.Second,
			Dialer:         kp.DialerConn,
		})
	}
	kp.read(pipe, fn)
	return nil
}

// read would access the KAFKA messages in a infinite loop. Callback method
// access is existing only in "goka" library.  Not available in "kafka-go".
func (kp *KafkaPacket) read(p string, fn MsgProcess) error {

	var d interface{}
	c := context.Background()
	// Infinite loop to constantly reading the message from KAFKA.
	for {
		m, e := kp.Readers[p].ReadMessage(c)
		if e != nil {
			log.Printf("%v", e)
			return e
		}
		if e = Decode(m.Value, &d); e != nil {
			return e
		}
		// Callback Function call.
		fn(d)
	}
}

// Remove will just remove the existing subscription. This API would check just
// the Reader map as to Distribute / Publish messages, we don't need subscription
func (kp *KafkaPacket) Remove(pipe string) error {

	es, ok := kp.Readers[pipe]
	if ok == false {
		e := fmt.Errorf("specified pipe is not subscribed yet. please check the pipe name passed")
		return e
	}
	es.Close()
	delete(kp.Readers, pipe)

	return nil
}

// Close will disconnect KAFKA Connection.
func (kp *KafkaPacket) Close() {

	// Closing all opened Readers Connections
	for rp, rc := range kp.Readers {
		rc.Close()
		delete(kp.Readers, rp)
	}
	// Closing all opened Writers Connections
	for wp, wc := range kp.Writers {
		wc.Close()
		delete(kp.Writers, wp)
	}
}
