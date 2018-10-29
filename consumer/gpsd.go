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

	"github.com/stratoberry/go-gpsd"
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
	gps                 *gpsd.Session
	address             string `config:"Address" default:"localhost:2947"`
}

func init() {
	core.TypeRegistry.Register(Gpsd{})
}

// Configure initializes this consumer with values from a plugin config.
func (cons *Gpsd) Configure(conf core.PluginConfigReader) {
	gps, err := gpsd.Dial(gpsd.DefaultAddress)
	if err != nil {
		cons.Logger.Error(err)
	}
	cons.gps = gps
}

// Consume listens to stdin.
func (cons *Gpsd) Consume(workers *sync.WaitGroup) {
	// Attach frame processing func to this bus
	cons.bus.SubscribeFunc(cons.ProcessFrame)

	// Kick off publish function
	go cons.bus.ConnectAndPublish()

	// Kill the connection after we have exit
	defer cons.bus.Disconnect()

	// Wait for an exit signal
	cons.ControlLoop()
}
