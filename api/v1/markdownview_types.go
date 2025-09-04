/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MarkdownViewSpec defines the desired state of MarkdownView
type MarkdownViewSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// Markdowns contain the markdown files you want to display.
    // The key indicates the file name and must not overlap with the keys.
    // The value is the content in markdown format.
    //+kubebuilder:validation:Required
    //+kubebuilder:validation:MinProperties=1
    Markdowns map[string]string `json:"markdowns,omitempty"`

    // Replicas is the number of viewers.
    // +kubebuilder:default=1
    // +optional
    Replicas int32 `json:"replicas,omitempty"`

    // ViewerImage is the image name of the viewer.
    // +optional
    ViewerImage string `json:"viewerImage,omitempty"`
}

// MarkdownViewStatus defines the observed state of MarkdownView.
type MarkdownViewStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the MarkdownView resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

const (
    TypeMarkdownViewAvailable = "Available"
    TypeMarkdownViewDegraded  = "Degraded"
)


// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.replicas"
// +kubebuilder:printcolumn:name="Available",type="string",JSONPath=".status.conditions[?(@.type==\"Available\")].status"

// MarkdownView is the Schema for the markdownviews API
type MarkdownView struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of MarkdownView
	// +required
	Spec MarkdownViewSpec `json:"spec"`

	// status defines the observed state of MarkdownView
	// +optional
	Status MarkdownViewStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// MarkdownViewList contains a list of MarkdownView
type MarkdownViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MarkdownView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MarkdownView{}, &MarkdownViewList{})
}
