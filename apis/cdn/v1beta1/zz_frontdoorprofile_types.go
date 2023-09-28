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

type FrontdoorProfileInitParameters struct {

	// Specifies the maximum response timeout in seconds. Possible values are between 16 and 240 seconds (inclusive). Defaults to 120 seconds.
	ResponseTimeoutSeconds *float64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`

	// Specifies the SKU for this Front Door Profile. Possible values include Standard_AzureFrontDoor and Premium_AzureFrontDoor. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies a mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontdoorProfileObservation struct {

	// The ID of this Front Door Profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The UUID of this Front Door Profile which will be sent in the HTTP Header as the X-Azure-FDID attribute.
	ResourceGUID *string `json:"resourceGuid,omitempty" tf:"resource_guid,omitempty"`

	// The name of the Resource Group where this Front Door Profile should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the maximum response timeout in seconds. Possible values are between 16 and 240 seconds (inclusive). Defaults to 120 seconds.
	ResponseTimeoutSeconds *float64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`

	// Specifies the SKU for this Front Door Profile. Possible values include Standard_AzureFrontDoor and Premium_AzureFrontDoor. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies a mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontdoorProfileParameters struct {

	// The name of the Resource Group where this Front Door Profile should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the maximum response timeout in seconds. Possible values are between 16 and 240 seconds (inclusive). Defaults to 120 seconds.
	// +kubebuilder:validation:Optional
	ResponseTimeoutSeconds *float64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`

	// Specifies the SKU for this Front Door Profile. Possible values include Standard_AzureFrontDoor and Premium_AzureFrontDoor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies a mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// FrontdoorProfileSpec defines the desired state of FrontdoorProfile
type FrontdoorProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorProfileParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FrontdoorProfileInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorProfileStatus defines the observed state of FrontdoorProfile.
type FrontdoorProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorProfile is the Schema for the FrontdoorProfiles API. Manages a Front Door (standard/premium) Profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   FrontdoorProfileSpec   `json:"spec"`
	Status FrontdoorProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorProfileList contains a list of FrontdoorProfiles
type FrontdoorProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorProfile `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorProfile_Kind             = "FrontdoorProfile"
	FrontdoorProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorProfile_Kind}.String()
	FrontdoorProfile_KindAPIVersion   = FrontdoorProfile_Kind + "." + CRDGroupVersion.String()
	FrontdoorProfile_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorProfile{}, &FrontdoorProfileList{})
}
