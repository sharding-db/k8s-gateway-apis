/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apisv1 "github.com/sharding-db/k8s-gateway-apis/networking/apis/v1"
)

// HTTPRouteMatchApplyConfiguration represents an declarative configuration of the HTTPRouteMatch type for use
// with apply.
type HTTPRouteMatchApplyConfiguration struct {
	Path        *HTTPPathMatchApplyConfiguration        `json:"path,omitempty"`
	Headers     []HTTPHeaderMatchApplyConfiguration     `json:"headers,omitempty"`
	QueryParams []HTTPQueryParamMatchApplyConfiguration `json:"queryParams,omitempty"`
	Method      *apisv1.HTTPMethod                      `json:"method,omitempty"`
}

// HTTPRouteMatchApplyConfiguration constructs an declarative configuration of the HTTPRouteMatch type for use with
// apply.
func HTTPRouteMatch() *HTTPRouteMatchApplyConfiguration {
	return &HTTPRouteMatchApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *HTTPRouteMatchApplyConfiguration) WithPath(value *HTTPPathMatchApplyConfiguration) *HTTPRouteMatchApplyConfiguration {
	b.Path = value
	return b
}

// WithHeaders adds the given value to the Headers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Headers field.
func (b *HTTPRouteMatchApplyConfiguration) WithHeaders(values ...*HTTPHeaderMatchApplyConfiguration) *HTTPRouteMatchApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHeaders")
		}
		b.Headers = append(b.Headers, *values[i])
	}
	return b
}

// WithQueryParams adds the given value to the QueryParams field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the QueryParams field.
func (b *HTTPRouteMatchApplyConfiguration) WithQueryParams(values ...*HTTPQueryParamMatchApplyConfiguration) *HTTPRouteMatchApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithQueryParams")
		}
		b.QueryParams = append(b.QueryParams, *values[i])
	}
	return b
}

// WithMethod sets the Method field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Method field is set to the value of the last call.
func (b *HTTPRouteMatchApplyConfiguration) WithMethod(value apisv1.HTTPMethod) *HTTPRouteMatchApplyConfiguration {
	b.Method = &value
	return b
}
