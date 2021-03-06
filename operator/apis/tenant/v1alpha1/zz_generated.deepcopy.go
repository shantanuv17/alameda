// +build !ignore_autogenerated

/*
Copyright 2020 The Alameda Authors.

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
func (in *AlamedaCluster) DeepCopyInto(out *AlamedaCluster) {
	*out = *in
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(AlamedaDataSource)
		(*in).DeepCopyInto(*out)
	}
	if in.WatchedNamespace != nil {
		in, out := &in.WatchedNamespace, &out.WatchedNamespace
		*out = new(AlamedaWatchedNamespace)
		(*in).DeepCopyInto(*out)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]AlamedaFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaCluster.
func (in *AlamedaCluster) DeepCopy() *AlamedaCluster {
	if in == nil {
		return nil
	}
	out := new(AlamedaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaDataSource) DeepCopyInto(out *AlamedaDataSource) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]AlamedaKey, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaDataSource.
func (in *AlamedaDataSource) DeepCopy() *AlamedaDataSource {
	if in == nil {
		return nil
	}
	out := new(AlamedaDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaFeature) DeepCopyInto(out *AlamedaFeature) {
	*out = *in
	in.AlamedaFeatureSet.DeepCopyInto(&out.AlamedaFeatureSet)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaFeature.
func (in *AlamedaFeature) DeepCopy() *AlamedaFeature {
	if in == nil {
		return nil
	}
	out := new(AlamedaFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaFeatureSet) DeepCopyInto(out *AlamedaFeatureSet) {
	*out = *in
	if in.ResourcePlanning != nil {
		in, out := &in.ResourcePlanning, &out.ResourcePlanning
		*out = new(ResourcePlanningFeature)
		**out = **in
	}
	if in.CostAnalysis != nil {
		in, out := &in.CostAnalysis, &out.CostAnalysis
		*out = new(CostAnalysisFeature)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaFeatureSet.
func (in *AlamedaFeatureSet) DeepCopy() *AlamedaFeatureSet {
	if in == nil {
		return nil
	}
	out := new(AlamedaFeatureSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaKey) DeepCopyInto(out *AlamedaKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaKey.
func (in *AlamedaKey) DeepCopy() *AlamedaKey {
	if in == nil {
		return nil
	}
	out := new(AlamedaKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaOrganization) DeepCopyInto(out *AlamedaOrganization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaOrganization.
func (in *AlamedaOrganization) DeepCopy() *AlamedaOrganization {
	if in == nil {
		return nil
	}
	out := new(AlamedaOrganization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaOrganization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaOrganizationList) DeepCopyInto(out *AlamedaOrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlamedaOrganization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaOrganizationList.
func (in *AlamedaOrganizationList) DeepCopy() *AlamedaOrganizationList {
	if in == nil {
		return nil
	}
	out := new(AlamedaOrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaOrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaOrganizationSpec) DeepCopyInto(out *AlamedaOrganizationSpec) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]AlamedaFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(AlamedaDataSource)
		(*in).DeepCopyInto(*out)
	}
	if in.WatchedNamespace != nil {
		in, out := &in.WatchedNamespace, &out.WatchedNamespace
		*out = new(AlamedaWatchedNamespace)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]AlamedaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaOrganizationSpec.
func (in *AlamedaOrganizationSpec) DeepCopy() *AlamedaOrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaOrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaOrganizationStatus) DeepCopyInto(out *AlamedaOrganizationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaOrganizationStatus.
func (in *AlamedaOrganizationStatus) DeepCopy() *AlamedaOrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(AlamedaOrganizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaWatchedNamespace) DeepCopyInto(out *AlamedaWatchedNamespace) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaWatchedNamespace.
func (in *AlamedaWatchedNamespace) DeepCopy() *AlamedaWatchedNamespace {
	if in == nil {
		return nil
	}
	out := new(AlamedaWatchedNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostAnalysisFeature) DeepCopyInto(out *CostAnalysisFeature) {
	*out = *in
	out.FeatureMeta = in.FeatureMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostAnalysisFeature.
func (in *CostAnalysisFeature) DeepCopy() *CostAnalysisFeature {
	if in == nil {
		return nil
	}
	out := new(CostAnalysisFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureMeta) DeepCopyInto(out *FeatureMeta) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMeta.
func (in *FeatureMeta) DeepCopy() *FeatureMeta {
	if in == nil {
		return nil
	}
	out := new(FeatureMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePlanningFeature) DeepCopyInto(out *ResourcePlanningFeature) {
	*out = *in
	out.FeatureMeta = in.FeatureMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePlanningFeature.
func (in *ResourcePlanningFeature) DeepCopy() *ResourcePlanningFeature {
	if in == nil {
		return nil
	}
	out := new(ResourcePlanningFeature)
	in.DeepCopyInto(out)
	return out
}
