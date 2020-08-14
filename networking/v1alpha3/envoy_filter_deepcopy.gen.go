// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/envoy_filter.proto

package v1alpha3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using EnvoyFilter within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter) DeepCopyInto(out *EnvoyFilter) {
	p := proto.Clone(in).(*EnvoyFilter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter. Required by controller-gen.
func (in *EnvoyFilter) DeepCopy() *EnvoyFilter {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ProxyMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ProxyMatch) DeepCopyInto(out *EnvoyFilter_ProxyMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ProxyMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ProxyMatch. Required by controller-gen.
func (in *EnvoyFilter_ProxyMatch) DeepCopy() *EnvoyFilter_ProxyMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ProxyMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ClusterMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ClusterMatch) DeepCopyInto(out *EnvoyFilter_ClusterMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ClusterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ClusterMatch. Required by controller-gen.
func (in *EnvoyFilter_ClusterMatch) DeepCopy() *EnvoyFilter_ClusterMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ClusterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_RouteConfigurationMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_RouteConfigurationMatch) DeepCopyInto(out *EnvoyFilter_RouteConfigurationMatch) {
	p := proto.Clone(in).(*EnvoyFilter_RouteConfigurationMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_RouteConfigurationMatch. Required by controller-gen.
func (in *EnvoyFilter_RouteConfigurationMatch) DeepCopy() *EnvoyFilter_RouteConfigurationMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_RouteConfigurationMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_RouteConfigurationMatch_RouteMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_RouteConfigurationMatch_RouteMatch) DeepCopyInto(out *EnvoyFilter_RouteConfigurationMatch_RouteMatch) {
	p := proto.Clone(in).(*EnvoyFilter_RouteConfigurationMatch_RouteMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_RouteConfigurationMatch_RouteMatch. Required by controller-gen.
func (in *EnvoyFilter_RouteConfigurationMatch_RouteMatch) DeepCopy() *EnvoyFilter_RouteConfigurationMatch_RouteMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_RouteConfigurationMatch_RouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch) DeepCopyInto(out *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch) {
	p := proto.Clone(in).(*EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch. Required by controller-gen.
func (in *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch) DeepCopy() *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch) DeepCopy() *EnvoyFilter_ListenerMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch_FilterChainMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch_FilterChainMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch_FilterChainMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch_FilterChainMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch_FilterChainMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch_FilterChainMatch) DeepCopy() *EnvoyFilter_ListenerMatch_FilterChainMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch_FilterChainMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch_FilterMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch_FilterMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch_FilterMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch_FilterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch_FilterMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch_FilterMatch) DeepCopy() *EnvoyFilter_ListenerMatch_FilterMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch_FilterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch_SubFilterMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch_SubFilterMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch_SubFilterMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch_SubFilterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch_SubFilterMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch_SubFilterMatch) DeepCopy() *EnvoyFilter_ListenerMatch_SubFilterMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch_SubFilterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_Patch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_Patch) DeepCopyInto(out *EnvoyFilter_Patch) {
	p := proto.Clone(in).(*EnvoyFilter_Patch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_Patch. Required by controller-gen.
func (in *EnvoyFilter_Patch) DeepCopy() *EnvoyFilter_Patch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_EnvoyConfigObjectMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_EnvoyConfigObjectMatch) DeepCopyInto(out *EnvoyFilter_EnvoyConfigObjectMatch) {
	p := proto.Clone(in).(*EnvoyFilter_EnvoyConfigObjectMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_EnvoyConfigObjectMatch. Required by controller-gen.
func (in *EnvoyFilter_EnvoyConfigObjectMatch) DeepCopy() *EnvoyFilter_EnvoyConfigObjectMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_EnvoyConfigObjectMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_EnvoyConfigObjectPatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_EnvoyConfigObjectPatch) DeepCopyInto(out *EnvoyFilter_EnvoyConfigObjectPatch) {
	p := proto.Clone(in).(*EnvoyFilter_EnvoyConfigObjectPatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_EnvoyConfigObjectPatch. Required by controller-gen.
func (in *EnvoyFilter_EnvoyConfigObjectPatch) DeepCopy() *EnvoyFilter_EnvoyConfigObjectPatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_EnvoyConfigObjectPatch)
	in.DeepCopyInto(out)
	return out
}
