//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlags) DeepCopyInto(out *FeatureFlags) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlags.
func (in *FeatureFlags) DeepCopy() *FeatureFlags {
	if in == nil {
		return nil
	}
	out := new(FeatureFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfig) DeepCopyInto(out *ProjectConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.Flags = in.Flags
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfig.
func (in *ProjectConfig) DeepCopy() *ProjectConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
