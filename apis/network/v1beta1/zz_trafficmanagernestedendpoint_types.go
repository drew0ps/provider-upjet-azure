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

type TrafficManagerNestedEndpointCustomHeaderInitParameters struct {

	// The name of the custom header.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TrafficManagerNestedEndpointCustomHeaderObservation struct {

	// The name of the custom header.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TrafficManagerNestedEndpointCustomHeaderParameters struct {

	// The name of the custom header.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of custom header. Applicable for HTTP and HTTPS protocol.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type TrafficManagerNestedEndpointInitParameters struct {

	// One or more custom_header blocks as defined below.
	CustomHeader []TrafficManagerNestedEndpointCustomHeaderInitParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure location of the Endpoint, this must be specified for Profiles using the Performance routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// This argument specifies the minimum number of endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This value must be larger than 0.
	MinimumChildEndpoints *float64 `json:"minimumChildEndpoints,omitempty" tf:"minimum_child_endpoints,omitempty"`

	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type nestedEndpoints and
	MinimumRequiredChildEndpointsIPv4 *float64 `json:"minimumRequiredChildEndpointsIpv4,omitempty" tf:"minimum_required_child_endpoints_ipv4,omitempty"`

	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type nestedEndpoints and
	MinimumRequiredChildEndpointsIPv6 *float64 `json:"minimumRequiredChildEndpointsIpv6,omitempty" tf:"minimum_required_child_endpoints_ipv6,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	Subnet []TrafficManagerNestedEndpointSubnetInitParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerNestedEndpointObservation struct {

	// One or more custom_header blocks as defined below.
	CustomHeader []TrafficManagerNestedEndpointCustomHeaderObservation `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure location of the Endpoint, this must be specified for Profiles using the Performance routing method.
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// The ID of the Nested Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This argument specifies the minimum number of endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This value must be larger than 0.
	MinimumChildEndpoints *float64 `json:"minimumChildEndpoints,omitempty" tf:"minimum_child_endpoints,omitempty"`

	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type nestedEndpoints and
	MinimumRequiredChildEndpointsIPv4 *float64 `json:"minimumRequiredChildEndpointsIpv4,omitempty" tf:"minimum_required_child_endpoints_ipv4,omitempty"`

	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type nestedEndpoints and
	MinimumRequiredChildEndpointsIPv6 *float64 `json:"minimumRequiredChildEndpointsIpv6,omitempty" tf:"minimum_required_child_endpoints_ipv6,omitempty"`

	// Specifies the priority of this Endpoint, this must be specified for Profiles using the Priority traffic routing method. Supports values between 1 and 1000, with no Endpoints sharing the same value. If omitted the value will be computed in order of creation.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the Traffic Manager Profile that this External Endpoint should be created within. Changing this forces a new resource to be created.
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// One or more subnet blocks as defined below. Changing this forces a new resource to be created.
	Subnet []TrafficManagerNestedEndpointSubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The resource id of an Azure resource to target.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerNestedEndpointParameters struct {

	// One or more custom_header blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomHeader []TrafficManagerNestedEndpointCustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// Is the endpoint enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the Azure location of the Endpoint, this must be specified for Profiles using the Performance routing method.
	// +kubebuilder:validation:Optional
	EndpointLocation *string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`

	// A list of Geographic Regions used to distribute traffic, such as WORLD, UK or DE. The same location can't be specified in two endpoints. See the Geographic Hierarchies documentation for more information.
	// +kubebuilder:validation:Optional
	GeoMappings []*string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`

	// This argument specifies the minimum number of endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This value must be larger than 0.
	// +kubebuilder:validation:Optional
	MinimumChildEndpoints *float64 `json:"minimumChildEndpoints,omitempty" tf:"minimum_child_endpoints,omitempty"`

	// This argument specifies the minimum number of IPv4 (DNS record type A) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type nestedEndpoints and
	// +kubebuilder:validation:Optional
	MinimumRequiredChildEndpointsIPv4 *float64 `json:"minimumRequiredChildEndpointsIpv4,omitempty" tf:"minimum_required_child_endpoints_ipv4,omitempty"`

	// This argument specifies the minimum number of IPv6 (DNS record type AAAA) endpoints that must be ‘online’ in the child profile in order for the parent profile to direct traffic to any of the endpoints in that child profile. This argument only applies to Endpoints of type nestedEndpoints and
	// +kubebuilder:validation:Optional
	MinimumRequiredChildEndpointsIPv6 *float64 `json:"minimumRequiredChildEndpointsIpv6,omitempty" tf:"minimum_required_child_endpoints_ipv6,omitempty"`

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
	Subnet []TrafficManagerNestedEndpointSubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The resource id of an Azure resource to target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.TrafficManagerProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a TrafficManagerProfile in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a TrafficManagerProfile in network to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`

	// Specifies how much traffic should be distributed to this endpoint, this must be specified for Profiles using the Weighted traffic routing method. Valid values are between 1 and 1000.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerNestedEndpointSubnetInitParameters struct {

	// The first IP Address in this subnet.
	First *string `json:"first,omitempty" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerNestedEndpointSubnetObservation struct {

	// The first IP Address in this subnet.
	First *string `json:"first,omitempty" tf:"first,omitempty"`

	// The last IP Address in this subnet.
	Last *string `json:"last,omitempty" tf:"last,omitempty"`

	// The block size (number of leading bits in the subnet mask).
	Scope *float64 `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerNestedEndpointSubnetParameters struct {

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

// TrafficManagerNestedEndpointSpec defines the desired state of TrafficManagerNestedEndpoint
type TrafficManagerNestedEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficManagerNestedEndpointParameters `json:"forProvider"`
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
	InitProvider TrafficManagerNestedEndpointInitParameters `json:"initProvider,omitempty"`
}

// TrafficManagerNestedEndpointStatus defines the observed state of TrafficManagerNestedEndpoint.
type TrafficManagerNestedEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficManagerNestedEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficManagerNestedEndpoint is the Schema for the TrafficManagerNestedEndpoints API. Manages a Nested Endpoint within a Traffic Manager Profile.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TrafficManagerNestedEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.minimumChildEndpoints) || (has(self.initProvider) && has(self.initProvider.minimumChildEndpoints))",message="spec.forProvider.minimumChildEndpoints is a required parameter"
	Spec   TrafficManagerNestedEndpointSpec   `json:"spec"`
	Status TrafficManagerNestedEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficManagerNestedEndpointList contains a list of TrafficManagerNestedEndpoints
type TrafficManagerNestedEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficManagerNestedEndpoint `json:"items"`
}

// Repository type metadata.
var (
	TrafficManagerNestedEndpoint_Kind             = "TrafficManagerNestedEndpoint"
	TrafficManagerNestedEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficManagerNestedEndpoint_Kind}.String()
	TrafficManagerNestedEndpoint_KindAPIVersion   = TrafficManagerNestedEndpoint_Kind + "." + CRDGroupVersion.String()
	TrafficManagerNestedEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(TrafficManagerNestedEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficManagerNestedEndpoint{}, &TrafficManagerNestedEndpointList{})
}
