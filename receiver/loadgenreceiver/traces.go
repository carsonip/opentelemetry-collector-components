// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package loadgenreceiver

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"math/rand"
	"os"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

//go:embed testdata/traces.jsonl
var demoTraces []byte

type receiverTraces struct {
	traces   ptrace.Traces
	jsonSize int
}

type tracesGenerator struct {
	cfg    *Config
	logger *zap.Logger

	sampleTraces    []receiverTraces
	lastSampleIndex int
	consumer        consumer.Traces

	cancelFn context.CancelFunc
}

func createTracesReceiver(
	ctx context.Context,
	set receiver.Settings,
	config component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	genConfig := config.(*Config)
	recv := tracesGenerator{
		cfg:             genConfig,
		logger:          set.Logger,
		consumer:        consumer,
		sampleTraces:    make([]receiverTraces, 0),
		lastSampleIndex: 0,
	}

	parser := ptrace.JSONUnmarshaler{}
	var err error
	sampleTraces := demoTraces

	if genConfig.Traces.JsonFile != "" {
		sampleTraces, err = os.ReadFile(string(genConfig.Traces.JsonFile))
		if err != nil {
			return nil, err
		}
	}

	scanner := bufio.NewScanner(bytes.NewReader(sampleTraces))
	for scanner.Scan() {
		traceBytes := scanner.Bytes()
		lineTraces, err := parser.UnmarshalTraces(traceBytes)
		if err != nil {
			return nil, err
		}
		recv.sampleTraces = append(recv.sampleTraces, receiverTraces{
			traces:   lineTraces,
			jsonSize: len(traceBytes),
		})
	}

	return &recv, nil
}

func (ar *tracesGenerator) Start(ctx context.Context, _ component.Host) error {
	startCtx, cancelFn := context.WithCancel(ctx)
	ar.cancelFn = cancelFn

	go func() {
		for {
			select {
			case <-startCtx.Done():
				return
			default:
				nTraces, _, err := ar.nextTraces()
				if err != nil {
					ar.logger.Error(err.Error())
					continue
				}
				err = ar.consumer.ConsumeTraces(startCtx, nTraces)
				if err != nil {
					ar.logger.Error(err.Error())
					continue
				}
			}
		}
	}()
	return nil
}

func (ar *tracesGenerator) Shutdown(context.Context) error {
	if ar.cancelFn != nil {
		ar.cancelFn()
	}
	return nil
}

func (ar *tracesGenerator) nextTraces() (ptrace.Traces, int, error) {
	nextLogs := ptrace.NewTraces()

	ar.sampleTraces[ar.lastSampleIndex].traces.CopyTo(nextLogs)
	sampledSize := ar.sampleTraces[ar.lastSampleIndex].jsonSize

	rm := nextLogs.ResourceSpans()
	for i := 0; i < rm.Len(); i++ {
		for j := 0; j < rm.At(i).ScopeSpans().Len(); j++ {
			for k := 0; k < rm.At(i).ScopeSpans().At(j).Spans().Len(); k++ {
				sspan := rm.At(i).ScopeSpans().At(j).Spans().At(k)
				now := time.Now()
				// Generate a random duration between 0 and 3 seconds
				sspan.SetStartTimestamp(pcommon.NewTimestampFromTime(now.Add(-time.Duration(rand.Int63n(int64(3 * time.Second))))))
				sspan.SetEndTimestamp(pcommon.NewTimestampFromTime(now))

			}
		}
	}

	ar.lastSampleIndex = (ar.lastSampleIndex + 1) % len(ar.sampleTraces)

	return nextLogs, sampledSize, nil
}

func (hmr *tracesGenerator) shutdown(_ context.Context) error {
	return nil
}