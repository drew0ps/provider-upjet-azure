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

type ContentKeyInitParameters struct {

	// ID of Content Key. Changing this forces a new Streaming Locator to be created.
	ContentKeyID *string `json:"contentKeyId,omitempty" tf:"content_key_id,omitempty"`

	// Label of Content Key as specified in the Streaming Policy. Changing this forces a new Streaming Locator to be created.
	LabelReferenceInStreamingPolicy *string `json:"labelReferenceInStreamingPolicy,omitempty" tf:"label_reference_in_streaming_policy,omitempty"`

	// Content Key Policy used by Content Key. Changing this forces a new Streaming Locator to be created.
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// Encryption type of Content Key. Supported values are CommonEncryptionCbcs, CommonEncryptionCenc or EnvelopeEncryption. Changing this forces a new Streaming Locator to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value of Content Key. Changing this forces a new Streaming Locator to be created.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContentKeyObservation struct {

	// ID of Content Key. Changing this forces a new Streaming Locator to be created.
	ContentKeyID *string `json:"contentKeyId,omitempty" tf:"content_key_id,omitempty"`

	// Label of Content Key as specified in the Streaming Policy. Changing this forces a new Streaming Locator to be created.
	LabelReferenceInStreamingPolicy *string `json:"labelReferenceInStreamingPolicy,omitempty" tf:"label_reference_in_streaming_policy,omitempty"`

	// Content Key Policy used by Content Key. Changing this forces a new Streaming Locator to be created.
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// Encryption type of Content Key. Supported values are CommonEncryptionCbcs, CommonEncryptionCenc or EnvelopeEncryption. Changing this forces a new Streaming Locator to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value of Content Key. Changing this forces a new Streaming Locator to be created.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ContentKeyParameters struct {

	// ID of Content Key. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	ContentKeyID *string `json:"contentKeyId,omitempty" tf:"content_key_id,omitempty"`

	// Label of Content Key as specified in the Streaming Policy. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	LabelReferenceInStreamingPolicy *string `json:"labelReferenceInStreamingPolicy,omitempty" tf:"label_reference_in_streaming_policy,omitempty"`

	// Content Key Policy used by Content Key. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// Encryption type of Content Key. Supported values are CommonEncryptionCbcs, CommonEncryptionCenc or EnvelopeEncryption. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value of Content Key. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StreamingLocatorInitParameters struct {

	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaID *string `json:"alternativeMediaId,omitempty" tf:"alternative_media_id,omitempty"`

	// Asset Name. Changing this forces a new Streaming Locator to be created.
	// +crossplane:generate:reference:type=Asset
	AssetName *string `json:"assetName,omitempty" tf:"asset_name,omitempty"`

	// Reference to a Asset to populate assetName.
	// +kubebuilder:validation:Optional
	AssetNameRef *v1.Reference `json:"assetNameRef,omitempty" tf:"-"`

	// Selector for a Asset to populate assetName.
	// +kubebuilder:validation:Optional
	AssetNameSelector *v1.Selector `json:"assetNameSelector,omitempty" tf:"-"`

	// One or more content_key blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKey []ContentKeyInitParameters `json:"contentKey,omitempty" tf:"content_key,omitempty"`

	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty" tf:"default_content_key_policy_name,omitempty"`

	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A list of names of asset or account filters which apply to this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	FilterNames []*string `json:"filterNames,omitempty" tf:"filter_names,omitempty"`

	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorID *string `json:"streamingLocatorId,omitempty" tf:"streaming_locator_id,omitempty"`

	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: Predefined_DownloadOnly, Predefined_ClearStreamingOnly, Predefined_DownloadAndClearStreaming, Predefined_ClearKey, Predefined_MultiDrmCencStreaming and Predefined_MultiDrmStreaming. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName *string `json:"streamingPolicyName,omitempty" tf:"streaming_policy_name,omitempty"`
}

type StreamingLocatorObservation struct {

	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	AlternativeMediaID *string `json:"alternativeMediaId,omitempty" tf:"alternative_media_id,omitempty"`

	// Asset Name. Changing this forces a new Streaming Locator to be created.
	AssetName *string `json:"assetName,omitempty" tf:"asset_name,omitempty"`

	// One or more content_key blocks as defined below. Changing this forces a new Streaming Locator to be created.
	ContentKey []ContentKeyObservation `json:"contentKey,omitempty" tf:"content_key,omitempty"`

	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty" tf:"default_content_key_policy_name,omitempty"`

	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A list of names of asset or account filters which apply to this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	FilterNames []*string `json:"filterNames,omitempty" tf:"filter_names,omitempty"`

	// The ID of the Streaming Locator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	StreamingLocatorID *string `json:"streamingLocatorId,omitempty" tf:"streaming_locator_id,omitempty"`

	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: Predefined_DownloadOnly, Predefined_ClearStreamingOnly, Predefined_DownloadAndClearStreaming, Predefined_ClearKey, Predefined_MultiDrmCencStreaming and Predefined_MultiDrmStreaming. Changing this forces a new Streaming Locator to be created.
	StreamingPolicyName *string `json:"streamingPolicyName,omitempty" tf:"streaming_policy_name,omitempty"`
}

type StreamingLocatorParameters struct {

	// Alternative Media ID of this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	AlternativeMediaID *string `json:"alternativeMediaId,omitempty" tf:"alternative_media_id,omitempty"`

	// Asset Name. Changing this forces a new Streaming Locator to be created.
	// +crossplane:generate:reference:type=Asset
	// +kubebuilder:validation:Optional
	AssetName *string `json:"assetName,omitempty" tf:"asset_name,omitempty"`

	// Reference to a Asset to populate assetName.
	// +kubebuilder:validation:Optional
	AssetNameRef *v1.Reference `json:"assetNameRef,omitempty" tf:"-"`

	// Selector for a Asset to populate assetName.
	// +kubebuilder:validation:Optional
	AssetNameSelector *v1.Selector `json:"assetNameSelector,omitempty" tf:"-"`

	// One or more content_key blocks as defined below. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	ContentKey []ContentKeyParameters `json:"contentKey,omitempty" tf:"content_key,omitempty"`

	// Name of the default Content Key Policy used by this Streaming Locator.Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty" tf:"default_content_key_policy_name,omitempty"`

	// The end time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A list of names of asset or account filters which apply to this Streaming Locator. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	FilterNames []*string `json:"filterNames,omitempty" tf:"filter_names,omitempty"`

	// The Media Services account name. Changing this forces a new Streaming Locator to be created.
	// +crossplane:generate:reference:type=ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the Streaming Locator should exist. Changing this forces a new Streaming Locator to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The start time of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The ID of the Streaming Locator. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	StreamingLocatorID *string `json:"streamingLocatorId,omitempty" tf:"streaming_locator_id,omitempty"`

	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: Predefined_DownloadOnly, Predefined_ClearStreamingOnly, Predefined_DownloadAndClearStreaming, Predefined_ClearKey, Predefined_MultiDrmCencStreaming and Predefined_MultiDrmStreaming. Changing this forces a new Streaming Locator to be created.
	// +kubebuilder:validation:Optional
	StreamingPolicyName *string `json:"streamingPolicyName,omitempty" tf:"streaming_policy_name,omitempty"`
}

// StreamingLocatorSpec defines the desired state of StreamingLocator
type StreamingLocatorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamingLocatorParameters `json:"forProvider"`
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
	InitProvider StreamingLocatorInitParameters `json:"initProvider,omitempty"`
}

// StreamingLocatorStatus defines the observed state of StreamingLocator.
type StreamingLocatorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamingLocatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StreamingLocator is the Schema for the StreamingLocators API. Manages a Media Streaming Locator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StreamingLocator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.streamingPolicyName) || (has(self.initProvider) && has(self.initProvider.streamingPolicyName))",message="spec.forProvider.streamingPolicyName is a required parameter"
	Spec   StreamingLocatorSpec   `json:"spec"`
	Status StreamingLocatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamingLocatorList contains a list of StreamingLocators
type StreamingLocatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamingLocator `json:"items"`
}

// Repository type metadata.
var (
	StreamingLocator_Kind             = "StreamingLocator"
	StreamingLocator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamingLocator_Kind}.String()
	StreamingLocator_KindAPIVersion   = StreamingLocator_Kind + "." + CRDGroupVersion.String()
	StreamingLocator_GroupVersionKind = CRDGroupVersion.WithKind(StreamingLocator_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamingLocator{}, &StreamingLocatorList{})
}
