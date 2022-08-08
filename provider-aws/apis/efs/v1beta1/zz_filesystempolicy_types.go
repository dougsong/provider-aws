/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FileSystemPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FileSystemPolicyParameters struct {

	// +kubebuilder:validation:Optional
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// +crossplane:generate:reference:type=FileSystem
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// FileSystemPolicySpec defines the desired state of FileSystemPolicy
type FileSystemPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileSystemPolicyParameters `json:"forProvider"`
}

// FileSystemPolicyStatus defines the observed state of FileSystemPolicy.
type FileSystemPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileSystemPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystemPolicy is the Schema for the FileSystemPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FileSystemPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSystemPolicySpec   `json:"spec"`
	Status            FileSystemPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystemPolicyList contains a list of FileSystemPolicys
type FileSystemPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystemPolicy `json:"items"`
}

// Repository type metadata.
var (
	FileSystemPolicy_Kind             = "FileSystemPolicy"
	FileSystemPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FileSystemPolicy_Kind}.String()
	FileSystemPolicy_KindAPIVersion   = FileSystemPolicy_Kind + "." + CRDGroupVersion.String()
	FileSystemPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FileSystemPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FileSystemPolicy{}, &FileSystemPolicyList{})
}