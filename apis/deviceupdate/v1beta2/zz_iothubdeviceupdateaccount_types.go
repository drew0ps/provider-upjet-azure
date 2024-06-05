// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IOTHubDeviceUpdateAccountInitParameters struct {

	// An identity block as defined below.
	Identity *IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the IoT Hub Device Update Account should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies whether the public network access is enabled for the IoT Hub Device Update Account. Possible values are true and false. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Sku of the IoT Hub Device Update Account. Possible values are Free and Standard. Defaults to Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags which should be assigned to the IoT Hub Device Update Account.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDeviceUpdateAccountObservation struct {

	// The API host name of the IoT Hub Device Update Account.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the IoT Hub Device Update Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the IoT Hub Device Update Account should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies whether the public network access is enabled for the IoT Hub Device Update Account. Possible values are true and false. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the name of the Resource Group where the IoT Hub Device Update Account should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Sku of the IoT Hub Device Update Account. Possible values are Free and Standard. Defaults to Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags which should be assigned to the IoT Hub Device Update Account.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDeviceUpdateAccountParameters struct {

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the IoT Hub Device Update Account should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies whether the public network access is enabled for the IoT Hub Device Update Account. Possible values are true and false. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Specifies the name of the Resource Group where the IoT Hub Device Update Account should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Sku of the IoT Hub Device Update Account. Possible values are Free and Standard. Defaults to Standard. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags which should be assigned to the IoT Hub Device Update Account.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityInitParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this IoT Hub Device Update Account.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub Device Update Account. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this IoT Hub Device Update Account.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this IoT Hub Device Update Account.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this IoT Hub Device Update Account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub Device Update Account. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this IoT Hub Device Update Account.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub Device Update Account. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// IOTHubDeviceUpdateAccountSpec defines the desired state of IOTHubDeviceUpdateAccount
type IOTHubDeviceUpdateAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubDeviceUpdateAccountParameters `json:"forProvider"`
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
	InitProvider IOTHubDeviceUpdateAccountInitParameters `json:"initProvider,omitempty"`
}

// IOTHubDeviceUpdateAccountStatus defines the observed state of IOTHubDeviceUpdateAccount.
type IOTHubDeviceUpdateAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubDeviceUpdateAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// IOTHubDeviceUpdateAccount is the Schema for the IOTHubDeviceUpdateAccounts API. Manages an IoT Hub Device Update Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubDeviceUpdateAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   IOTHubDeviceUpdateAccountSpec   `json:"spec"`
	Status IOTHubDeviceUpdateAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDeviceUpdateAccountList contains a list of IOTHubDeviceUpdateAccounts
type IOTHubDeviceUpdateAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubDeviceUpdateAccount `json:"items"`
}

// Repository type metadata.
var (
	IOTHubDeviceUpdateAccount_Kind             = "IOTHubDeviceUpdateAccount"
	IOTHubDeviceUpdateAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubDeviceUpdateAccount_Kind}.String()
	IOTHubDeviceUpdateAccount_KindAPIVersion   = IOTHubDeviceUpdateAccount_Kind + "." + CRDGroupVersion.String()
	IOTHubDeviceUpdateAccount_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubDeviceUpdateAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubDeviceUpdateAccount{}, &IOTHubDeviceUpdateAccountList{})
}