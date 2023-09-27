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

type GlobalSchemaInitParameters struct {

	// The description of the schema.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The content type of the Schema. Possible values are xml and json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The string defining the document representing the Schema.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GlobalSchemaObservation struct {

	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The description of the schema.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the API Management API Schema.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The content type of the Schema. Possible values are xml and json.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The string defining the document representing the Schema.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GlobalSchemaParameters struct {

	// The Name of the API Management Service where the API exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The description of the schema.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The content type of the Schema. Possible values are xml and json.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The string defining the document representing the Schema.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// GlobalSchemaSpec defines the desired state of GlobalSchema
type GlobalSchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalSchemaParameters `json:"forProvider"`
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
	InitProvider GlobalSchemaInitParameters `json:"initProvider,omitempty"`
}

// GlobalSchemaStatus defines the observed state of GlobalSchema.
type GlobalSchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalSchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalSchema is the Schema for the GlobalSchemas API. Manages a Global Schema within an API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type GlobalSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   GlobalSchemaSpec   `json:"spec"`
	Status GlobalSchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalSchemaList contains a list of GlobalSchemas
type GlobalSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalSchema `json:"items"`
}

// Repository type metadata.
var (
	GlobalSchema_Kind             = "GlobalSchema"
	GlobalSchema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalSchema_Kind}.String()
	GlobalSchema_KindAPIVersion   = GlobalSchema_Kind + "." + CRDGroupVersion.String()
	GlobalSchema_GroupVersionKind = CRDGroupVersion.WithKind(GlobalSchema_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalSchema{}, &GlobalSchemaList{})
}
