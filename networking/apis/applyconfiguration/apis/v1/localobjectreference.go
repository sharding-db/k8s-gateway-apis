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
	v1 "github.com/sharding-db/k8s-gateway-apis/networking/apis/v1"
)

// LocalObjectReferenceApplyConfiguration represents an declarative configuration of the LocalObjectReference type for use
// with apply.
type LocalObjectReferenceApplyConfiguration struct {
	Group *v1.Group      `json:"group,omitempty"`
	Kind  *v1.Kind       `json:"kind,omitempty"`
	Name  *v1.ObjectName `json:"name,omitempty"`
}

// LocalObjectReferenceApplyConfiguration constructs an declarative configuration of the LocalObjectReference type for use with
// apply.
func LocalObjectReference() *LocalObjectReferenceApplyConfiguration {
	return &LocalObjectReferenceApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *LocalObjectReferenceApplyConfiguration) WithGroup(value v1.Group) *LocalObjectReferenceApplyConfiguration {
	b.Group = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *LocalObjectReferenceApplyConfiguration) WithKind(value v1.Kind) *LocalObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *LocalObjectReferenceApplyConfiguration) WithName(value v1.ObjectName) *LocalObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}
