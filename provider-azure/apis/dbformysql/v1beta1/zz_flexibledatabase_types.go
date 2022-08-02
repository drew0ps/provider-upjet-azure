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

type FlexibleDatabaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FlexibleDatabaseParameters struct {

	// +kubebuilder:validation:Required
	Charset *string `json:"charset" tf:"charset,omitempty"`

	// +kubebuilder:validation:Required
	Collation *string `json:"collation" tf:"collation,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=FlexibleServer
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`
}

// FlexibleDatabaseSpec defines the desired state of FlexibleDatabase
type FlexibleDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleDatabaseParameters `json:"forProvider"`
}

// FlexibleDatabaseStatus defines the observed state of FlexibleDatabase.
type FlexibleDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleDatabase is the Schema for the FlexibleDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleDatabaseSpec   `json:"spec"`
	Status            FlexibleDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleDatabaseList contains a list of FlexibleDatabases
type FlexibleDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleDatabase `json:"items"`
}

// Repository type metadata.
var (
	FlexibleDatabase_Kind             = "FlexibleDatabase"
	FlexibleDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleDatabase_Kind}.String()
	FlexibleDatabase_KindAPIVersion   = FlexibleDatabase_Kind + "." + CRDGroupVersion.String()
	FlexibleDatabase_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleDatabase{}, &FlexibleDatabaseList{})
}