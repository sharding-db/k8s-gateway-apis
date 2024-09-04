//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"github.com/sharding-db/k8s-gateway-apis/networking/apis/v1"
	"github.com/sharding-db/k8s-gateway-apis/networking/apis/v1alpha2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendTLSPolicy) DeepCopyInto(out *BackendTLSPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendTLSPolicy.
func (in *BackendTLSPolicy) DeepCopy() *BackendTLSPolicy {
	if in == nil {
		return nil
	}
	out := new(BackendTLSPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendTLSPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendTLSPolicyList) DeepCopyInto(out *BackendTLSPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackendTLSPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendTLSPolicyList.
func (in *BackendTLSPolicyList) DeepCopy() *BackendTLSPolicyList {
	if in == nil {
		return nil
	}
	out := new(BackendTLSPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendTLSPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendTLSPolicySpec) DeepCopyInto(out *BackendTLSPolicySpec) {
	*out = *in
	if in.TargetRefs != nil {
		in, out := &in.TargetRefs, &out.TargetRefs
		*out = make([]v1alpha2.LocalPolicyTargetReferenceWithSectionName, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Validation.DeepCopyInto(&out.Validation)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendTLSPolicySpec.
func (in *BackendTLSPolicySpec) DeepCopy() *BackendTLSPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BackendTLSPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendTLSPolicyValidation) DeepCopyInto(out *BackendTLSPolicyValidation) {
	*out = *in
	if in.CACertificateRefs != nil {
		in, out := &in.CACertificateRefs, &out.CACertificateRefs
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.WellKnownCACertificates != nil {
		in, out := &in.WellKnownCACertificates, &out.WellKnownCACertificates
		*out = new(WellKnownCACertificatesType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendTLSPolicyValidation.
func (in *BackendTLSPolicyValidation) DeepCopy() *BackendTLSPolicyValidation {
	if in == nil {
		return nil
	}
	out := new(BackendTLSPolicyValidation)
	in.DeepCopyInto(out)
	return out
}
