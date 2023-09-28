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

type LogAnalyticsSolutionInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A plan block as documented below.
	Plan []PlanInitParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Specifies the name of the solution to be deployed. See here for options.Changing this forces a new resource to be created.
	SolutionName *string `json:"solutionName,omitempty" tf:"solution_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LogAnalyticsSolutionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A plan block as documented below.
	Plan []PlanObservation `json:"plan,omitempty" tf:"plan,omitempty"`

	// The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the solution to be deployed. See here for options.Changing this forces a new resource to be created.
	SolutionName *string `json:"solutionName,omitempty" tf:"solution_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
	WorkspaceName *string `json:"workspaceName,omitempty" tf:"workspace_name,omitempty"`

	// The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`
}

type LogAnalyticsSolutionParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A plan block as documented below.
	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and its related workspace can only exist in the same resource group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the solution to be deployed. See here for options.Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SolutionName *string `json:"solutionName,omitempty" tf:"solution_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The full name of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +kubebuilder:validation:Optional
	WorkspaceName *string `json:"workspaceName,omitempty" tf:"workspace_name,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceName.
	// +kubebuilder:validation:Optional
	WorkspaceNameRef *v1.Reference `json:"workspaceNameRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceName.
	// +kubebuilder:validation:Optional
	WorkspaceNameSelector *v1.Selector `json:"workspaceNameSelector,omitempty" tf:"-"`

	// The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDRef *v1.Reference `json:"workspaceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDSelector *v1.Selector `json:"workspaceResourceIdSelector,omitempty" tf:"-"`
}

type PlanInitParameters struct {

	// The product name of the solution. For example OMSGallery/Containers. Changing this forces a new resource to be created.
	Product *string `json:"product,omitempty" tf:"product,omitempty"`

	// A promotion code to be used with the solution. Changing this forces a new resource to be created.
	PromotionCode *string `json:"promotionCode,omitempty" tf:"promotion_code,omitempty"`

	// The publisher of the solution. For example Microsoft. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type PlanObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The product name of the solution. For example OMSGallery/Containers. Changing this forces a new resource to be created.
	Product *string `json:"product,omitempty" tf:"product,omitempty"`

	// A promotion code to be used with the solution. Changing this forces a new resource to be created.
	PromotionCode *string `json:"promotionCode,omitempty" tf:"promotion_code,omitempty"`

	// The publisher of the solution. For example Microsoft. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`
}

type PlanParameters struct {

	// The product name of the solution. For example OMSGallery/Containers. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Product *string `json:"product" tf:"product,omitempty"`

	// A promotion code to be used with the solution. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PromotionCode *string `json:"promotionCode,omitempty" tf:"promotion_code,omitempty"`

	// The publisher of the solution. For example Microsoft. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

// LogAnalyticsSolutionSpec defines the desired state of LogAnalyticsSolution
type LogAnalyticsSolutionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsSolutionParameters `json:"forProvider"`
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
	InitProvider LogAnalyticsSolutionInitParameters `json:"initProvider,omitempty"`
}

// LogAnalyticsSolutionStatus defines the observed state of LogAnalyticsSolution.
type LogAnalyticsSolutionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsSolutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSolution is the Schema for the LogAnalyticsSolutions API. Manages a Log Analytics (formally Operational Insights) Solution.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsSolution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plan) || (has(self.initProvider) && has(self.initProvider.plan))",message="spec.forProvider.plan is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.solutionName) || (has(self.initProvider) && has(self.initProvider.solutionName))",message="spec.forProvider.solutionName is a required parameter"
	Spec   LogAnalyticsSolutionSpec   `json:"spec"`
	Status LogAnalyticsSolutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSolutionList contains a list of LogAnalyticsSolutions
type LogAnalyticsSolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsSolution `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsSolution_Kind             = "LogAnalyticsSolution"
	LogAnalyticsSolution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsSolution_Kind}.String()
	LogAnalyticsSolution_KindAPIVersion   = LogAnalyticsSolution_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsSolution_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsSolution_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsSolution{}, &LogAnalyticsSolutionList{})
}
