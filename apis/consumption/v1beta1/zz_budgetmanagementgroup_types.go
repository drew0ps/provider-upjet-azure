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

type BudgetManagementGroupInitParameters struct {

	// The total amount of cost to track with the budget.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Management Group Consumption Budget.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The name which should be used for this Management Group Consumption Budget. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more notification blocks as defined below.
	Notification []NotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	TimePeriod []TimePeriodInitParameters `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetManagementGroupObservation struct {

	// The total amount of cost to track with the budget.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Management Group Consumption Budget.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the Management Group Consumption Budget.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Management Group. Changing this forces a new resource to be created.
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// The name which should be used for this Management Group Consumption Budget. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more notification blocks as defined below.
	Notification []NotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	TimePeriod []TimePeriodObservation `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetManagementGroupParameters struct {

	// The total amount of cost to track with the budget.
	// +kubebuilder:validation:Optional
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Management Group Consumption Budget.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the Management Group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagementGroupID *string `json:"managementGroupId,omitempty" tf:"management_group_id,omitempty"`

	// Reference to a ManagementGroup in management to populate managementGroupId.
	// +kubebuilder:validation:Optional
	ManagementGroupIDRef *v1.Reference `json:"managementGroupIdRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate managementGroupId.
	// +kubebuilder:validation:Optional
	ManagementGroupIDSelector *v1.Selector `json:"managementGroupIdSelector,omitempty" tf:"-"`

	// The name which should be used for this Management Group Consumption Budget. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more notification blocks as defined below.
	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	// +kubebuilder:validation:Optional
	TimePeriod []TimePeriodParameters `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type DimensionInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type DimensionObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type DimensionParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type FilterInitParameters struct {

	// One or more dimension blocks as defined below to filter the budget on.
	Dimension []DimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	Not []NotInitParameters `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	Tag []FilterTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FilterObservation struct {

	// One or more dimension blocks as defined below to filter the budget on.
	Dimension []DimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	Not []NotObservation `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	Tag []FilterTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FilterParameters struct {

	// One or more dimension blocks as defined below to filter the budget on.
	// +kubebuilder:validation:Optional
	Dimension []DimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	// +kubebuilder:validation:Optional
	Not []NotParameters `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	// +kubebuilder:validation:Optional
	Tag []FilterTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FilterTagInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterTagObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterTagParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type NotDimensionInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NotDimensionObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NotDimensionParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type NotInitParameters struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	Dimension []NotDimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	Tag []TagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotObservation struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	Dimension []NotDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotParameters struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	// +kubebuilder:validation:Optional
	Dimension []NotDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotificationInitParameters struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Should the notification be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type NotificationObservation struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Should the notification be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type NotificationParameters struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	// +kubebuilder:validation:Optional
	ContactEmails []*string `json:"contactEmails" tf:"contact_emails,omitempty"`

	// Should the notification be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type TagInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TagObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TagParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type TimePeriodInitParameters struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new resource to be created.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type TimePeriodObservation struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new resource to be created.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type TimePeriodParameters struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate" tf:"start_date,omitempty"`
}

// BudgetManagementGroupSpec defines the desired state of BudgetManagementGroup
type BudgetManagementGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetManagementGroupParameters `json:"forProvider"`
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
	InitProvider BudgetManagementGroupInitParameters `json:"initProvider,omitempty"`
}

// BudgetManagementGroupStatus defines the observed state of BudgetManagementGroup.
type BudgetManagementGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetManagementGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetManagementGroup is the Schema for the BudgetManagementGroups API. Manages a Consumption Budget for a Management Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BudgetManagementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.amount) || (has(self.initProvider) && has(self.initProvider.amount))",message="spec.forProvider.amount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notification) || (has(self.initProvider) && has(self.initProvider.notification))",message="spec.forProvider.notification is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timePeriod) || (has(self.initProvider) && has(self.initProvider.timePeriod))",message="spec.forProvider.timePeriod is a required parameter"
	Spec   BudgetManagementGroupSpec   `json:"spec"`
	Status BudgetManagementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetManagementGroupList contains a list of BudgetManagementGroups
type BudgetManagementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetManagementGroup `json:"items"`
}

// Repository type metadata.
var (
	BudgetManagementGroup_Kind             = "BudgetManagementGroup"
	BudgetManagementGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BudgetManagementGroup_Kind}.String()
	BudgetManagementGroup_KindAPIVersion   = BudgetManagementGroup_Kind + "." + CRDGroupVersion.String()
	BudgetManagementGroup_GroupVersionKind = CRDGroupVersion.WithKind(BudgetManagementGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&BudgetManagementGroup{}, &BudgetManagementGroupList{})
}
