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

type CapacityReservationInitParameters struct {

	// A sku block as defined below.
	Sku []SkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Availability Zone for this Capacity Reservation. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type CapacityReservationObservation struct {

	// The ID of the Capacity Reservation Group where the Capacity Reservation exists. Changing this forces a new resource to be created.
	CapacityReservationGroupID *string `json:"capacityReservationGroupId,omitempty" tf:"capacity_reservation_group_id,omitempty"`

	// The ID of the Capacity Reservation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A sku block as defined below.
	Sku []SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Availability Zone for this Capacity Reservation. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type CapacityReservationParameters struct {

	// The ID of the Capacity Reservation Group where the Capacity Reservation exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.CapacityReservationGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CapacityReservationGroupID *string `json:"capacityReservationGroupId,omitempty" tf:"capacity_reservation_group_id,omitempty"`

	// Reference to a CapacityReservationGroup in compute to populate capacityReservationGroupId.
	// +kubebuilder:validation:Optional
	CapacityReservationGroupIDRef *v1.Reference `json:"capacityReservationGroupIdRef,omitempty" tf:"-"`

	// Selector for a CapacityReservationGroup in compute to populate capacityReservationGroupId.
	// +kubebuilder:validation:Optional
	CapacityReservationGroupIDSelector *v1.Selector `json:"capacityReservationGroupIdSelector,omitempty" tf:"-"`

	// A sku block as defined below.
	// +kubebuilder:validation:Optional
	Sku []SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Availability Zone for this Capacity Reservation. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SkuInitParameters struct {

	// Specifies the number of instances to be reserved. It must be a positive integer and not exceed the quota in the subscription.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Name of the sku, such as Standard_F2. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuObservation struct {

	// Specifies the number of instances to be reserved. It must be a positive integer and not exceed the quota in the subscription.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// Name of the sku, such as Standard_F2. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuParameters struct {

	// Specifies the number of instances to be reserved. It must be a positive integer and not exceed the quota in the subscription.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// Name of the sku, such as Standard_F2. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// CapacityReservationSpec defines the desired state of CapacityReservation
type CapacityReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CapacityReservationParameters `json:"forProvider"`
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
	InitProvider CapacityReservationInitParameters `json:"initProvider,omitempty"`
}

// CapacityReservationStatus defines the observed state of CapacityReservation.
type CapacityReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CapacityReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CapacityReservation is the Schema for the CapacityReservations API. Manages a Capacity Reservation within a Capacity Reservation Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CapacityReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   CapacityReservationSpec   `json:"spec"`
	Status CapacityReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CapacityReservationList contains a list of CapacityReservations
type CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CapacityReservation `json:"items"`
}

// Repository type metadata.
var (
	CapacityReservation_Kind             = "CapacityReservation"
	CapacityReservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CapacityReservation_Kind}.String()
	CapacityReservation_KindAPIVersion   = CapacityReservation_Kind + "." + CRDGroupVersion.String()
	CapacityReservation_GroupVersionKind = CRDGroupVersion.WithKind(CapacityReservation_Kind)
)

func init() {
	SchemeBuilder.Register(&CapacityReservation{}, &CapacityReservationList{})
}
