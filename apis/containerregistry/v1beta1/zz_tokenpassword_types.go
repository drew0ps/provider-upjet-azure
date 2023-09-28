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

type Password1InitParameters struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type Password1Observation struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type Password1Parameters struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type Password2InitParameters struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type Password2Observation struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type Password2Parameters struct {

	// The expiration date of the password in RFC3339 format. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`
}

type TokenPasswordInitParameters struct {

	// One password block as defined below.
	Password1 []Password1InitParameters `json:"password1,omitempty" tf:"password1,omitempty"`

	// One password block as defined below.
	Password2 []Password2InitParameters `json:"password2,omitempty" tf:"password2,omitempty"`
}

type TokenPasswordObservation struct {

	// The ID of the Container Registry Token that this Container Registry Token Password resides in. Changing this forces a new Container Registry Token Password to be created.
	ContainerRegistryTokenID *string `json:"containerRegistryTokenId,omitempty" tf:"container_registry_token_id,omitempty"`

	// The ID of the Container Registry Token Password.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One password block as defined below.
	Password1 []Password1Observation `json:"password1,omitempty" tf:"password1,omitempty"`

	// One password block as defined below.
	Password2 []Password2Observation `json:"password2,omitempty" tf:"password2,omitempty"`
}

type TokenPasswordParameters struct {

	// The ID of the Container Registry Token that this Container Registry Token Password resides in. Changing this forces a new Container Registry Token Password to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerregistry/v1beta1.Token
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerRegistryTokenID *string `json:"containerRegistryTokenId,omitempty" tf:"container_registry_token_id,omitempty"`

	// Reference to a Token in containerregistry to populate containerRegistryTokenId.
	// +kubebuilder:validation:Optional
	ContainerRegistryTokenIDRef *v1.Reference `json:"containerRegistryTokenIdRef,omitempty" tf:"-"`

	// Selector for a Token in containerregistry to populate containerRegistryTokenId.
	// +kubebuilder:validation:Optional
	ContainerRegistryTokenIDSelector *v1.Selector `json:"containerRegistryTokenIdSelector,omitempty" tf:"-"`

	// One password block as defined below.
	// +kubebuilder:validation:Optional
	Password1 []Password1Parameters `json:"password1,omitempty" tf:"password1,omitempty"`

	// One password block as defined below.
	// +kubebuilder:validation:Optional
	Password2 []Password2Parameters `json:"password2,omitempty" tf:"password2,omitempty"`
}

// TokenPasswordSpec defines the desired state of TokenPassword
type TokenPasswordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenPasswordParameters `json:"forProvider"`
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
	InitProvider TokenPasswordInitParameters `json:"initProvider,omitempty"`
}

// TokenPasswordStatus defines the observed state of TokenPassword.
type TokenPasswordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenPasswordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TokenPassword is the Schema for the TokenPasswords API. Manages a Container Registry Token Password.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TokenPassword struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.password1) || (has(self.initProvider) && has(self.initProvider.password1))",message="spec.forProvider.password1 is a required parameter"
	Spec   TokenPasswordSpec   `json:"spec"`
	Status TokenPasswordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenPasswordList contains a list of TokenPasswords
type TokenPasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TokenPassword `json:"items"`
}

// Repository type metadata.
var (
	TokenPassword_Kind             = "TokenPassword"
	TokenPassword_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TokenPassword_Kind}.String()
	TokenPassword_KindAPIVersion   = TokenPassword_Kind + "." + CRDGroupVersion.String()
	TokenPassword_GroupVersionKind = CRDGroupVersion.WithKind(TokenPassword_Kind)
)

func init() {
	SchemeBuilder.Register(&TokenPassword{}, &TokenPasswordList{})
}
