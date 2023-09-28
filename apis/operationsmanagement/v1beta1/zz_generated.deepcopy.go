//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolution) DeepCopyInto(out *LogAnalyticsSolution) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolution.
func (in *LogAnalyticsSolution) DeepCopy() *LogAnalyticsSolution {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogAnalyticsSolution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolutionInitParameters) DeepCopyInto(out *LogAnalyticsSolutionInitParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]PlanInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SolutionName != nil {
		in, out := &in.SolutionName, &out.SolutionName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolutionInitParameters.
func (in *LogAnalyticsSolutionInitParameters) DeepCopy() *LogAnalyticsSolutionInitParameters {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolutionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolutionList) DeepCopyInto(out *LogAnalyticsSolutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogAnalyticsSolution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolutionList.
func (in *LogAnalyticsSolutionList) DeepCopy() *LogAnalyticsSolutionList {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogAnalyticsSolutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolutionObservation) DeepCopyInto(out *LogAnalyticsSolutionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]PlanObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SolutionName != nil {
		in, out := &in.SolutionName, &out.SolutionName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WorkspaceName != nil {
		in, out := &in.WorkspaceName, &out.WorkspaceName
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceResourceID != nil {
		in, out := &in.WorkspaceResourceID, &out.WorkspaceResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolutionObservation.
func (in *LogAnalyticsSolutionObservation) DeepCopy() *LogAnalyticsSolutionObservation {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolutionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolutionParameters) DeepCopyInto(out *LogAnalyticsSolutionParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make([]PlanParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SolutionName != nil {
		in, out := &in.SolutionName, &out.SolutionName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WorkspaceName != nil {
		in, out := &in.WorkspaceName, &out.WorkspaceName
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceNameRef != nil {
		in, out := &in.WorkspaceNameRef, &out.WorkspaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceNameSelector != nil {
		in, out := &in.WorkspaceNameSelector, &out.WorkspaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceResourceID != nil {
		in, out := &in.WorkspaceResourceID, &out.WorkspaceResourceID
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceResourceIDRef != nil {
		in, out := &in.WorkspaceResourceIDRef, &out.WorkspaceResourceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceResourceIDSelector != nil {
		in, out := &in.WorkspaceResourceIDSelector, &out.WorkspaceResourceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolutionParameters.
func (in *LogAnalyticsSolutionParameters) DeepCopy() *LogAnalyticsSolutionParameters {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolutionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolutionSpec) DeepCopyInto(out *LogAnalyticsSolutionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolutionSpec.
func (in *LogAnalyticsSolutionSpec) DeepCopy() *LogAnalyticsSolutionSpec {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAnalyticsSolutionStatus) DeepCopyInto(out *LogAnalyticsSolutionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAnalyticsSolutionStatus.
func (in *LogAnalyticsSolutionStatus) DeepCopy() *LogAnalyticsSolutionStatus {
	if in == nil {
		return nil
	}
	out := new(LogAnalyticsSolutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanInitParameters) DeepCopyInto(out *PlanInitParameters) {
	*out = *in
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.PromotionCode != nil {
		in, out := &in.PromotionCode, &out.PromotionCode
		*out = new(string)
		**out = **in
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanInitParameters.
func (in *PlanInitParameters) DeepCopy() *PlanInitParameters {
	if in == nil {
		return nil
	}
	out := new(PlanInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanObservation) DeepCopyInto(out *PlanObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.PromotionCode != nil {
		in, out := &in.PromotionCode, &out.PromotionCode
		*out = new(string)
		**out = **in
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanObservation.
func (in *PlanObservation) DeepCopy() *PlanObservation {
	if in == nil {
		return nil
	}
	out := new(PlanObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanParameters) DeepCopyInto(out *PlanParameters) {
	*out = *in
	if in.Product != nil {
		in, out := &in.Product, &out.Product
		*out = new(string)
		**out = **in
	}
	if in.PromotionCode != nil {
		in, out := &in.PromotionCode, &out.PromotionCode
		*out = new(string)
		**out = **in
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanParameters.
func (in *PlanParameters) DeepCopy() *PlanParameters {
	if in == nil {
		return nil
	}
	out := new(PlanParameters)
	in.DeepCopyInto(out)
	return out
}
