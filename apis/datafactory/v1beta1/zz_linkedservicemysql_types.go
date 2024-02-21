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

type LinkedServiceMySQLInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with MySQL.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The description for the Data Factory Linked Service MySQL.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceMySQLObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with MySQL.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service MySQL.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory MySQL Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceMySQLParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service MySQL.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service MySQL.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with MySQL.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service MySQL.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service MySQL.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service MySQL.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// LinkedServiceMySQLSpec defines the desired state of LinkedServiceMySQL
type LinkedServiceMySQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceMySQLParameters `json:"forProvider"`
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
	InitProvider LinkedServiceMySQLInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceMySQLStatus defines the observed state of LinkedServiceMySQL.
type LinkedServiceMySQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceMySQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkedServiceMySQL is the Schema for the LinkedServiceMySQLs API. Manages a Linked Service (connection) between MySQL and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceMySQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionString) || (has(self.initProvider) && has(self.initProvider.connectionString))",message="spec.forProvider.connectionString is a required parameter"
	Spec   LinkedServiceMySQLSpec   `json:"spec"`
	Status LinkedServiceMySQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceMySQLList contains a list of LinkedServiceMySQLs
type LinkedServiceMySQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceMySQL `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceMySQL_Kind             = "LinkedServiceMySQL"
	LinkedServiceMySQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceMySQL_Kind}.String()
	LinkedServiceMySQL_KindAPIVersion   = LinkedServiceMySQL_Kind + "." + CRDGroupVersion.String()
	LinkedServiceMySQL_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceMySQL_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceMySQL{}, &LinkedServiceMySQLList{})
}
