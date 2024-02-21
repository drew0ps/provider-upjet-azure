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

type BackupVaultInitParameters struct {

	// Specifies the type of the data store. Possible values are ArchiveStore, SnapshotStore and VaultStore. Changing this forces a new resource to be created.
	DatastoreType *string `json:"datastoreType,omitempty" tf:"datastore_type,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the backup storage redundancy. Possible values are GeoRedundant and LocallyRedundant. Changing this forces a new Backup Vault to be created.
	Redundancy *string `json:"redundancy,omitempty" tf:"redundancy,omitempty"`

	// A mapping of tags which should be assigned to the Backup Vault.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BackupVaultObservation struct {

	// Specifies the type of the data store. Possible values are ArchiveStore, SnapshotStore and VaultStore. Changing this forces a new resource to be created.
	DatastoreType *string `json:"datastoreType,omitempty" tf:"datastore_type,omitempty"`

	// The ID of the Backup Vault.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the backup storage redundancy. Possible values are GeoRedundant and LocallyRedundant. Changing this forces a new Backup Vault to be created.
	Redundancy *string `json:"redundancy,omitempty" tf:"redundancy,omitempty"`

	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the Backup Vault.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BackupVaultParameters struct {

	// Specifies the type of the data store. Possible values are ArchiveStore, SnapshotStore and VaultStore. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DatastoreType *string `json:"datastoreType,omitempty" tf:"datastore_type,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the backup storage redundancy. Possible values are GeoRedundant and LocallyRedundant. Changing this forces a new Backup Vault to be created.
	// +kubebuilder:validation:Optional
	Redundancy *string `json:"redundancy,omitempty" tf:"redundancy,omitempty"`

	// The name of the Resource Group where the Backup Vault should exist. Changing this forces a new Backup Vault to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Backup Vault.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Backup Vault. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// The Principal ID for the Service Principal associated with the Identity of this Backup Vault.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Identity of this Backup Vault.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Backup Vault. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Backup Vault. The only possible value is SystemAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// BackupVaultSpec defines the desired state of BackupVault
type BackupVaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupVaultParameters `json:"forProvider"`
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
	InitProvider BackupVaultInitParameters `json:"initProvider,omitempty"`
}

// BackupVaultStatus defines the observed state of BackupVault.
type BackupVaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupVaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackupVault is the Schema for the BackupVaults API. Manages a Backup Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.datastoreType) || (has(self.initProvider) && has(self.initProvider.datastoreType))",message="spec.forProvider.datastoreType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.redundancy) || (has(self.initProvider) && has(self.initProvider.redundancy))",message="spec.forProvider.redundancy is a required parameter"
	Spec   BackupVaultSpec   `json:"spec"`
	Status BackupVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultList contains a list of BackupVaults
type BackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVault `json:"items"`
}

// Repository type metadata.
var (
	BackupVault_Kind             = "BackupVault"
	BackupVault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupVault_Kind}.String()
	BackupVault_KindAPIVersion   = BackupVault_Kind + "." + CRDGroupVersion.String()
	BackupVault_GroupVersionKind = CRDGroupVersion.WithKind(BackupVault_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupVault{}, &BackupVaultList{})
}
