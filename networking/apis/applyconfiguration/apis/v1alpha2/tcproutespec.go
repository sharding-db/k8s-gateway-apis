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

package v1alpha2

import (
	v1 "github.com/sharding-db/k8s-gateway-apis/networking/apis/applyconfiguration/apis/v1"
)

// TCPRouteSpecApplyConfiguration represents an declarative configuration of the TCPRouteSpec type for use
// with apply.
type TCPRouteSpecApplyConfiguration struct {
	v1.CommonRouteSpecApplyConfiguration `json:",inline"`
	Rules                                []TCPRouteRuleApplyConfiguration `json:"rules,omitempty"`
}

// TCPRouteSpecApplyConfiguration constructs an declarative configuration of the TCPRouteSpec type for use with
// apply.
func TCPRouteSpec() *TCPRouteSpecApplyConfiguration {
	return &TCPRouteSpecApplyConfiguration{}
}

// WithParentRefs adds the given value to the ParentRefs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ParentRefs field.
func (b *TCPRouteSpecApplyConfiguration) WithParentRefs(values ...*v1.ParentReferenceApplyConfiguration) *TCPRouteSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithParentRefs")
		}
		b.ParentRefs = append(b.ParentRefs, *values[i])
	}
	return b
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *TCPRouteSpecApplyConfiguration) WithRules(values ...*TCPRouteRuleApplyConfiguration) *TCPRouteSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}
