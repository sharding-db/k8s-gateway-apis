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

package v1alpha3

import (
	v1alpha2 "github.com/sharding-db/k8s-gateway-apis/networking/apis/applyconfiguration/apis/v1alpha2"
)

// BackendTLSPolicySpecApplyConfiguration represents an declarative configuration of the BackendTLSPolicySpec type for use
// with apply.
type BackendTLSPolicySpecApplyConfiguration struct {
	TargetRefs []v1alpha2.LocalPolicyTargetReferenceWithSectionNameApplyConfiguration `json:"targetRefs,omitempty"`
	Validation *BackendTLSPolicyValidationApplyConfiguration                          `json:"validation,omitempty"`
}

// BackendTLSPolicySpecApplyConfiguration constructs an declarative configuration of the BackendTLSPolicySpec type for use with
// apply.
func BackendTLSPolicySpec() *BackendTLSPolicySpecApplyConfiguration {
	return &BackendTLSPolicySpecApplyConfiguration{}
}

// WithTargetRefs adds the given value to the TargetRefs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TargetRefs field.
func (b *BackendTLSPolicySpecApplyConfiguration) WithTargetRefs(values ...*v1alpha2.LocalPolicyTargetReferenceWithSectionNameApplyConfiguration) *BackendTLSPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTargetRefs")
		}
		b.TargetRefs = append(b.TargetRefs, *values[i])
	}
	return b
}

// WithValidation sets the Validation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Validation field is set to the value of the last call.
func (b *BackendTLSPolicySpecApplyConfiguration) WithValidation(value *BackendTLSPolicyValidationApplyConfiguration) *BackendTLSPolicySpecApplyConfiguration {
	b.Validation = value
	return b
}
