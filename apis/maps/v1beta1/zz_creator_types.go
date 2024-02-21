// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type CreatorInitParameters struct {

	// The Azure Region where the Azure Maps Creator should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The storage units to be allocated. Integer values from 1 to 100, inclusive.
	StorageUnits *float64 `json:"storageUnits,omitempty" tf:"storage_units,omitempty"`

	// A mapping of tags which should be assigned to the Azure Maps Creator.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CreatorObservation struct {

	// The ID of the Azure Maps Creator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Azure Maps Creator should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Azure Maps Creator. Changing this forces a new resource to be created.
	MapsAccountID *string `json:"mapsAccountId,omitempty" tf:"maps_account_id,omitempty"`

	// The storage units to be allocated. Integer values from 1 to 100, inclusive.
	StorageUnits *float64 `json:"storageUnits,omitempty" tf:"storage_units,omitempty"`

	// A mapping of tags which should be assigned to the Azure Maps Creator.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CreatorParameters struct {

	// The Azure Region where the Azure Maps Creator should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Azure Maps Creator. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/maps/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MapsAccountID *string `json:"mapsAccountId,omitempty" tf:"maps_account_id,omitempty"`

	// Reference to a Account in maps to populate mapsAccountId.
	// +kubebuilder:validation:Optional
	MapsAccountIDRef *v1.Reference `json:"mapsAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in maps to populate mapsAccountId.
	// +kubebuilder:validation:Optional
	MapsAccountIDSelector *v1.Selector `json:"mapsAccountIdSelector,omitempty" tf:"-"`

	// The storage units to be allocated. Integer values from 1 to 100, inclusive.
	// +kubebuilder:validation:Optional
	StorageUnits *float64 `json:"storageUnits,omitempty" tf:"storage_units,omitempty"`

	// A mapping of tags which should be assigned to the Azure Maps Creator.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CreatorSpec defines the desired state of Creator
type CreatorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CreatorParameters `json:"forProvider"`
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
	InitProvider CreatorInitParameters `json:"initProvider,omitempty"`
}

// CreatorStatus defines the observed state of Creator.
type CreatorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CreatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Creator is the Schema for the Creators API. Manages an Azure Maps Creator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Creator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageUnits) || (has(self.initProvider) && has(self.initProvider.storageUnits))",message="spec.forProvider.storageUnits is a required parameter"
	Spec   CreatorSpec   `json:"spec"`
	Status CreatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CreatorList contains a list of Creators
type CreatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Creator `json:"items"`
}

// Repository type metadata.
var (
	Creator_Kind             = "Creator"
	Creator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Creator_Kind}.String()
	Creator_KindAPIVersion   = Creator_Kind + "." + CRDGroupVersion.String()
	Creator_GroupVersionKind = CRDGroupVersion.WithKind(Creator_Kind)
)

func init() {
	SchemeBuilder.Register(&Creator{}, &CreatorList{})
}
