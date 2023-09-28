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

type FrontdoorOriginGroupInitParameters struct {

	// A health_probe block as defined below.
	HealthProbe []HealthProbeInitParameters `json:"healthProbe,omitempty" tf:"health_probe,omitempty"`

	// A load_balancing block as defined below.
	LoadBalancing []LoadBalancingInitParameters `json:"loadBalancing,omitempty" tf:"load_balancing,omitempty"`

	// Specifies the amount of time which should elapse before shifting traffic to another endpoint when a healthy endpoint becomes unhealthy or a new endpoint is added. Possible values are between 0 and 50 minutes (inclusive). Default is 10 minutes.
	RestoreTrafficTimeToHealedOrNewEndpointInMinutes *float64 `json:"restoreTrafficTimeToHealedOrNewEndpointInMinutes,omitempty" tf:"restore_traffic_time_to_healed_or_new_endpoint_in_minutes,omitempty"`

	// Specifies whether session affinity should be enabled on this host. Defaults to true.
	SessionAffinityEnabled *bool `json:"sessionAffinityEnabled,omitempty" tf:"session_affinity_enabled,omitempty"`
}

type FrontdoorOriginGroupObservation struct {

	// The ID of the Front Door Profile within which this Front Door Origin Group should exist. Changing this forces a new Front Door Origin Group to be created.
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// A health_probe block as defined below.
	HealthProbe []HealthProbeObservation `json:"healthProbe,omitempty" tf:"health_probe,omitempty"`

	// The ID of the Front Door Origin Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A load_balancing block as defined below.
	LoadBalancing []LoadBalancingObservation `json:"loadBalancing,omitempty" tf:"load_balancing,omitempty"`

	// Specifies the amount of time which should elapse before shifting traffic to another endpoint when a healthy endpoint becomes unhealthy or a new endpoint is added. Possible values are between 0 and 50 minutes (inclusive). Default is 10 minutes.
	RestoreTrafficTimeToHealedOrNewEndpointInMinutes *float64 `json:"restoreTrafficTimeToHealedOrNewEndpointInMinutes,omitempty" tf:"restore_traffic_time_to_healed_or_new_endpoint_in_minutes,omitempty"`

	// Specifies whether session affinity should be enabled on this host. Defaults to true.
	SessionAffinityEnabled *bool `json:"sessionAffinityEnabled,omitempty" tf:"session_affinity_enabled,omitempty"`
}

type FrontdoorOriginGroupParameters struct {

	// The ID of the Front Door Profile within which this Front Door Origin Group should exist. Changing this forces a new Front Door Origin Group to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cdn/v1beta1.FrontdoorProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileID *string `json:"cdnFrontdoorProfileId,omitempty" tf:"cdn_frontdoor_profile_id,omitempty"`

	// Reference to a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDRef *v1.Reference `json:"cdnFrontdoorProfileIdRef,omitempty" tf:"-"`

	// Selector for a FrontdoorProfile in cdn to populate cdnFrontdoorProfileId.
	// +kubebuilder:validation:Optional
	CdnFrontdoorProfileIDSelector *v1.Selector `json:"cdnFrontdoorProfileIdSelector,omitempty" tf:"-"`

	// A health_probe block as defined below.
	// +kubebuilder:validation:Optional
	HealthProbe []HealthProbeParameters `json:"healthProbe,omitempty" tf:"health_probe,omitempty"`

	// A load_balancing block as defined below.
	// +kubebuilder:validation:Optional
	LoadBalancing []LoadBalancingParameters `json:"loadBalancing,omitempty" tf:"load_balancing,omitempty"`

	// Specifies the amount of time which should elapse before shifting traffic to another endpoint when a healthy endpoint becomes unhealthy or a new endpoint is added. Possible values are between 0 and 50 minutes (inclusive). Default is 10 minutes.
	// +kubebuilder:validation:Optional
	RestoreTrafficTimeToHealedOrNewEndpointInMinutes *float64 `json:"restoreTrafficTimeToHealedOrNewEndpointInMinutes,omitempty" tf:"restore_traffic_time_to_healed_or_new_endpoint_in_minutes,omitempty"`

	// Specifies whether session affinity should be enabled on this host. Defaults to true.
	// +kubebuilder:validation:Optional
	SessionAffinityEnabled *bool `json:"sessionAffinityEnabled,omitempty" tf:"session_affinity_enabled,omitempty"`
}

type HealthProbeInitParameters struct {

	// Specifies the number of seconds between health probes. Possible values are between 5 and 31536000 seconds (inclusive).
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// Specifies the path relative to the origin that is used to determine the health of the origin. Defaults to /.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the protocol to use for health probe. Possible values are Http and Https.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the type of health probe request that is made. Possible values are GET and HEAD. Defaults to HEAD.
	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`
}

type HealthProbeObservation struct {

	// Specifies the number of seconds between health probes. Possible values are between 5 and 31536000 seconds (inclusive).
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// Specifies the path relative to the origin that is used to determine the health of the origin. Defaults to /.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the protocol to use for health probe. Possible values are Http and Https.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies the type of health probe request that is made. Possible values are GET and HEAD. Defaults to HEAD.
	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`
}

type HealthProbeParameters struct {

	// Specifies the number of seconds between health probes. Possible values are between 5 and 31536000 seconds (inclusive).
	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds" tf:"interval_in_seconds,omitempty"`

	// Specifies the path relative to the origin that is used to determine the health of the origin. Defaults to /.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Specifies the protocol to use for health probe. Possible values are Http and Https.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Specifies the type of health probe request that is made. Possible values are GET and HEAD. Defaults to HEAD.
	// +kubebuilder:validation:Optional
	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`
}

type LoadBalancingInitParameters struct {

	// Specifies the additional latency in milliseconds for probes to fall into the lowest latency bucket. Possible values are between 0 and 1000 milliseconds (inclusive). Defaults to 50.
	AdditionalLatencyInMilliseconds *float64 `json:"additionalLatencyInMilliseconds,omitempty" tf:"additional_latency_in_milliseconds,omitempty"`

	// Specifies the number of samples to consider for load balancing decisions. Possible values are between 0 and 255 (inclusive). Defaults to 4.
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`

	// Specifies the number of samples within the sample period that must succeed. Possible values are between 0 and 255 (inclusive). Defaults to 3.
	SuccessfulSamplesRequired *float64 `json:"successfulSamplesRequired,omitempty" tf:"successful_samples_required,omitempty"`
}

type LoadBalancingObservation struct {

	// Specifies the additional latency in milliseconds for probes to fall into the lowest latency bucket. Possible values are between 0 and 1000 milliseconds (inclusive). Defaults to 50.
	AdditionalLatencyInMilliseconds *float64 `json:"additionalLatencyInMilliseconds,omitempty" tf:"additional_latency_in_milliseconds,omitempty"`

	// Specifies the number of samples to consider for load balancing decisions. Possible values are between 0 and 255 (inclusive). Defaults to 4.
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`

	// Specifies the number of samples within the sample period that must succeed. Possible values are between 0 and 255 (inclusive). Defaults to 3.
	SuccessfulSamplesRequired *float64 `json:"successfulSamplesRequired,omitempty" tf:"successful_samples_required,omitempty"`
}

type LoadBalancingParameters struct {

	// Specifies the additional latency in milliseconds for probes to fall into the lowest latency bucket. Possible values are between 0 and 1000 milliseconds (inclusive). Defaults to 50.
	// +kubebuilder:validation:Optional
	AdditionalLatencyInMilliseconds *float64 `json:"additionalLatencyInMilliseconds,omitempty" tf:"additional_latency_in_milliseconds,omitempty"`

	// Specifies the number of samples to consider for load balancing decisions. Possible values are between 0 and 255 (inclusive). Defaults to 4.
	// +kubebuilder:validation:Optional
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`

	// Specifies the number of samples within the sample period that must succeed. Possible values are between 0 and 255 (inclusive). Defaults to 3.
	// +kubebuilder:validation:Optional
	SuccessfulSamplesRequired *float64 `json:"successfulSamplesRequired,omitempty" tf:"successful_samples_required,omitempty"`
}

// FrontdoorOriginGroupSpec defines the desired state of FrontdoorOriginGroup
type FrontdoorOriginGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorOriginGroupParameters `json:"forProvider"`
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
	InitProvider FrontdoorOriginGroupInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorOriginGroupStatus defines the observed state of FrontdoorOriginGroup.
type FrontdoorOriginGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorOriginGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorOriginGroup is the Schema for the FrontdoorOriginGroups API. Manages a Front Door (standard/premium) Origin Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorOriginGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.loadBalancing) || (has(self.initProvider) && has(self.initProvider.loadBalancing))",message="spec.forProvider.loadBalancing is a required parameter"
	Spec   FrontdoorOriginGroupSpec   `json:"spec"`
	Status FrontdoorOriginGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorOriginGroupList contains a list of FrontdoorOriginGroups
type FrontdoorOriginGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorOriginGroup `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorOriginGroup_Kind             = "FrontdoorOriginGroup"
	FrontdoorOriginGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorOriginGroup_Kind}.String()
	FrontdoorOriginGroup_KindAPIVersion   = FrontdoorOriginGroup_Kind + "." + CRDGroupVersion.String()
	FrontdoorOriginGroup_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorOriginGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorOriginGroup{}, &FrontdoorOriginGroupList{})
}
