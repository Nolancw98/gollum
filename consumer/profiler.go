// Copyright 2015 trivago GmbH
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
	"github.com/trivago/gollum/core"
	"github.com/trivago/gollum/core/log"
	"github.com/trivago/gollum/shared"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Profiler consumer plugin
// Configuration example
//
//   - "consumer.Console":
//     Enable: true
//
// This consumer does not define any options beside the standard ones.
type Profiler struct {
	core.ConsumerBase
	profileRuns int
	batches     int
	length      int
	quit        bool
}

func init() {
	shared.RuntimeType.Register(Profiler{})
}

// Configure initializes this consumer with values from a plugin config.
func (cons *Profiler) Configure(conf core.PluginConfig) error {
	err := cons.ConsumerBase.Configure(conf)
	if err != nil {
		return err
	}

	cons.profileRuns = conf.GetInt("Runs", 10000)
	cons.batches = conf.GetInt("Batches", 10)
	cons.length = conf.GetInt("Length", 256)

	return nil
}

var stringBase = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890 ")

func (cons *Profiler) profile() {

	randString := make([]byte, cons.length)
	for i := 0; i < cons.length; i++ {
		randString[i] = stringBase[rand.Intn(len(stringBase))]
	}

	testStart := time.Now()
	minTime := math.MaxFloat64
	maxTime := 0.0
	buffer := shared.NewByteStream(cons.length + 64)

	for b := 0; b < cons.batches; b++ {
		Log.Note.Print(fmt.Sprintf("run %d/%d:", b, cons.batches))

		start := time.Now()
		for i := 0; i < cons.profileRuns; i++ {
			buffer.Reset()
			fmt.Fprintf(&buffer, "%d ", i)
			buffer.Write(randString)

			cons.PostCopy(buffer.Bytes(), uint64(b*cons.profileRuns+i))
			if cons.quit {
				return
			}
		}

		runTime := time.Since(start)
		minTime = math.Min(minTime, runTime.Seconds())
		maxTime = math.Max(maxTime, runTime.Seconds())
	}

	runTime := time.Since(testStart)

	Log.Note.Print(fmt.Sprintf(
		"Avg: %.4f sec = %4.f msg/sec",
		runTime.Seconds(),
		float64(cons.profileRuns*cons.batches)/runTime.Seconds()))

	Log.Note.Print(fmt.Sprintf(
		"Best: %.4f sec = %4.f msg/sec",
		minTime,
		float64(cons.profileRuns)/minTime))

	Log.Note.Print(fmt.Sprintf(
		"Worst: %.4f sec = %4.f msg/sec",
		maxTime,
		float64(cons.profileRuns)/maxTime))

	proc, _ := os.FindProcess(os.Getpid())
	proc.Signal(os.Interrupt)
}

// Consume starts a profile run and exits gollum when done
func (cons *Profiler) Consume(workers *sync.WaitGroup) {
	cons.quit = false
	go func() {
		defer shared.RecoverShutdown()
		cons.profile()
	}()

	defer func() {
		cons.quit = true
	}()

	cons.DefaultControlLoop(nil)
}
