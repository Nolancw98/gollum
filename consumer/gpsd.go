// Copyright 2019 John Ott
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
	"net"
	"sync"

	"github.com/trivago/gollum/core"
)

// CAN bus consumer
//
// This consumer reads from a CAN interface. A message is generated for each
// message
//
// Metadata
//
// *NOTE: The metadata will only set if the parameter `SetMetadata` is active.*
//
// - iface: Name of the interface the message was received on (set)
//
// Parameters
//
// - Interface: Defines the interface to read from. This parameter is required.
//
// - SetMetadata: When this value is set to "true", the fields mentioned in the metadata
// section will be added to each message. Adding metadata will have a
// performance impact on systems with high throughput.
// By default this parameter is set to "false".
//
// Examples
//
// This config reads data from can0
//
//  GpsdIn:
//    Type: consumer.Gpsd
//    Streams: gpsd
type Gpsd struct {
	core.SimpleConsumer `gollumdoc:"embed_type"`
	address             string `config:"Address" default:"localhost:2947"`
	socket              net.Conn
	reader              *bufio.Reader
}

func init() {
	core.TypeRegistry.Register(Gpsd{})
}

// Configure initializes this consumer with values from a plugin config.
func (cons *Gpsd) Configure(conf core.PluginConfigReader) {
	socket, err := net.Dial("tcp4", cons.address)
	if err != nil {
		cons.Logger.Error(err)
	}
	cons.socket = socket
	cons.reader = bufio.NewReader(cons.socket)
	cons.reader.ReadString('\n')
}

func (cons *Gpsd) readFromSocket() {
	defer cons.socket.Close()

	// Send watch command
	fmt.Fprintf(cons.socket, "?WATCH={\"enable\":true,\"json\":true}")

	// Read while active
	for cons.IsActive() {
		if line, err := cons.reader.ReadString('\n'); err == nil {
			cons.Enqueue([]byte(line))
		} else {
			cons.Logger.Error(err)
		}
	}
}

// Consume listens to stdin.
func (cons *Gpsd) Consume(workers *sync.WaitGroup) {
	go cons.readFromSocket()
	cons.ControlLoop()
}
