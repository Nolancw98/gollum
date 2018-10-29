// Copyright 2018 John Ott
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consumer

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/trivago/gollum/core"
)

// GPSD consumer
//
// This consumer reads from the gpsd service.
//
// Parameters
//
// - Address: Defines the adddress to dial. By default this is set to "localhost:2947"
//
// Examples
//
// This config reads data from can0
//
//  GpsdIn:
// 	  Type: consumer.Gpsd
//    Streams: gpsd
//    Address: localhost:2947
type Gpsd struct {
	core.SimpleConsumer `gollumdoc:"embed_type"`
	address             string `config:"Address" default:"localhost:2947"`
	socket net.Conn
}

func init() {
	core.TypeRegistry.Register(Gpsd{})
}

// Configure initializes this consumer with values from a plugin config.
func (cons *Gpsd) Configure(conf core.PluginConfigReader) {
	sock, err := net.Dial("tcp4", cons.address)
	if err != nil {
		cons.Logger.Error(err)
	}
	cons.socket = sock
}

// Consume listens to stdin.
func (cons *Gpsd) Consume(workers *sync.WaitGroup) {
	reader := bufio.NewReader(cons.socket)

	// Send watch command
	fmt.Fprintf(cons.socket, "?WATCH={\"enable\":true,\"json\":true}")


	go func(){
		// Close the socket on exit
		defer cons.socket.Close()

		// Run until we're done
		for cons.IsActive() {
			switch line, err := reader.ReadString('\n'); err {
			case nil:
				cons.Enqueue([]byte(line))
			case io.EOF:
				return
			default:
				cons.Logger.Error(err)
				return
			}
		}
	}()

	// Wait for an exit signal
	cons.ControlLoop()
}
