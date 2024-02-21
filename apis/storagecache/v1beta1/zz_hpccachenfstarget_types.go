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

type HPCCacheNFSTargetInitParameters struct {

	// The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storagecache/v1beta1.HPCCache
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// Reference to a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameRef *v1.Reference `json:"cacheNameRef,omitempty" tf:"-"`

	// Selector for a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameSelector *v1.Selector `json:"cacheNameSelector,omitempty" tf:"-"`

	// Can be specified multiple times to define multiple namespace_junction. Each namespace_juntion block supports fields documented below.
	NamespaceJunction []NamespaceJunctionInitParameters `json:"namespaceJunction,omitempty" tf:"namespace_junction,omitempty"`

	// The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
	TargetHostName *string `json:"targetHostName,omitempty" tf:"target_host_name,omitempty"`

	// The type of usage of the HPC Cache NFS Target. Possible values are: READ_HEAVY_INFREQ, READ_HEAVY_CHECK_180, WRITE_WORKLOAD_15, WRITE_AROUND, WRITE_WORKLOAD_CHECK_30, WRITE_WORKLOAD_CHECK_60 and WRITE_WORKLOAD_CLOUDWS.
	UsageModel *string `json:"usageModel,omitempty" tf:"usage_model,omitempty"`
}

type HPCCacheNFSTargetObservation struct {

	// The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// The ID of the HPC Cache NFS Target.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Can be specified multiple times to define multiple namespace_junction. Each namespace_juntion block supports fields documented below.
	NamespaceJunction []NamespaceJunctionObservation `json:"namespaceJunction,omitempty" tf:"namespace_junction,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
	TargetHostName *string `json:"targetHostName,omitempty" tf:"target_host_name,omitempty"`

	// The type of usage of the HPC Cache NFS Target. Possible values are: READ_HEAVY_INFREQ, READ_HEAVY_CHECK_180, WRITE_WORKLOAD_15, WRITE_AROUND, WRITE_WORKLOAD_CHECK_30, WRITE_WORKLOAD_CHECK_60 and WRITE_WORKLOAD_CLOUDWS.
	UsageModel *string `json:"usageModel,omitempty" tf:"usage_model,omitempty"`
}

type HPCCacheNFSTargetParameters struct {

	// The name HPC Cache, which the HPC Cache NFS Target will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storagecache/v1beta1.HPCCache
	// +kubebuilder:validation:Optional
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// Reference to a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameRef *v1.Reference `json:"cacheNameRef,omitempty" tf:"-"`

	// Selector for a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameSelector *v1.Selector `json:"cacheNameSelector,omitempty" tf:"-"`

	// Can be specified multiple times to define multiple namespace_junction. Each namespace_juntion block supports fields documented below.
	// +kubebuilder:validation:Optional
	NamespaceJunction []NamespaceJunctionParameters `json:"namespaceJunction,omitempty" tf:"namespace_junction,omitempty"`

	// The name of the Resource Group in which to create the HPC Cache NFS Target. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The IP address or fully qualified domain name (FQDN) of the HPC Cache NFS target. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TargetHostName *string `json:"targetHostName,omitempty" tf:"target_host_name,omitempty"`

	// The type of usage of the HPC Cache NFS Target. Possible values are: READ_HEAVY_INFREQ, READ_HEAVY_CHECK_180, WRITE_WORKLOAD_15, WRITE_AROUND, WRITE_WORKLOAD_CHECK_30, WRITE_WORKLOAD_CHECK_60 and WRITE_WORKLOAD_CLOUDWS.
	// +kubebuilder:validation:Optional
	UsageModel *string `json:"usageModel,omitempty" tf:"usage_model,omitempty"`
}

type NamespaceJunctionInitParameters struct {

	// The name of the access policy applied to this target. Defaults to default.
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The NFS export of this NFS target within the HPC Cache NFS Target.
	NFSExport *string `json:"nfsExport,omitempty" tf:"nfs_export,omitempty"`

	// The client-facing file path of this NFS target within the HPC Cache NFS Target.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The relative subdirectory path from the nfs_export to map to the namespace_path. Defaults to "", in which case the whole nfs_export is exported.
	TargetPath *string `json:"targetPath,omitempty" tf:"target_path,omitempty"`
}

type NamespaceJunctionObservation struct {

	// The name of the access policy applied to this target. Defaults to default.
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The NFS export of this NFS target within the HPC Cache NFS Target.
	NFSExport *string `json:"nfsExport,omitempty" tf:"nfs_export,omitempty"`

	// The client-facing file path of this NFS target within the HPC Cache NFS Target.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The relative subdirectory path from the nfs_export to map to the namespace_path. Defaults to "", in which case the whole nfs_export is exported.
	TargetPath *string `json:"targetPath,omitempty" tf:"target_path,omitempty"`
}

type NamespaceJunctionParameters struct {

	// The name of the access policy applied to this target. Defaults to default.
	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The NFS export of this NFS target within the HPC Cache NFS Target.
	// +kubebuilder:validation:Optional
	NFSExport *string `json:"nfsExport" tf:"nfs_export,omitempty"`

	// The client-facing file path of this NFS target within the HPC Cache NFS Target.
	// +kubebuilder:validation:Optional
	NamespacePath *string `json:"namespacePath" tf:"namespace_path,omitempty"`

	// The relative subdirectory path from the nfs_export to map to the namespace_path. Defaults to "", in which case the whole nfs_export is exported.
	// +kubebuilder:validation:Optional
	TargetPath *string `json:"targetPath,omitempty" tf:"target_path,omitempty"`
}

// HPCCacheNFSTargetSpec defines the desired state of HPCCacheNFSTarget
type HPCCacheNFSTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheNFSTargetParameters `json:"forProvider"`
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
	InitProvider HPCCacheNFSTargetInitParameters `json:"initProvider,omitempty"`
}

// HPCCacheNFSTargetStatus defines the observed state of HPCCacheNFSTarget.
type HPCCacheNFSTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheNFSTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HPCCacheNFSTarget is the Schema for the HPCCacheNFSTargets API. Manages a NFS Target within a HPC Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HPCCacheNFSTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.namespaceJunction) || (has(self.initProvider) && has(self.initProvider.namespaceJunction))",message="spec.forProvider.namespaceJunction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetHostName) || (has(self.initProvider) && has(self.initProvider.targetHostName))",message="spec.forProvider.targetHostName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.usageModel) || (has(self.initProvider) && has(self.initProvider.usageModel))",message="spec.forProvider.usageModel is a required parameter"
	Spec   HPCCacheNFSTargetSpec   `json:"spec"`
	Status HPCCacheNFSTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheNFSTargetList contains a list of HPCCacheNFSTargets
type HPCCacheNFSTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCacheNFSTarget `json:"items"`
}

// Repository type metadata.
var (
	HPCCacheNFSTarget_Kind             = "HPCCacheNFSTarget"
	HPCCacheNFSTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCacheNFSTarget_Kind}.String()
	HPCCacheNFSTarget_KindAPIVersion   = HPCCacheNFSTarget_Kind + "." + CRDGroupVersion.String()
	HPCCacheNFSTarget_GroupVersionKind = CRDGroupVersion.WithKind(HPCCacheNFSTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCacheNFSTarget{}, &HPCCacheNFSTargetList{})
}
