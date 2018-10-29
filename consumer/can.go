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
	"fmt"
	"sync"

	"github.com/brutella/can"
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
//  Can0In:
//    Type: consumer.Can
//    Streams: can0
//    Interface: can0
type Can struct {
	core.SimpleConsumer `gollumdoc:"embed_type"`
	bus                 *can.Bus
	iface               string `config:"Interface"`
	hasToSetMetadata    bool   `config:"SetMetadata" default:"false"`
}

func init() {
	core.TypeRegistry.Register(Can{})
}

// Configure initializes this consumer with values from a plugin config.
func (cons *Can) Configure(conf core.PluginConfigReader) {
	bus, err := can.NewBusForInterfaceWithName(cons.iface)
	if err != nil {
		cons.Logger.Error(err)
	}
	cons.bus = bus
}

// Enqueue creates a new message
func (cons *Can) Enqueue(data []byte) {
	if cons.hasToSetMetadata {
		metaData := core.Metadata{}
		metaData.SetValue("iface", []byte(cons.iface))

		cons.EnqueueWithMetadata(data, metaData)
	} else {
		cons.SimpleConsumer.Enqueue(data)
	}
}

func (cons *Can) ProcessFrame(frm can.Frame) {
	str := fmt.Sprintf("%-4x [%d] % -24X", frm.ID, frm.Length, frm.Data[:frm.Length])
	cons.Enqueue([]byte(str))
}

// Consume listens to stdin.
func (cons *Can) Consume(workers *sync.WaitGroup) {
	// Attach frame processing func to this bus
	cons.bus.SubscribeFunc(cons.ProcessFrame)

	// Kick off publish function
	go cons.bus.ConnectAndPublish()

	// Kill the connection after we have exit
	defer cons.bus.Disconnect()

	// Wait for an exit signal
	cons.ControlLoop()
}
