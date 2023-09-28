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

type StandardEnvironmentInitParameters struct {

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime *string `json:"dataRetentionTime,omitempty" tf:"data_retention_time,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(_).The first part is the name, valid values include: S1 and S2. The second part is the capacity (e.g. the number of deployed units of the sku), which must be a positive integer (e.g. S1_1). Possible values are S1_1, S1_2, S1_3, S1_4, S1_5, S1_6, S1_7, S1_8, S1_9, S1_10, S2_1, S2_2, S2_3, S2_4, S2_5, S2_6, S2_7, S2_8, S2_9 and S2_10. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include PauseIngress and PurgeOldData. Defaults to PurgeOldData.
	StorageLimitExceededBehavior *string `json:"storageLimitExceededBehavior,omitempty" tf:"storage_limit_exceeded_behavior,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StandardEnvironmentObservation struct {

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	DataRetentionTime *string `json:"dataRetentionTime,omitempty" tf:"data_retention_time,omitempty"`

	// The ID of the IoT Time Series Insights Standard Environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(_).The first part is the name, valid values include: S1 and S2. The second part is the capacity (e.g. the number of deployed units of the sku), which must be a positive integer (e.g. S1_1). Possible values are S1_1, S1_2, S1_3, S1_4, S1_5, S1_6, S1_7, S1_8, S1_9, S1_10, S2_1, S2_2, S2_3, S2_4, S2_5, S2_6, S2_7, S2_8, S2_9 and S2_10. Changing this forces a new resource to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include PauseIngress and PurgeOldData. Defaults to PurgeOldData.
	StorageLimitExceededBehavior *string `json:"storageLimitExceededBehavior,omitempty" tf:"storage_limit_exceeded_behavior,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StandardEnvironmentParameters struct {

	// Specifies the ISO8601 timespan specifying the minimum number of days the environment's events will be available for query. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataRetentionTime *string `json:"dataRetentionTime,omitempty" tf:"data_retention_time,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the event property which will be used to partition data. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The name of the resource group in which to create the Azure IoT Time Series Insights Standard Environment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the SKU Name for this IoT Time Series Insights Standard Environment. It is string consisting of two parts separated by an underscore(_).The first part is the name, valid values include: S1 and S2. The second part is the capacity (e.g. the number of deployed units of the sku), which must be a positive integer (e.g. S1_1). Possible values are S1_1, S1_2, S1_3, S1_4, S1_5, S1_6, S1_7, S1_8, S1_9, S1_10, S2_1, S2_2, S2_3, S2_4, S2_5, S2_6, S2_7, S2_8, S2_9 and S2_10. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Specifies the behaviour the IoT Time Series Insights service should take when the environment's capacity has been exceeded. Valid values include PauseIngress and PurgeOldData. Defaults to PurgeOldData.
	// +kubebuilder:validation:Optional
	StorageLimitExceededBehavior *string `json:"storageLimitExceededBehavior,omitempty" tf:"storage_limit_exceeded_behavior,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StandardEnvironmentSpec defines the desired state of StandardEnvironment
type StandardEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StandardEnvironmentParameters `json:"forProvider"`
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
	InitProvider StandardEnvironmentInitParameters `json:"initProvider,omitempty"`
}

// StandardEnvironmentStatus defines the observed state of StandardEnvironment.
type StandardEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StandardEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StandardEnvironment is the Schema for the StandardEnvironments API. Manages an Azure IoT Time Series Insights Standard Environment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StandardEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataRetentionTime) || (has(self.initProvider) && has(self.initProvider.dataRetentionTime))",message="spec.forProvider.dataRetentionTime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   StandardEnvironmentSpec   `json:"spec"`
	Status StandardEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StandardEnvironmentList contains a list of StandardEnvironments
type StandardEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StandardEnvironment `json:"items"`
}

// Repository type metadata.
var (
	StandardEnvironment_Kind             = "StandardEnvironment"
	StandardEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StandardEnvironment_Kind}.String()
	StandardEnvironment_KindAPIVersion   = StandardEnvironment_Kind + "." + CRDGroupVersion.String()
	StandardEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(StandardEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&StandardEnvironment{}, &StandardEnvironmentList{})
}
