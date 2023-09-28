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

type TrafficManagerExternalEndpointCustomHeaderInitParameters struct {

	// The name of the custom header.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TrafficManagerExternalEndpointCustomHeaderObservation struct {

	// The name of the custom header.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TrafficManagerExternalEndpointCustomHeaderParameters struct {

	// The name of the custom header.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type TrafficManagerExternalEndpointInitParameters struct {

	// One or more custom_header blocks as defined below.
	CustomHeader []TrafficManagerExternalEndpointCustomHeaderInitParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure location of the Endpoint, this must be specified for Profiles using the Performance routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	Subnet []TrafficManagerExternalEndpointSubnetInitParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The FQDN DNS name of the target.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerExternalEndpointObservation struct {

	// One or more custom_header blocks as defined below.
	CustomHeader []TrafficManagerExternalEndpointCustomHeaderObservation `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure location of the Endpoint, this must be specified for Profiles using the Performance routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// The ID of the External Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the Traffic Manager Profile that this External Endpoint should be created within. Changing this forces a new resource to be created.
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	Subnet []TrafficManagerExternalEndpointSubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The FQDN DNS name of the target.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerExternalEndpointParameters struct {

	// One or more custom_header blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomHeader []TrafficManagerExternalEndpointCustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure location of the Endpoint, this must be specified for Profiles using the Performance routing method.
	// +kubebuilder:validation:Optional
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	// +kubebuilder:validation:Optional
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the Traffic Manager Profile that this External Endpoint should be created within. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.TrafficManagerProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// Reference to a TrafficManagerProfile in network to populate profileId.
	// +kubebuilder:validation:Optional
	ProfileIDRef *v1.Reference `json:"profileIdRef,omitempty" tf:"-"`

	// Selector for a TrafficManagerProfile in network to populate profileId.
	// +kubebuilder:validation:Optional
	ProfileIDSelector *v1.Selector `json:"profileIdSelector,omitempty" tf:"-"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Subnet []TrafficManagerExternalEndpointSubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The FQDN DNS name of the target.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerExternalEndpointSubnetInitParameters struct {

	// The first IP Address in this subnet.
	First *string `json:"first,omitempty" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerExternalEndpointSubnetObservation struct {

	// The first IP Address in this subnet.
	First *string `json:"first,omitempty" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerExternalEndpointSubnetParameters struct {

	// The first IP Address in this subnet.
	// +kubebuilder:validation:Optional
	First *string `json:"first" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	// +kubebuilder:validation:Optional
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	// +kubebuilder:validation:Optional
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

// TrafficManagerExternalEndpointSpec defines the desired state of TrafficManagerExternalEndpoint
type TrafficManagerExternalEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficManagerExternalEndpointParameters `json:"forProvider"`
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
	InitProvider TrafficManagerExternalEndpointInitParameters `json:"initProvider,omitempty"`
}

// TrafficManagerExternalEndpointStatus defines the observed state of TrafficManagerExternalEndpoint.
type TrafficManagerExternalEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficManagerExternalEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficManagerExternalEndpoint is the Schema for the TrafficManagerExternalEndpoints API. Manages an External Endpoint within a Traffic Manager Profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TrafficManagerExternalEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.target) || (has(self.initProvider) && has(self.initProvider.target))",message="spec.forProvider.target is a required parameter"
	Spec   TrafficManagerExternalEndpointSpec   `json:"spec"`
	Status TrafficManagerExternalEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficManagerExternalEndpointList contains a list of TrafficManagerExternalEndpoints
type TrafficManagerExternalEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficManagerExternalEndpoint `json:"items"`
}

// Repository type metadata.
var (
	TrafficManagerExternalEndpoint_Kind             = "TrafficManagerExternalEndpoint"
	TrafficManagerExternalEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficManagerExternalEndpoint_Kind}.String()
	TrafficManagerExternalEndpoint_KindAPIVersion   = TrafficManagerExternalEndpoint_Kind + "." + CRDGroupVersion.String()
	TrafficManagerExternalEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(TrafficManagerExternalEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficManagerExternalEndpoint{}, &TrafficManagerExternalEndpointList{})
}
