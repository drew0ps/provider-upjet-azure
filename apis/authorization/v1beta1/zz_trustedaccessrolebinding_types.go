// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TrustedAccessRoleBindingInitParameters struct {

	// Specifies the Kubernetes Cluster Id within which this Kubernetes Cluster Trusted Access Role Binding should exist. Changing this forces a new Kubernetes Cluster Trusted Access Role Binding to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerservice/v1beta2.KubernetesCluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	KubernetesClusterID *string `json:"kubernetesClusterId,omitempty" tf:"kubernetes_cluster_id,omitempty"`

	// Reference to a KubernetesCluster in containerservice to populate kubernetesClusterId.
	// +kubebuilder:validation:Optional
	KubernetesClusterIDRef *v1.Reference `json:"kubernetesClusterIdRef,omitempty" tf:"-"`

	// Selector for a KubernetesCluster in containerservice to populate kubernetesClusterId.
	// +kubebuilder:validation:Optional
	KubernetesClusterIDSelector *v1.Selector `json:"kubernetesClusterIdSelector,omitempty" tf:"-"`

	// A list of roles to bind, each item is a resource type qualified role name.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// The ARM resource ID of source resource that trusted access is configured for. Changing this forces a new Kubernetes Cluster Trusted Access Role Binding to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/machinelearningservices/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Reference to a Workspace in machinelearningservices to populate sourceResourceId.
	// +kubebuilder:validation:Optional
	SourceResourceIDRef *v1.Reference `json:"sourceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in machinelearningservices to populate sourceResourceId.
	// +kubebuilder:validation:Optional
	SourceResourceIDSelector *v1.Selector `json:"sourceResourceIdSelector,omitempty" tf:"-"`
}

type TrustedAccessRoleBindingObservation struct {

	// The ID of the Kubernetes Cluster Trusted Access Role Binding.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Kubernetes Cluster Id within which this Kubernetes Cluster Trusted Access Role Binding should exist. Changing this forces a new Kubernetes Cluster Trusted Access Role Binding to be created.
	KubernetesClusterID *string `json:"kubernetesClusterId,omitempty" tf:"kubernetes_cluster_id,omitempty"`

	// A list of roles to bind, each item is a resource type qualified role name.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// The ARM resource ID of source resource that trusted access is configured for. Changing this forces a new Kubernetes Cluster Trusted Access Role Binding to be created.
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`
}

type TrustedAccessRoleBindingParameters struct {

	// Specifies the Kubernetes Cluster Id within which this Kubernetes Cluster Trusted Access Role Binding should exist. Changing this forces a new Kubernetes Cluster Trusted Access Role Binding to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerservice/v1beta2.KubernetesCluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KubernetesClusterID *string `json:"kubernetesClusterId,omitempty" tf:"kubernetes_cluster_id,omitempty"`

	// Reference to a KubernetesCluster in containerservice to populate kubernetesClusterId.
	// +kubebuilder:validation:Optional
	KubernetesClusterIDRef *v1.Reference `json:"kubernetesClusterIdRef,omitempty" tf:"-"`

	// Selector for a KubernetesCluster in containerservice to populate kubernetesClusterId.
	// +kubebuilder:validation:Optional
	KubernetesClusterIDSelector *v1.Selector `json:"kubernetesClusterIdSelector,omitempty" tf:"-"`

	// A list of roles to bind, each item is a resource type qualified role name.
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// The ARM resource ID of source resource that trusted access is configured for. Changing this forces a new Kubernetes Cluster Trusted Access Role Binding to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/machinelearningservices/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceResourceID *string `json:"sourceResourceId,omitempty" tf:"source_resource_id,omitempty"`

	// Reference to a Workspace in machinelearningservices to populate sourceResourceId.
	// +kubebuilder:validation:Optional
	SourceResourceIDRef *v1.Reference `json:"sourceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in machinelearningservices to populate sourceResourceId.
	// +kubebuilder:validation:Optional
	SourceResourceIDSelector *v1.Selector `json:"sourceResourceIdSelector,omitempty" tf:"-"`
}

// TrustedAccessRoleBindingSpec defines the desired state of TrustedAccessRoleBinding
type TrustedAccessRoleBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrustedAccessRoleBindingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TrustedAccessRoleBindingInitParameters `json:"initProvider,omitempty"`
}

// TrustedAccessRoleBindingStatus defines the observed state of TrustedAccessRoleBinding.
type TrustedAccessRoleBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrustedAccessRoleBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrustedAccessRoleBinding is the Schema for the TrustedAccessRoleBindings API. Manages a Kubernetes Cluster Trusted Access Role Binding.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TrustedAccessRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roles) || (has(self.initProvider) && has(self.initProvider.roles))",message="spec.forProvider.roles is a required parameter"
	Spec   TrustedAccessRoleBindingSpec   `json:"spec"`
	Status TrustedAccessRoleBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustedAccessRoleBindingList contains a list of TrustedAccessRoleBindings
type TrustedAccessRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustedAccessRoleBinding `json:"items"`
}

// Repository type metadata.
var (
	TrustedAccessRoleBinding_Kind             = "TrustedAccessRoleBinding"
	TrustedAccessRoleBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrustedAccessRoleBinding_Kind}.String()
	TrustedAccessRoleBinding_KindAPIVersion   = TrustedAccessRoleBinding_Kind + "." + CRDGroupVersion.String()
	TrustedAccessRoleBinding_GroupVersionKind = CRDGroupVersion.WithKind(TrustedAccessRoleBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&TrustedAccessRoleBinding{}, &TrustedAccessRoleBindingList{})
}
