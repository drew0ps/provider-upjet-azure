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

type SubscriptionPolicyAssignmentIdentityInitParameters struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned or UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriptionPolicyAssignmentIdentityObservation struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID of the Policy Assignment for this Subscription.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID of the Policy Assignment for this Subscription.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned or UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubscriptionPolicyAssignmentIdentityParameters struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned or UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SubscriptionPolicyAssignmentInitParameters struct {

	// A description which should be used for this Policy Assignment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not? Defaults to true.
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// An identity block as defined below.
	Identity []SubscriptionPolicyAssignmentIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	NonComplianceMessage []SubscriptionPolicyAssignmentNonComplianceMessageInitParameters `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// One or more overrides blocks as defined below. More detail about overrides and resource_selectors see policy assignment structure
	Overrides []SubscriptionPolicyAssignmentOverridesInitParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// One or more resource_selectors blocks as defined below to filter polices by resource properties.
	ResourceSelectors []SubscriptionPolicyAssignmentResourceSelectorsInitParameters `json:"resourceSelectors,omitempty" tf:"resource_selectors,omitempty"`

	// The ID of the Subscription where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyAssignmentNonComplianceMessageInitParameters struct {

	// The non-compliance message text. When assigning policy sets (initiatives), unless policy_definition_reference_id is specified then this message will be the default for all policies.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type SubscriptionPolicyAssignmentNonComplianceMessageObservation struct {

	// The non-compliance message text. When assigning policy sets (initiatives), unless policy_definition_reference_id is specified then this message will be the default for all policies.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type SubscriptionPolicyAssignmentNonComplianceMessageParameters struct {

	// The non-compliance message text. When assigning policy sets (initiatives), unless policy_definition_reference_id is specified then this message will be the default for all policies.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// When assigning policy sets (initiatives), this is the ID of the policy definition that the non-compliance message applies to.
	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type SubscriptionPolicyAssignmentObservation struct {

	// A description which should be used for this Policy Assignment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not? Defaults to true.
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// The ID of the Subscription Policy Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []SubscriptionPolicyAssignmentIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	NonComplianceMessage []SubscriptionPolicyAssignmentNonComplianceMessageObservation `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// One or more overrides blocks as defined below. More detail about overrides and resource_selectors see policy assignment structure
	Overrides []SubscriptionPolicyAssignmentOverridesObservation `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// One or more resource_selectors blocks as defined below to filter polices by resource properties.
	ResourceSelectors []SubscriptionPolicyAssignmentResourceSelectorsObservation `json:"resourceSelectors,omitempty" tf:"resource_selectors,omitempty"`

	// The ID of the Subscription where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyAssignmentOverridesInitParameters struct {

	// One or more override_selector as defined below.
	Selectors []SubscriptionPolicyAssignmentOverridesSelectorsInitParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`

	// Specifies the value to override the policy property. Possible values for policyEffect override listed policy effects.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SubscriptionPolicyAssignmentOverridesObservation struct {

	// One or more override_selector as defined below.
	Selectors []SubscriptionPolicyAssignmentOverridesSelectorsObservation `json:"selectors,omitempty" tf:"selectors,omitempty"`

	// Specifies the value to override the policy property. Possible values for policyEffect override listed policy effects.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type SubscriptionPolicyAssignmentOverridesParameters struct {

	// One or more override_selector as defined below.
	// +kubebuilder:validation:Optional
	Selectors []SubscriptionPolicyAssignmentOverridesSelectorsParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`

	// Specifies the value to override the policy property. Possible values for policyEffect override listed policy effects.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type SubscriptionPolicyAssignmentOverridesSelectorsInitParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SubscriptionPolicyAssignmentOverridesSelectorsObservation struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SubscriptionPolicyAssignmentOverridesSelectorsParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SubscriptionPolicyAssignmentParameters struct {

	// A description which should be used for this Policy Assignment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not? Defaults to true.
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []SubscriptionPolicyAssignmentIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	// +kubebuilder:validation:Optional
	NonComplianceMessage []SubscriptionPolicyAssignmentNonComplianceMessageParameters `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes (for example a Subscription, or a Resource Group) within this Management Group which are excluded from this Policy.
	// +kubebuilder:validation:Optional
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// One or more overrides blocks as defined below. More detail about overrides and resource_selectors see policy assignment structure
	// +kubebuilder:validation:Optional
	Overrides []SubscriptionPolicyAssignmentOverridesParameters `json:"overrides,omitempty" tf:"overrides,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/authorization/v1beta1.PolicyDefinition
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// Reference to a PolicyDefinition in authorization to populate policyDefinitionId.
	// +kubebuilder:validation:Optional
	PolicyDefinitionIDRef *v1.Reference `json:"policyDefinitionIdRef,omitempty" tf:"-"`

	// Selector for a PolicyDefinition in authorization to populate policyDefinitionId.
	// +kubebuilder:validation:Optional
	PolicyDefinitionIDSelector *v1.Selector `json:"policyDefinitionIdSelector,omitempty" tf:"-"`

	// One or more resource_selectors blocks as defined below to filter polices by resource properties.
	// +kubebuilder:validation:Optional
	ResourceSelectors []SubscriptionPolicyAssignmentResourceSelectorsParameters `json:"resourceSelectors,omitempty" tf:"resource_selectors,omitempty"`

	// The ID of the Subscription where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`
}

type SubscriptionPolicyAssignmentResourceSelectorsInitParameters struct {

	// Specifies a name for the resource selector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more resource_selector block as defined below.
	Selectors []SubscriptionPolicyAssignmentResourceSelectorsSelectorsInitParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`
}

type SubscriptionPolicyAssignmentResourceSelectorsObservation struct {

	// Specifies a name for the resource selector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more resource_selector block as defined below.
	Selectors []SubscriptionPolicyAssignmentResourceSelectorsSelectorsObservation `json:"selectors,omitempty" tf:"selectors,omitempty"`
}

type SubscriptionPolicyAssignmentResourceSelectorsParameters struct {

	// Specifies a name for the resource selector.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more resource_selector block as defined below.
	// +kubebuilder:validation:Optional
	Selectors []SubscriptionPolicyAssignmentResourceSelectorsSelectorsParameters `json:"selectors" tf:"selectors,omitempty"`
}

type SubscriptionPolicyAssignmentResourceSelectorsSelectorsInitParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SubscriptionPolicyAssignmentResourceSelectorsSelectorsObservation struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

type SubscriptionPolicyAssignmentResourceSelectorsSelectorsParameters struct {

	// The list of allowed values for the specified kind. Cannot be used with not_in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	In []*string `json:"in,omitempty" tf:"in,omitempty"`

	// Specifies which characteristic will narrow down the set of evaluated resources. Possible values are resourceLocation,  resourceType and resourceWithoutLocation.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// The list of not-allowed values for the specified kind. Cannot be used with in. Can contain up to 50 values.
	// +kubebuilder:validation:Optional
	NotIn []*string `json:"notIn,omitempty" tf:"not_in,omitempty"`
}

// SubscriptionPolicyAssignmentSpec defines the desired state of SubscriptionPolicyAssignment
type SubscriptionPolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionPolicyAssignmentParameters `json:"forProvider"`
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
	InitProvider SubscriptionPolicyAssignmentInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionPolicyAssignmentStatus defines the observed state of SubscriptionPolicyAssignment.
type SubscriptionPolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionPolicyAssignment is the Schema for the SubscriptionPolicyAssignments API. Manages a Subscription Policy Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubscriptionPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionId) || (has(self.initProvider) && has(self.initProvider.subscriptionId))",message="spec.forProvider.subscriptionId is a required parameter"
	Spec   SubscriptionPolicyAssignmentSpec   `json:"spec"`
	Status SubscriptionPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionPolicyAssignmentList contains a list of SubscriptionPolicyAssignments
type SubscriptionPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionPolicyAssignment_Kind             = "SubscriptionPolicyAssignment"
	SubscriptionPolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionPolicyAssignment_Kind}.String()
	SubscriptionPolicyAssignment_KindAPIVersion   = SubscriptionPolicyAssignment_Kind + "." + CRDGroupVersion.String()
	SubscriptionPolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionPolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionPolicyAssignment{}, &SubscriptionPolicyAssignmentList{})
}
