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

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// AttributeValkeyClientConnectionState specifies the a value valkey.client.connection.state attribute.
type AttributeValkeyClientConnectionState int

const (
	_ AttributeValkeyClientConnectionState = iota
	AttributeValkeyClientConnectionStateUsed
	AttributeValkeyClientConnectionStateBlocked
	AttributeValkeyClientConnectionStateTracking
)

// String returns the string representation of the AttributeValkeyClientConnectionState.
func (av AttributeValkeyClientConnectionState) String() string {
	switch av {
	case AttributeValkeyClientConnectionStateUsed:
		return "used"
	case AttributeValkeyClientConnectionStateBlocked:
		return "blocked"
	case AttributeValkeyClientConnectionStateTracking:
		return "tracking"
	}
	return ""
}

// MapAttributeValkeyClientConnectionState is a helper map of string to AttributeValkeyClientConnectionState attribute value.
var MapAttributeValkeyClientConnectionState = map[string]AttributeValkeyClientConnectionState{
	"used":     AttributeValkeyClientConnectionStateUsed,
	"blocked":  AttributeValkeyClientConnectionStateBlocked,
	"tracking": AttributeValkeyClientConnectionStateTracking,
}

type metricValkeyClientConnectionCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills valkey.client.connection.count metric with initial data.
func (m *metricValkeyClientConnectionCount) init() {
	m.data.SetName("valkey.client.connection.count")
	m.data.SetDescription("The number of connections that are currently in state described by the state attribute")
	m.data.SetUnit("{connection}")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricValkeyClientConnectionCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, valkeyClientConnectionStateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("valkey.client.connection.state", valkeyClientConnectionStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricValkeyClientConnectionCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricValkeyClientConnectionCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricValkeyClientConnectionCount(cfg MetricConfig) metricValkeyClientConnectionCount {
	m := metricValkeyClientConnectionCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                            MetricsBuilderConfig // config of the metrics builder.
	startTime                         pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                   int                  // maximum observed number of metrics per resource.
	metricsBuffer                     pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                         component.BuildInfo  // contains version information.
	resourceAttributeIncludeFilter    map[string]filter.Filter
	resourceAttributeExcludeFilter    map[string]filter.Filter
	metricValkeyClientConnectionCount metricValkeyClientConnectionCount
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                            mbc,
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		buildInfo:                         settings.BuildInfo,
		metricValkeyClientConnectionCount: newMetricValkeyClientConnectionCount(mbc.Metrics.ValkeyClientConnectionCount),
		resourceAttributeIncludeFilter:    make(map[string]filter.Filter),
		resourceAttributeExcludeFilter:    make(map[string]filter.Filter),
	}
	if mbc.ResourceAttributes.ServerAddress.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["server.address"] = filter.CreateFilter(mbc.ResourceAttributes.ServerAddress.MetricsInclude)
	}
	if mbc.ResourceAttributes.ServerAddress.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["server.address"] = filter.CreateFilter(mbc.ResourceAttributes.ServerAddress.MetricsExclude)
	}
	if mbc.ResourceAttributes.ServerPort.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["server.port"] = filter.CreateFilter(mbc.ResourceAttributes.ServerPort.MetricsInclude)
	}
	if mbc.ResourceAttributes.ServerPort.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["server.port"] = filter.CreateFilter(mbc.ResourceAttributes.ServerPort.MetricsExclude)
	}
	if mbc.ResourceAttributes.ValkeyVersion.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["valkey.version"] = filter.CreateFilter(mbc.ResourceAttributes.ValkeyVersion.MetricsInclude)
	}
	if mbc.ResourceAttributes.ValkeyVersion.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["valkey.version"] = filter.CreateFilter(mbc.ResourceAttributes.ValkeyVersion.MetricsExclude)
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/rogercoll/telemetrygen")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricValkeyClientConnectionCount.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}
	for attr, filter := range mb.resourceAttributeIncludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && !filter.Matches(val.AsString()) {
			return
		}
	}
	for attr, filter := range mb.resourceAttributeExcludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && filter.Matches(val.AsString()) {
			return
		}
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordValkeyClientConnectionCountDataPoint adds a data point to valkey.client.connection.count metric.
func (mb *MetricsBuilder) RecordValkeyClientConnectionCountDataPoint(ts pcommon.Timestamp, val int64, valkeyClientConnectionStateAttributeValue AttributeValkeyClientConnectionState) {
	mb.metricValkeyClientConnectionCount.recordDataPoint(mb.startTime, ts, val, valkeyClientConnectionStateAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
