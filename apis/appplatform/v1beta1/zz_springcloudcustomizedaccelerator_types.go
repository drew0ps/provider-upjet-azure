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

type BasicAuthInitParameters struct {

	// Specifies the username of git repository basic auth.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicAuthObservation struct {

	// Specifies the username of git repository basic auth.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicAuthParameters struct {

	// Specifies the password of git repository basic auth.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Specifies the username of git repository basic auth.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type GitRepositoryInitParameters struct {

	// A basic_auth block as defined below. Conflicts with git_repository.0.ssh_auth. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	BasicAuth []BasicAuthInitParameters `json:"basicAuth,omitempty" tf:"basic_auth,omitempty"`

	// Specifies the Git repository branch to be used.
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Specifies the Git repository commit to be used.
	Commit *string `json:"commit,omitempty" tf:"commit,omitempty"`

	// Specifies the Git repository tag to be used.
	GitTag *string `json:"gitTag,omitempty" tf:"git_tag,omitempty"`

	// Specifies the interval for checking for updates to Git or image repository. It should be greater than 10.
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// A ssh_auth block as defined below. Conflicts with git_repository.0.basic_auth. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	SSHAuth []SSHAuthInitParameters `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// Specifies Git repository URL for the accelerator.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type GitRepositoryObservation struct {

	// A basic_auth block as defined below. Conflicts with git_repository.0.ssh_auth. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	BasicAuth []BasicAuthObservation `json:"basicAuth,omitempty" tf:"basic_auth,omitempty"`

	// Specifies the Git repository branch to be used.
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Specifies the Git repository commit to be used.
	Commit *string `json:"commit,omitempty" tf:"commit,omitempty"`

	// Specifies the Git repository tag to be used.
	GitTag *string `json:"gitTag,omitempty" tf:"git_tag,omitempty"`

	// Specifies the interval for checking for updates to Git or image repository. It should be greater than 10.
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// A ssh_auth block as defined below. Conflicts with git_repository.0.basic_auth. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	SSHAuth []SSHAuthObservation `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// Specifies Git repository URL for the accelerator.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type GitRepositoryParameters struct {

	// A basic_auth block as defined below. Conflicts with git_repository.0.ssh_auth. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	// +kubebuilder:validation:Optional
	BasicAuth []BasicAuthParameters `json:"basicAuth,omitempty" tf:"basic_auth,omitempty"`

	// Specifies the Git repository branch to be used.
	// +kubebuilder:validation:Optional
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// Specifies the Git repository commit to be used.
	// +kubebuilder:validation:Optional
	Commit *string `json:"commit,omitempty" tf:"commit,omitempty"`

	// Specifies the Git repository tag to be used.
	// +kubebuilder:validation:Optional
	GitTag *string `json:"gitTag,omitempty" tf:"git_tag,omitempty"`

	// Specifies the interval for checking for updates to Git or image repository. It should be greater than 10.
	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// A ssh_auth block as defined below. Conflicts with git_repository.0.basic_auth. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	// +kubebuilder:validation:Optional
	SSHAuth []SSHAuthParameters `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`

	// Specifies Git repository URL for the accelerator.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

type SSHAuthInitParameters struct {

	// Specifies the SSH Key algorithm of git repository basic auth.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`
}

type SSHAuthObservation struct {

	// Specifies the SSH Key algorithm of git repository basic auth.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`
}

type SSHAuthParameters struct {

	// Specifies the SSH Key algorithm of git repository basic auth.
	// +kubebuilder:validation:Optional
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Specifies the Public SSH Key of git repository basic auth.
	// +kubebuilder:validation:Optional
	HostKeySecretRef *v1.SecretKeySelector `json:"hostKeySecretRef,omitempty" tf:"-"`

	// Specifies the Private SSH Key of git repository basic auth.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`
}

type SpringCloudCustomizedAcceleratorInitParameters struct {

	// Specifies a list of accelerator tags.
	AcceleratorTags []*string `json:"acceleratorTags,omitempty" tf:"accelerator_tags,omitempty"`

	// Specifies the description of the Spring Cloud Customized Accelerator.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the display name of the Spring Cloud Customized Accelerator..
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A git_repository block as defined below.
	GitRepository []GitRepositoryInitParameters `json:"gitRepository,omitempty" tf:"git_repository,omitempty"`

	// Specifies the icon URL of the Spring Cloud Customized Accelerator..
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`
}

type SpringCloudCustomizedAcceleratorObservation struct {

	// Specifies a list of accelerator tags.
	AcceleratorTags []*string `json:"acceleratorTags,omitempty" tf:"accelerator_tags,omitempty"`

	// Specifies the description of the Spring Cloud Customized Accelerator.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the display name of the Spring Cloud Customized Accelerator..
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A git_repository block as defined below.
	GitRepository []GitRepositoryObservation `json:"gitRepository,omitempty" tf:"git_repository,omitempty"`

	// The ID of the Spring Cloud Customized Accelerator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the icon URL of the Spring Cloud Customized Accelerator..
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// The ID of the Spring Cloud Accelerator. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	SpringCloudAcceleratorID *string `json:"springCloudAcceleratorId,omitempty" tf:"spring_cloud_accelerator_id,omitempty"`
}

type SpringCloudCustomizedAcceleratorParameters struct {

	// Specifies a list of accelerator tags.
	// +kubebuilder:validation:Optional
	AcceleratorTags []*string `json:"acceleratorTags,omitempty" tf:"accelerator_tags,omitempty"`

	// Specifies the description of the Spring Cloud Customized Accelerator.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the display name of the Spring Cloud Customized Accelerator..
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A git_repository block as defined below.
	// +kubebuilder:validation:Optional
	GitRepository []GitRepositoryParameters `json:"gitRepository,omitempty" tf:"git_repository,omitempty"`

	// Specifies the icon URL of the Spring Cloud Customized Accelerator..
	// +kubebuilder:validation:Optional
	IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

	// The ID of the Spring Cloud Accelerator. Changing this forces a new Spring Cloud Customized Accelerator to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudAccelerator
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAcceleratorID *string `json:"springCloudAcceleratorId,omitempty" tf:"spring_cloud_accelerator_id,omitempty"`

	// Reference to a SpringCloudAccelerator in appplatform to populate springCloudAcceleratorId.
	// +kubebuilder:validation:Optional
	SpringCloudAcceleratorIDRef *v1.Reference `json:"springCloudAcceleratorIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudAccelerator in appplatform to populate springCloudAcceleratorId.
	// +kubebuilder:validation:Optional
	SpringCloudAcceleratorIDSelector *v1.Selector `json:"springCloudAcceleratorIdSelector,omitempty" tf:"-"`
}

// SpringCloudCustomizedAcceleratorSpec defines the desired state of SpringCloudCustomizedAccelerator
type SpringCloudCustomizedAcceleratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudCustomizedAcceleratorParameters `json:"forProvider"`
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
	InitProvider SpringCloudCustomizedAcceleratorInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudCustomizedAcceleratorStatus defines the observed state of SpringCloudCustomizedAccelerator.
type SpringCloudCustomizedAcceleratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudCustomizedAcceleratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudCustomizedAccelerator is the Schema for the SpringCloudCustomizedAccelerators API. Manages a Spring Cloud Customized Accelerator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudCustomizedAccelerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gitRepository) || (has(self.initProvider) && has(self.initProvider.gitRepository))",message="spec.forProvider.gitRepository is a required parameter"
	Spec   SpringCloudCustomizedAcceleratorSpec   `json:"spec"`
	Status SpringCloudCustomizedAcceleratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudCustomizedAcceleratorList contains a list of SpringCloudCustomizedAccelerators
type SpringCloudCustomizedAcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudCustomizedAccelerator `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudCustomizedAccelerator_Kind             = "SpringCloudCustomizedAccelerator"
	SpringCloudCustomizedAccelerator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudCustomizedAccelerator_Kind}.String()
	SpringCloudCustomizedAccelerator_KindAPIVersion   = SpringCloudCustomizedAccelerator_Kind + "." + CRDGroupVersion.String()
	SpringCloudCustomizedAccelerator_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudCustomizedAccelerator_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudCustomizedAccelerator{}, &SpringCloudCustomizedAcceleratorList{})
}
