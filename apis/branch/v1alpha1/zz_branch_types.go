/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BranchObservation struct {
	CanPush *bool `json:"canPush,omitempty" tf:"can_push,omitempty"`

	Commit []CommitObservation `json:"commit,omitempty" tf:"commit,omitempty"`

	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	DeveloperCanMerge *bool `json:"developerCanMerge,omitempty" tf:"developer_can_merge,omitempty"`

	DeveloperCanPush *bool `json:"developerCanPush,omitempty" tf:"developer_can_push,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Merged *bool `json:"merged,omitempty" tf:"merged,omitempty"`

	Protected *bool `json:"protected,omitempty" tf:"protected,omitempty"`

	WebURL *string `json:"webUrl,omitempty" tf:"web_url,omitempty"`
}

type BranchParameters struct {

	// The name for this branch.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID or full path of the project which the branch is created against.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-gitlab/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// The ref which the branch is created from.
	// +kubebuilder:validation:Required
	Ref *string `json:"ref" tf:"ref,omitempty"`
}

type CommitObservation struct {
}

type CommitParameters struct {

	// +kubebuilder:validation:Required
	AuthorEmail *string `json:"authorEmail" tf:"author_email,omitempty"`

	// +kubebuilder:validation:Required
	AuthorName *string `json:"authorName" tf:"author_name,omitempty"`

	// +kubebuilder:validation:Required
	AuthoredDate *string `json:"authoredDate" tf:"authored_date,omitempty"`

	// +kubebuilder:validation:Required
	CommittedDate *string `json:"committedDate" tf:"committed_date,omitempty"`

	// +kubebuilder:validation:Required
	CommitterEmail *string `json:"committerEmail" tf:"committer_email,omitempty"`

	// +kubebuilder:validation:Required
	CommitterName *string `json:"committerName" tf:"committer_name,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Message *string `json:"message" tf:"message,omitempty"`

	// +kubebuilder:validation:Required
	ParentIds []*string `json:"parentIds" tf:"parent_ids,omitempty"`

	// +kubebuilder:validation:Required
	ShortID *string `json:"shortId" tf:"short_id,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

// BranchSpec defines the desired state of Branch
type BranchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BranchParameters `json:"forProvider"`
}

// BranchStatus defines the observed state of Branch.
type BranchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BranchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Branch is the Schema for the Branchs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gitlabjet}
type Branch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BranchSpec   `json:"spec"`
	Status            BranchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BranchList contains a list of Branchs
type BranchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Branch `json:"items"`
}

// Repository type metadata.
var (
	Branch_Kind             = "Branch"
	Branch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Branch_Kind}.String()
	Branch_KindAPIVersion   = Branch_Kind + "." + CRDGroupVersion.String()
	Branch_GroupVersionKind = CRDGroupVersion.WithKind(Branch_Kind)
)

func init() {
	SchemeBuilder.Register(&Branch{}, &BranchList{})
}
