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

type ReferenceInputMSSQLInitParameters struct {

	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when refresh_type is Static.
	DeltaSnapshotQuery *string `json:"deltaSnapshotQuery,omitempty" tf:"delta_snapshot_query,omitempty"`

	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery *string `json:"fullSnapshotQuery,omitempty" tf:"full_snapshot_query,omitempty"`

	// The frequency in hh:mm:ss with which the reference data should be retrieved from the MS SQL database e.g. 00:20:00 for every 20 minutes. Must be set when refresh_type is RefreshPeriodicallyWithFull or RefreshPeriodicallyWithDelta.
	RefreshIntervalDuration *string `json:"refreshIntervalDuration,omitempty" tf:"refresh_interval_duration,omitempty"`

	// Defines whether and how the reference data should be refreshed. Accepted values are Static, RefreshPeriodicallyWithFull and RefreshPeriodicallyWithDelta.
	RefreshType *string `json:"refreshType,omitempty" tf:"refresh_type,omitempty"`

	// The name of the table in the Azure SQL database.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The username to connect to the MS SQL database.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ReferenceInputMSSQLObservation struct {

	// The MS SQL database name where the reference data exists.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when refresh_type is Static.
	DeltaSnapshotQuery *string `json:"deltaSnapshotQuery,omitempty" tf:"delta_snapshot_query,omitempty"`

	// The query used to retrieve the reference data from the MS SQL database.
	FullSnapshotQuery *string `json:"fullSnapshotQuery,omitempty" tf:"full_snapshot_query,omitempty"`

	// The ID of the Stream Analytics.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The frequency in hh:mm:ss with which the reference data should be retrieved from the MS SQL database e.g. 00:20:00 for every 20 minutes. Must be set when refresh_type is RefreshPeriodicallyWithFull or RefreshPeriodicallyWithDelta.
	RefreshIntervalDuration *string `json:"refreshIntervalDuration,omitempty" tf:"refresh_interval_duration,omitempty"`

	// Defines whether and how the reference data should be refreshed. Accepted values are Static, RefreshPeriodicallyWithFull and RefreshPeriodicallyWithDelta.
	RefreshType *string `json:"refreshType,omitempty" tf:"refresh_type,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The fully qualified domain name of the MS SQL server.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// The name of the table in the Azure SQL database.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The username to connect to the MS SQL database.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ReferenceInputMSSQLParameters struct {

	// The MS SQL database name where the reference data exists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLDatabase
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Reference to a MSSQLDatabase in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a MSSQLDatabase in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// The query used to retrieve incremental changes in the reference data from the MS SQL database. Cannot be set when refresh_type is Static.
	// +kubebuilder:validation:Optional
	DeltaSnapshotQuery *string `json:"deltaSnapshotQuery,omitempty" tf:"delta_snapshot_query,omitempty"`

	// The query used to retrieve the reference data from the MS SQL database.
	// +kubebuilder:validation:Optional
	FullSnapshotQuery *string `json:"fullSnapshotQuery,omitempty" tf:"full_snapshot_query,omitempty"`

	// The password to connect to the MS SQL database.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The frequency in hh:mm:ss with which the reference data should be retrieved from the MS SQL database e.g. 00:20:00 for every 20 minutes. Must be set when refresh_type is RefreshPeriodicallyWithFull or RefreshPeriodicallyWithDelta.
	// +kubebuilder:validation:Optional
	RefreshIntervalDuration *string `json:"refreshIntervalDuration,omitempty" tf:"refresh_interval_duration,omitempty"`

	// Defines whether and how the reference data should be refreshed. Accepted values are Static, RefreshPeriodicallyWithFull and RefreshPeriodicallyWithDelta.
	// +kubebuilder:validation:Optional
	RefreshType *string `json:"refreshType,omitempty" tf:"refresh_type,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The fully qualified domain name of the MS SQL server.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/sql/v1beta1.MSSQLServer
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("fully_qualified_domain_name",true)
	// +kubebuilder:validation:Optional
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// Reference to a MSSQLServer in sql to populate server.
	// +kubebuilder:validation:Optional
	ServerRef *v1.Reference `json:"serverRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer in sql to populate server.
	// +kubebuilder:validation:Optional
	ServerSelector *v1.Selector `json:"serverSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`

	// The name of the table in the Azure SQL database.
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The username to connect to the MS SQL database.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// ReferenceInputMSSQLSpec defines the desired state of ReferenceInputMSSQL
type ReferenceInputMSSQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReferenceInputMSSQLParameters `json:"forProvider"`
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
	InitProvider ReferenceInputMSSQLInitParameters `json:"initProvider,omitempty"`
}

// ReferenceInputMSSQLStatus defines the observed state of ReferenceInputMSSQL.
type ReferenceInputMSSQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReferenceInputMSSQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceInputMSSQL is the Schema for the ReferenceInputMSSQLs API. Manages a Stream Analytics Reference Input from MS SQL.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ReferenceInputMSSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fullSnapshotQuery) || (has(self.initProvider) && has(self.initProvider.fullSnapshotQuery))",message="spec.forProvider.fullSnapshotQuery is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.refreshType) || (has(self.initProvider) && has(self.initProvider.refreshType))",message="spec.forProvider.refreshType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   ReferenceInputMSSQLSpec   `json:"spec"`
	Status ReferenceInputMSSQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceInputMSSQLList contains a list of ReferenceInputMSSQLs
type ReferenceInputMSSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceInputMSSQL `json:"items"`
}

// Repository type metadata.
var (
	ReferenceInputMSSQL_Kind             = "ReferenceInputMSSQL"
	ReferenceInputMSSQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReferenceInputMSSQL_Kind}.String()
	ReferenceInputMSSQL_KindAPIVersion   = ReferenceInputMSSQL_Kind + "." + CRDGroupVersion.String()
	ReferenceInputMSSQL_GroupVersionKind = CRDGroupVersion.WithKind(ReferenceInputMSSQL_Kind)
)

func init() {
	SchemeBuilder.Register(&ReferenceInputMSSQL{}, &ReferenceInputMSSQLList{})
}
