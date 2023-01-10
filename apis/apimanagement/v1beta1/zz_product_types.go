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

type ProductObservation struct {

	// The ID of the API Management Product.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProductParameters struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Do subscribers need to be approved prior to being able to use the Product?
	// +kubebuilder:validation:Optional
	ApprovalRequired *bool `json:"approvalRequired,omitempty" tf:"approval_required,omitempty"`

	// A description of this Product, which may include HTML formatting tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this API Management Product.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Is this Product Published?
	// +kubebuilder:validation:Required
	Published *bool `json:"published" tf:"published,omitempty"`

	// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Is a Subscription required to access API's included in this Product?
	// +kubebuilder:validation:Optional
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required,omitempty"`

	// The number of subscriptions a user can have to this Product at the same time.
	// +kubebuilder:validation:Optional
	SubscriptionsLimit *float64 `json:"subscriptionsLimit,omitempty" tf:"subscriptions_limit,omitempty"`

	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	// +kubebuilder:validation:Optional
	Terms *string `json:"terms,omitempty" tf:"terms,omitempty"`
}

// ProductSpec defines the desired state of Product
type ProductSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProductParameters `json:"forProvider"`
}

// ProductStatus defines the observed state of Product.
type ProductStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProductObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Product is the Schema for the Products API. Manages an API Management Product.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Product struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductSpec   `json:"spec"`
	Status            ProductStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductList contains a list of Products
type ProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Product `json:"items"`
}

// Repository type metadata.
var (
	Product_Kind             = "Product"
	Product_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Product_Kind}.String()
	Product_KindAPIVersion   = Product_Kind + "." + CRDGroupVersion.String()
	Product_GroupVersionKind = CRDGroupVersion.WithKind(Product_Kind)
)

func init() {
	SchemeBuilder.Register(&Product{}, &ProductList{})
}
