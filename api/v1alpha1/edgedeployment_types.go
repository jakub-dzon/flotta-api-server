/*
Copyright 2021.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EdgeDeploymentSpec defines the desired state of EdgeDeployment
type EdgeDeploymentSpec struct {
	DeviceSelector  *metav1.LabelSelector          `json:"deviceSelector,omitempty"`
	Device          string                         `json:"device,omitempty"`
	Type            EdgeDeploymentType             `json:"type"`
	Pod             Pod                            `json:"pod,omitempty"`
	Data            *DataConfiguration             `json:"data,omitempty"`
	ImageRegistries *ImageRegistriesConfiguration  `json:"imageRegistries,omitempty"`
	Metrics         *ContainerMetricsConfiguration `json:"metrics,omitempty"`
}

type ImageRegistriesConfiguration struct {
	AuthFileSecret *NameRef `json:"secretRef,omitempty"`
}

type MetricsConfigEntity struct {
	// Path to use when retrieving metrics
	// +kubebuilder:default=/
	Path string `json:"path,omitempty"`

	// Port to use when retrieve the metrics
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`

	Disabled bool `json:"disabled,omitempty"`
}

type ContainerMetricsConfiguration struct {
	// Path to use when retrieving metrics
	// +kubebuilder:default=/
	Path string `json:"path,omitempty"`

	// Port to use when retrieve the metrics
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`

	// Interval(in seconds) to scrape metrics endpoint.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=60
	Interval int32 `json:"interval,omitempty"`

	// Specification of workload metrics to be collected
	AllowList *NameRef `json:"allowList,omitempty"`

	Containers map[string]*MetricsConfigEntity `json:"containers,omitempty"`
}

type NameRef struct {
	Name string `json:"name"`
}

type DataConfiguration struct {
	Paths []DataPath `json:"paths,omitempty"`
}

type DataPath struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Pod struct {
	Spec v1.PodSpec `json:"spec"`
}

type EdgeDeploymentType string

const (
	PodDeploymentType EdgeDeploymentType = "pod"
)

// EdgeDeploymentStatus defines the observed state of EdgeDeployment
type EdgeDeploymentStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EdgeDeployment is the Schema for the edgedeployments API
type EdgeDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeDeploymentSpec   `json:"spec,omitempty"`
	Status EdgeDeploymentStatus `json:"status,omitempty"`
}

var _ resource.Object = &EdgeDeployment{}

func (in *EdgeDeployment) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *EdgeDeployment) NamespaceScoped() bool {
	return true
}

func (in *EdgeDeployment) New() runtime.Object {
	return &EdgeDeployment{}
}

func (in *EdgeDeployment) NewList() runtime.Object {
	return &EdgeDeploymentList{}
}

func (in *EdgeDeployment) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "management.project-flotta.io",
		Version:  "v1alpha1",
		Resource: "edgedeployments",
	}
}

func (in *EdgeDeployment) IsStorageVersion() bool {
	return true
}

//+kubebuilder:object:root=true

// EdgeDeploymentList contains a list of EdgeDeployment
type EdgeDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeDeployment `json:"items"`
}

var _ resource.ObjectList = &EdgeDeploymentList{}

func (in *EdgeDeploymentList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

func (eds EdgeDeploymentStatus) SubResourceName() string {
	return "status"
}

// EdgeDeployment implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &EdgeDeployment{}

func (in *EdgeDeployment) GetStatus() resource.StatusSubResource {
	return in.Status
}

// EdgeDeploymentStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &EdgeDeploymentStatus{}

func (eds EdgeDeploymentStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*EdgeDeployment).Status = eds
}
