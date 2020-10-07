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

// +kubebuilder:validation:Enum=email
type ChannelType string

const (
	EmailChannelType = "email"
)

type AlamedaEmail struct {
	Server     string `json:"server"`
	Port       uint16 `json:"port"`
	From       string `json:"from"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Encryption string `json:"encryption,omitempty"`
}

type Channel struct {
	Email *AlamedaEmail `json:"email,omitempty"`
}

// AlamedaNotificationChannelSpec defines the desired state of AlamedaNotificationChannel
type AlamedaNotificationChannelSpec struct {
	Channel `json:",inline"`
	Type    ChannelType `json:"type"`
}

// AlamedaNotificationChannelStatus defines the observed state of AlamedaNotificationChannel
type AlamedaNotificationChannelStatus struct {
	ChannelTest *AlamedaChannelTest `json:"channelTest,omitempty"`
}

type AlamedaChannelTest struct {
	Success bool   `json:"success"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=alamedanotificationchannels,scope=Cluster
// AlamedaNotificationChannel is the Schema for the alamedanotificationchannels API
type AlamedaNotificationChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlamedaNotificationChannelSpec   `json:"spec,omitempty"`
	Status AlamedaNotificationChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// AlamedaNotificationChannelList contains a list of AlamedaNotificationChannel
type AlamedaNotificationChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlamedaNotificationChannel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlamedaNotificationChannel{}, &AlamedaNotificationChannelList{})
}
