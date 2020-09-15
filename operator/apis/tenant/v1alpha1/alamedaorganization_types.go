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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=include;exclude
type LogicOperator string

// +kubebuilder:validation:Enum=datadog;prometheus
type DataSourceType string

// +kubebuilder:validation:Enum=resourcePlanning;costAnalysis
type FeatureType string

// +kubebuilder:validation:Enum=true;false
type FeatureEnabled bool

// +kubebuilder:validation:Enum=uploadResult;localOnly
type FeatureMode string

const (
	ResourcePlanningFeatureType = "resourcePlanning"
	CostAnalysisFeatureType     = "costAnalysis"
)

type FeatureMeta struct {
	Enabled FeatureEnabled `json:"enabled"`
	Mode    FeatureMode    `json:"mode"`
}
type ResourcePlanningFeature struct {
	FeatureMeta `json:",inline"`
}

type CostAnalysisFeature struct {
	FeatureMeta `json:",inline"`
}

type AlamedaKey struct {
	Key string `json:"key"`
}
type AlamedaDataSource struct {
	Type DataSourceType `json:"type"`
	// +optional
	Keys    []AlamedaKey `json:"keys,omitempty"`
	Account string       `json:"account"`
	Address string       `json:"address"`
}

type AlamedaFeature struct {
	Type FeatureType `json:"type,omitempty"`
	// +optional
	ResourcePlanning ResourcePlanningFeature `json:"resourcePlanning,omitempty"`
	// +optional
	CostAnalysis CostAnalysisFeature `json:"costAnalysis,omitempty"`
}

type AlamedaWatchedNamespace struct {
	// +optional
	Names    []string      `json:"names,omitempty"`
	Operator LogicOperator `json:"operator"`
}

type AlamedaCluster struct {
	Name string `json:"name"`
	// +optional
	DataSource AlamedaDataSource `json:"dataSource,omitempty"`
	// +optional
	Features []AlamedaFeature `json:"features,omitempty"`
	// +optional
	WatchedNamespace AlamedaWatchedNamespace `json:"watchedNamespace,omitempty"`
}

// AlamedaOrganizationSpec defines the desired state of AlamedaOrganization
type AlamedaOrganizationSpec struct {
	Tenant string `json:"tenant"`
	// +optional
	Clusters []AlamedaCluster `json:"clusters,omitempty"`
}

// AlamedaOrganizationStatus defines the observed state of AlamedaOrganization
type AlamedaOrganizationStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// AlamedaOrganization is the Schema for the alamedaorganizations API
type AlamedaOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaOrganizationSpec   `json:"spec,omitempty"`
	Status AlamedaOrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlamedaOrganizationList contains a list of AlamedaOrganization
type AlamedaOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaOrganization `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaOrganization{}, &AlamedaOrganizationList{})
}
