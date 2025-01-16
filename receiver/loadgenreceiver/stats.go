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

package loadgenreceiver // import "github.com/elastic/opentelemetry-collector-components/receiver/loadgenreceiver"

type TelemetryStats struct {
	Requests       int
	FailedRequests int

	LogRecords       int
	MetricDataPoints int
	Spans            int

	FailedLogRecords       int
	FailedMetricDataPoints int
	FailedSpans            int
}

func (s TelemetryStats) Add(other TelemetryStats) TelemetryStats {
	s.Requests += other.Requests
	s.FailedRequests += other.FailedRequests
	s.LogRecords += other.LogRecords
	s.MetricDataPoints += other.MetricDataPoints
	s.Spans += other.Spans
	s.FailedLogRecords += other.FailedLogRecords
	s.FailedMetricDataPoints += other.FailedMetricDataPoints
	s.FailedSpans += other.FailedSpans
	return s
}
