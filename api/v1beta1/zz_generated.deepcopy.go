//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneNodeSection) DeepCopyInto(out *DataPlaneNodeSection) {
	*out = *in
	in.Node.DeepCopyInto(&out.Node)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneNodeSection.
func (in *DataPlaneNodeSection) DeepCopy() *DataPlaneNodeSection {
	if in == nil {
		return nil
	}
	out := new(DataPlaneNodeSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployStrategySection) DeepCopyInto(out *DeployStrategySection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployStrategySection.
func (in *DeployStrategySection) DeepCopy() *DeployStrategySection {
	if in == nil {
		return nil
	}
	out := new(DeployStrategySection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigSection) DeepCopyInto(out *NetworkConfigSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigSection.
func (in *NetworkConfigSection) DeepCopy() *NetworkConfigSection {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksSection) DeepCopyInto(out *NetworksSection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksSection.
func (in *NetworksSection) DeepCopy() *NetworksSection {
	if in == nil {
		return nil
	}
	out := new(NetworksSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSection) DeepCopyInto(out *NodeSection) {
	*out = *in
	out.NetworkConfig = in.NetworkConfig
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworksSection, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSection.
func (in *NodeSection) DeepCopy() *NodeSection {
	if in == nil {
		return nil
	}
	out := new(NodeSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlane) DeepCopyInto(out *OpenStackDataPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlane.
func (in *OpenStackDataPlane) DeepCopy() *OpenStackDataPlane {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneList) DeepCopyInto(out *OpenStackDataPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneList.
func (in *OpenStackDataPlaneList) DeepCopy() *OpenStackDataPlaneList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNode) DeepCopyInto(out *OpenStackDataPlaneNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNode.
func (in *OpenStackDataPlaneNode) DeepCopy() *OpenStackDataPlaneNode {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNodeList) DeepCopyInto(out *OpenStackDataPlaneNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlaneNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNodeList.
func (in *OpenStackDataPlaneNodeList) DeepCopy() *OpenStackDataPlaneNodeList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNodeSpec) DeepCopyInto(out *OpenStackDataPlaneNodeSpec) {
	*out = *in
	in.Node.DeepCopyInto(&out.Node)
	out.DeployStrategy = in.DeployStrategy
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNodeSpec.
func (in *OpenStackDataPlaneNodeSpec) DeepCopy() *OpenStackDataPlaneNodeSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneNodeStatus) DeepCopyInto(out *OpenStackDataPlaneNodeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneNodeStatus.
func (in *OpenStackDataPlaneNodeStatus) DeepCopy() *OpenStackDataPlaneNodeStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRole) DeepCopyInto(out *OpenStackDataPlaneRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRole.
func (in *OpenStackDataPlaneRole) DeepCopy() *OpenStackDataPlaneRole {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRoleList) DeepCopyInto(out *OpenStackDataPlaneRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackDataPlaneRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRoleList.
func (in *OpenStackDataPlaneRoleList) DeepCopy() *OpenStackDataPlaneRoleList {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackDataPlaneRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRoleSpec) DeepCopyInto(out *OpenStackDataPlaneRoleSpec) {
	*out = *in
	if in.DataPlaneNodes != nil {
		in, out := &in.DataPlaneNodes, &out.DataPlaneNodes
		*out = make([]OpenStackDataPlaneNodeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NodeTemplate.DeepCopyInto(&out.NodeTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRoleSpec.
func (in *OpenStackDataPlaneRoleSpec) DeepCopy() *OpenStackDataPlaneRoleSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneRoleStatus) DeepCopyInto(out *OpenStackDataPlaneRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneRoleStatus.
func (in *OpenStackDataPlaneRoleStatus) DeepCopy() *OpenStackDataPlaneRoleStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneSpec) DeepCopyInto(out *OpenStackDataPlaneSpec) {
	*out = *in
	if in.DataPlaneRoles != nil {
		in, out := &in.DataPlaneRoles, &out.DataPlaneRoles
		*out = make([]OpenStackDataPlaneRoleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneSpec.
func (in *OpenStackDataPlaneSpec) DeepCopy() *OpenStackDataPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackDataPlaneStatus) DeepCopyInto(out *OpenStackDataPlaneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackDataPlaneStatus.
func (in *OpenStackDataPlaneStatus) DeepCopy() *OpenStackDataPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackDataPlaneStatus)
	in.DeepCopyInto(out)
	return out
}
