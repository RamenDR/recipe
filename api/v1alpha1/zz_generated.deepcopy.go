//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 IBM Corp.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Check) DeepCopyInto(out *Check) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Check.
func (in *Check) DeepCopy() *Check {
	if in == nil {
		return nil
	}
	out := new(Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	if in.IncludedResourceTypes != nil {
		in, out := &in.IncludedResourceTypes, &out.IncludedResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedResourceTypes != nil {
		in, out := &in.ExcludedResourceTypes, &out.ExcludedResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	if in.IncludeClusterResources != nil {
		in, out := &in.IncludeClusterResources, &out.IncludeClusterResources
		*out = new(bool)
		**out = **in
	}
	if in.Essential != nil {
		in, out := &in.Essential, &out.Essential
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hook) DeepCopyInto(out *Hook) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	if in.Ops != nil {
		in, out := &in.Ops, &out.Ops
		*out = make([]*Operation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Operation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Chks != nil {
		in, out := &in.Chks, &out.Chks
		*out = make([]*Check, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Check)
				**out = **in
			}
		}
	}
	if in.Essential != nil {
		in, out := &in.Essential, &out.Essential
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hook.
func (in *Hook) DeepCopy() *Hook {
	if in == nil {
		return nil
	}
	out := new(Hook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recipe) DeepCopyInto(out *Recipe) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recipe.
func (in *Recipe) DeepCopy() *Recipe {
	if in == nil {
		return nil
	}
	out := new(Recipe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Recipe) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecipeList) DeepCopyInto(out *RecipeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Recipe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecipeList.
func (in *RecipeList) DeepCopy() *RecipeList {
	if in == nil {
		return nil
	}
	out := new(RecipeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecipeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecipeSpec) DeepCopyInto(out *RecipeSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*Group, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Group)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = make([]*Hook, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Hook)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Workflows != nil {
		in, out := &in.Workflows, &out.Workflows
		*out = make([]*Workflow, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Workflow)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecipeSpec.
func (in *RecipeSpec) DeepCopy() *RecipeSpec {
	if in == nil {
		return nil
	}
	out := new(RecipeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecipeStatus) DeepCopyInto(out *RecipeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecipeStatus.
func (in *RecipeStatus) DeepCopy() *RecipeStatus {
	if in == nil {
		return nil
	}
	out := new(RecipeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workflow) DeepCopyInto(out *Workflow) {
	*out = *in
	if in.Sequence != nil {
		in, out := &in.Sequence, &out.Sequence
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workflow.
func (in *Workflow) DeepCopy() *Workflow {
	if in == nil {
		return nil
	}
	out := new(Workflow)
	in.DeepCopyInto(out)
	return out
}
