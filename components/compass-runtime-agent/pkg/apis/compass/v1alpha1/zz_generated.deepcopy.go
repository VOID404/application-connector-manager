//go:build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	in.Acquired.DeepCopyInto(&out.Acquired)
	in.NotBefore.DeepCopyInto(&out.NotBefore)
	in.NotAfter.DeepCopyInto(&out.NotAfter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassConnection) DeepCopyInto(out *CompassConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassConnection.
func (in *CompassConnection) DeepCopy() *CompassConnection {
	if in == nil {
		return nil
	}
	out := new(CompassConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompassConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassConnectionList) DeepCopyInto(out *CompassConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompassConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassConnectionList.
func (in *CompassConnectionList) DeepCopy() *CompassConnectionList {
	if in == nil {
		return nil
	}
	out := new(CompassConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompassConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassConnectionSpec) DeepCopyInto(out *CompassConnectionSpec) {
	*out = *in
	out.ManagementInfo = in.ManagementInfo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassConnectionSpec.
func (in *CompassConnectionSpec) DeepCopy() *CompassConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(CompassConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassConnectionStatus) DeepCopyInto(out *CompassConnectionStatus) {
	*out = *in
	if in.ConnectionStatus != nil {
		in, out := &in.ConnectionStatus, &out.ConnectionStatus
		*out = new(ConnectionStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.SynchronizationStatus != nil {
		in, out := &in.SynchronizationStatus, &out.SynchronizationStatus
		*out = new(SynchronizationStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassConnectionStatus.
func (in *CompassConnectionStatus) DeepCopy() *CompassConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(CompassConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionProcessStatus) DeepCopyInto(out *ConnectionProcessStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionProcessStatus.
func (in *ConnectionProcessStatus) DeepCopy() *ConnectionProcessStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionProcessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStatus) DeepCopyInto(out *ConnectionStatus) {
	*out = *in
	in.Established.DeepCopyInto(&out.Established)
	in.Renewed.DeepCopyInto(&out.Renewed)
	in.LastSync.DeepCopyInto(&out.LastSync)
	in.LastSuccess.DeepCopyInto(&out.LastSuccess)
	in.CertificateStatus.DeepCopyInto(&out.CertificateStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStatus.
func (in *ConnectionStatus) DeepCopy() *ConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementInfo) DeepCopyInto(out *ManagementInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementInfo.
func (in *ManagementInfo) DeepCopy() *ManagementInfo {
	if in == nil {
		return nil
	}
	out := new(ManagementInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SynchronizationStatus) DeepCopyInto(out *SynchronizationStatus) {
	*out = *in
	in.LastAttempt.DeepCopyInto(&out.LastAttempt)
	in.LastSuccessfulFetch.DeepCopyInto(&out.LastSuccessfulFetch)
	in.LastSuccessfulApplication.DeepCopyInto(&out.LastSuccessfulApplication)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SynchronizationStatus.
func (in *SynchronizationStatus) DeepCopy() *SynchronizationStatus {
	if in == nil {
		return nil
	}
	out := new(SynchronizationStatus)
	in.DeepCopyInto(out)
	return out
}
