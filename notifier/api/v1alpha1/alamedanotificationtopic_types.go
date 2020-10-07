/*
Copyright 2019 The Alameda Authors.

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

// +kubebuilder:validation:Enum=AlamedaScalerCreate;AlamedaScalerDelete;NodeRegister;DeploymentRegister;DeploymentConfigRegister;PodRegister;NodeDeregister;DeploymentDeregister;DeploymentConfigDeregister;PodDeregister;NodePredictionCreate;PodPredictionCreate;HPARecommendationCreate;HPARecommendationExecute;AnomalyMetricDetect;AnomalyAnalysisCreate;License;EmailNotification;AnomalyForecastDetect;AnomalyRealtimeDetect
type TopicType string

const (
	ScalerCreateTopicType               TopicType = "AlamedaScalerCreate"
	ScalerDeleteTopicType               TopicType = "AlamedaScalerDelete"
	NodeRegisterTopicType               TopicType = "NodeRegister"
	DeploymentRegisterTopicType         TopicType = "DeploymentRegister"
	DeploymentConfigRegisterTopicType   TopicType = "DeploymentConfigRegister"
	PodRegisterTopicType                TopicType = "PodRegister"
	NodeDeregisterTopicType             TopicType = "NodeDeregister"
	DeploymentDeregisterTopicType       TopicType = "DeploymentDeregister"
	DeploymentConfigDeregisterTopicType TopicType = "DeploymentConfigDeregister"
	PodDeregisterTopicType              TopicType = "PodDeregister"
	NodePredictionCreateTopicType       TopicType = "NodePredictionCreate"
	PodPredictionCreateTopicType        TopicType = "PodPredictionCreate"
	VPARecommendationCreateTopicType    TopicType = "VPARecommendationCreate"
	HPARecommendationCreateTopicType    TopicType = "HPARecommendationCreate"
	VPARecommendationExecuteTopicType   TopicType = "VPARecommendationExecute"
	HPARecommendationExecuteTopicType   TopicType = "HPARecommendationExecute"
	AnomalyMetricDetectTopicType        TopicType = "AnomalyMetricDetect"
	AnomalyAnalysisCreateTopicType      TopicType = "AnomalyAnalysisCreate"
	LicenseTopicType                    TopicType = "License"
	EmailNotificationTopicType          TopicType = "EmailNotification"
	AnomalyForecastDetectTopicType      TopicType = "AnomalyForecastDetect"
	AnomalyRealtimeDetectTopicType      TopicType = "AnomalyRealtimeDetect"
)

var (
	TopicTypes = []TopicType{
		ScalerCreateTopicType,
		ScalerDeleteTopicType,
		NodeRegisterTopicType,
		DeploymentRegisterTopicType,
		DeploymentConfigRegisterTopicType,
		PodRegisterTopicType,
		NodeDeregisterTopicType,
		DeploymentDeregisterTopicType,
		DeploymentConfigDeregisterTopicType,
		PodDeregisterTopicType,
		NodePredictionCreateTopicType,
		PodPredictionCreateTopicType,
		VPARecommendationCreateTopicType,
		HPARecommendationCreateTopicType,
		VPARecommendationExecuteTopicType,
		HPARecommendationExecuteTopicType,
		AnomalyMetricDetectTopicType,
		AnomalyAnalysisCreateTopicType,
		LicenseTopicType,
		EmailNotificationTopicType,
		AnomalyForecastDetectTopicType,
		AnomalyRealtimeDetectTopicType,
	}
	// datahub event type mapping
	TopicTypeName = map[int32]TopicType{
		1:  ScalerCreateTopicType,
		2:  ScalerDeleteTopicType,
		3:  NodeRegisterTopicType,
		4:  DeploymentRegisterTopicType,
		5:  DeploymentConfigRegisterTopicType,
		6:  PodRegisterTopicType,
		7:  NodeDeregisterTopicType,
		8:  DeploymentDeregisterTopicType,
		9:  DeploymentConfigDeregisterTopicType,
		10: PodDeregisterTopicType,
		11: NodePredictionCreateTopicType,
		12: PodPredictionCreateTopicType,
		13: VPARecommendationCreateTopicType,
		14: HPARecommendationCreateTopicType,
		15: VPARecommendationExecuteTopicType,
		16: HPARecommendationExecuteTopicType,
		17: AnomalyMetricDetectTopicType,
		18: AnomalyAnalysisCreateTopicType,
		19: LicenseTopicType,
		20: EmailNotificationTopicType,
		21: AnomalyForecastDetectTopicType,
		22: AnomalyRealtimeDetectTopicType,
	}
	TopicTypeValue = map[TopicType]int32{
		ScalerCreateTopicType:               1,
		ScalerDeleteTopicType:               2,
		NodeRegisterTopicType:               3,
		DeploymentRegisterTopicType:         4,
		DeploymentConfigRegisterTopicType:   5,
		PodRegisterTopicType:                6,
		NodeDeregisterTopicType:             7,
		DeploymentDeregisterTopicType:       8,
		DeploymentConfigDeregisterTopicType: 9,
		PodDeregisterTopicType:              10,
		NodePredictionCreateTopicType:       11,
		PodPredictionCreateTopicType:        12,
		VPARecommendationCreateTopicType:    13,
		HPARecommendationCreateTopicType:    14,
		VPARecommendationExecuteTopicType:   15,
		HPARecommendationExecuteTopicType:   16,
		AnomalyMetricDetectTopicType:        17,
		AnomalyAnalysisCreateTopicType:      18,
		LicenseTopicType:                    19,
		EmailNotificationTopicType:          20,
		AnomalyForecastDetectTopicType:      21,
		AnomalyRealtimeDetectTopicType:      22,
	}
)

// +kubebuilder:validation:Enum=debug;info;warning;error;fatal
type LevelType string

const (
	DebugLevelType   LevelType = "debug"
	InfoLevelType    LevelType = "info"
	WarningLevelType LevelType = "warning"
	ErrorLevelType   LevelType = "error"
	FatalLevelType   LevelType = "fatal"
)

var (
	LevelTypes = []LevelType{
		DebugLevelType,
		InfoLevelType,
		WarningLevelType,
		ErrorLevelType,
		FatalLevelType,
	}
	// datahub event level mapping
	LevelTypeName = map[int32]LevelType{
		1: DebugLevelType,
		2: InfoLevelType,
		3: WarningLevelType,
		4: ErrorLevelType,
		5: FatalLevelType,
	}
	LevelTypeValue = map[LevelType]int32{
		DebugLevelType:   1,
		InfoLevelType:    2,
		WarningLevelType: 3,
		ErrorLevelType:   4,
		FatalLevelType:   5,
	}
)

type AlamedaSubject struct {
	Kind       string `json:"kind,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Name       string `json:"name,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type AlamedaTopic struct {
	Type    []TopicType       `json:"type,omitempty"`
	Subject []*AlamedaSubject `json:"subject,omitempty"`
	Level   []LevelType       `json:"level,omitempty"`
	Source  []*AlamedaSource  `json:"source,omitempty"`
}

type AlamedaSource struct {
	Host      string `json:"host,omitempty"`
	Component string `json:"component,omitempty"`
}

type AlamedaChannel struct {
	Emails []*AlamedaEmailChannel `json:"emails,omitempty"`
}

type AlamedaEmailChannel struct {
	Name string   `json:"name"`
	To   []string `json:"to"`
	Cc   []string `json:"cc,omitempty"`
}

// AlamedaNotificationTopicSpec defines the desired state of AlamedaNotificationTopic
type AlamedaNotificationTopicSpec struct {
	Disabled bool            `json:"disabled,omitempty"`
	Topics   []*AlamedaTopic `json:"topics"`
	Channel  *AlamedaChannel `json:"channel"`
}

type AlamedaChannelCondition struct {
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
	Success bool        `json:"success,omitempty"`
	Time    string      `json:"time,omitempty"`
	Message string      `json:"message,omitempty"`
}

// AlamedaNotificationTopicStatus defines the observed state of AlamedaNotificationTopic
type AlamedaNotificationTopicStatus struct {
	ChannelCondictions []*AlamedaChannelCondition `json:"channelConditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=alamedanotificationtopics,scope=Cluster
// AlamedaNotificationTopic is the Schema for the alamedanotificationtopics API
type AlamedaNotificationTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaNotificationTopicSpec   `json:"spec,omitempty"`
	Status AlamedaNotificationTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlamedaNotificationTopicList contains a list of AlamedaNotificationTopic
type AlamedaNotificationTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaNotificationTopic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaNotificationTopic{}, &AlamedaNotificationTopicList{})
}
