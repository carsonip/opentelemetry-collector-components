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
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetServerAddress sets provided value as "server.address" attribute.
func (rb *ResourceBuilder) SetServerAddress(val string) {
	if rb.config.ServerAddress.Enabled {
		rb.res.Attributes().PutStr("server.address", val)
	}
}

// SetServerPort sets provided value as "server.port" attribute.
func (rb *ResourceBuilder) SetServerPort(val string) {
	if rb.config.ServerPort.Enabled {
		rb.res.Attributes().PutStr("server.port", val)
	}
}

// SetValkeyVersion sets provided value as "valkey.version" attribute.
func (rb *ResourceBuilder) SetValkeyVersion(val string) {
	if rb.config.ValkeyVersion.Enabled {
		rb.res.Attributes().PutStr("valkey.version", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
