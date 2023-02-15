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

type RoutingPropagatedRouteTableObservation struct {
}

type RoutingPropagatedRouteTableParameters struct {

	// The list of labels to assign to this route table.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A list of Route Table IDs to associated with this Virtual Hub Connection.
	// +kubebuilder:validation:Optional
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`
}

type StaticVnetRouteObservation struct {
}

type StaticVnetRouteParameters struct {

	// A list of CIDR Ranges which should be used as Address Prefixes.
	// +kubebuilder:validation:Optional
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`

	// The name which should be used for this Static Route.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The IP Address which should be used for the Next Hop.
	// +kubebuilder:validation:Optional
	NextHopIPAddress *string `json:"nextHopIpAddress,omitempty" tf:"next_hop_ip_address,omitempty"`
}

type VirtualHubConnectionObservation struct {

	// The ID of the Virtual Hub Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualHubConnectionParameters struct {

	// Should Internet Security be enabled to secure internet traffic? Defaults to false.
	// +kubebuilder:validation:Optional
	InternetSecurityEnabled *bool `json:"internetSecurityEnabled,omitempty" tf:"internet_security_enabled,omitempty"`

	// The ID of the Virtual Network which the Virtual Hub should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RemoteVirtualNetworkID *string `json:"remoteVirtualNetworkId,omitempty" tf:"remote_virtual_network_id,omitempty"`

	// Reference to a VirtualNetwork in network to populate remoteVirtualNetworkId.
	// +kubebuilder:validation:Optional
	RemoteVirtualNetworkIDRef *v1.Reference `json:"remoteVirtualNetworkIdRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork in network to populate remoteVirtualNetworkId.
	// +kubebuilder:validation:Optional
	RemoteVirtualNetworkIDSelector *v1.Selector `json:"remoteVirtualNetworkIdSelector,omitempty" tf:"-"`

	// A routing block as defined below.
	// +kubebuilder:validation:Optional
	Routing []VirtualHubConnectionRoutingParameters `json:"routing,omitempty" tf:"routing,omitempty"`

	// The ID of the Virtual Hub within which this connection should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`

	// Reference to a VirtualHub in network to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDRef *v1.Reference `json:"virtualHubIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHub in network to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDSelector *v1.Selector `json:"virtualHubIdSelector,omitempty" tf:"-"`
}

type VirtualHubConnectionRoutingObservation struct {
}

type VirtualHubConnectionRoutingParameters struct {

	// The ID of the route table associated with this Virtual Hub connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHubRouteTable
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AssociatedRouteTableID *string `json:"associatedRouteTableId,omitempty" tf:"associated_route_table_id,omitempty"`

	// Reference to a VirtualHubRouteTable in network to populate associatedRouteTableId.
	// +kubebuilder:validation:Optional
	AssociatedRouteTableIDRef *v1.Reference `json:"associatedRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHubRouteTable in network to populate associatedRouteTableId.
	// +kubebuilder:validation:Optional
	AssociatedRouteTableIDSelector *v1.Selector `json:"associatedRouteTableIdSelector,omitempty" tf:"-"`

	// A propagated_route_table block as defined below.
	// +kubebuilder:validation:Optional
	PropagatedRouteTable []RoutingPropagatedRouteTableParameters `json:"propagatedRouteTable,omitempty" tf:"propagated_route_table,omitempty"`

	// A static_vnet_route block as defined below.
	// +kubebuilder:validation:Optional
	StaticVnetRoute []StaticVnetRouteParameters `json:"staticVnetRoute,omitempty" tf:"static_vnet_route,omitempty"`
}

// VirtualHubConnectionSpec defines the desired state of VirtualHubConnection
type VirtualHubConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubConnectionParameters `json:"forProvider"`
}

// VirtualHubConnectionStatus defines the observed state of VirtualHubConnection.
type VirtualHubConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubConnection is the Schema for the VirtualHubConnections API. Manages a Connection for a Virtual Hub.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubConnectionSpec   `json:"spec"`
	Status            VirtualHubConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubConnectionList contains a list of VirtualHubConnections
type VirtualHubConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubConnection `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubConnection_Kind             = "VirtualHubConnection"
	VirtualHubConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualHubConnection_Kind}.String()
	VirtualHubConnection_KindAPIVersion   = VirtualHubConnection_Kind + "." + CRDGroupVersion.String()
	VirtualHubConnection_GroupVersionKind = CRDGroupVersion.WithKind(VirtualHubConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubConnection{}, &VirtualHubConnectionList{})
}