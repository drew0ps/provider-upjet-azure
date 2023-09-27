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

type VirtualMachineExtensionInitParameters struct {

	// Specifies if the platform deploys the latest minor version update to the type_handler_version specified.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`

	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension?
	AutomaticUpgradeEnabled *bool `json:"automaticUpgradeEnabled,omitempty" tf:"automatic_upgrade_enabled,omitempty"`

	// Should failures from the extension be suppressed? Possible values are true or false. Defaults to false.
	FailureSuppressionEnabled *bool `json:"failureSuppressionEnabled,omitempty" tf:"failure_suppression_enabled,omitempty"`

	// A protected_settings_from_key_vault block as defined below.
	ProtectedSettingsFromKeyVault []VirtualMachineExtensionProtectedSettingsFromKeyVaultInitParameters `json:"protectedSettingsFromKeyVault,omitempty" tf:"protected_settings_from_key_vault,omitempty"`

	// The publisher of the extension, available publishers can be found by using the Azure CLI. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The settings passed to the extension, these are specified as a JSON object in a string.
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of extension, available types for a publisher can be found using the Azure CLI.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion *string `json:"typeHandlerVersion,omitempty" tf:"type_handler_version,omitempty"`
}

type VirtualMachineExtensionObservation struct {

	// Specifies if the platform deploys the latest minor version update to the type_handler_version specified.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`

	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension?
	AutomaticUpgradeEnabled *bool `json:"automaticUpgradeEnabled,omitempty" tf:"automatic_upgrade_enabled,omitempty"`

	// Should failures from the extension be suppressed? Possible values are true or false. Defaults to false.
	FailureSuppressionEnabled *bool `json:"failureSuppressionEnabled,omitempty" tf:"failure_suppression_enabled,omitempty"`

	// The ID of the Virtual Machine Extension.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A protected_settings_from_key_vault block as defined below.
	ProtectedSettingsFromKeyVault []VirtualMachineExtensionProtectedSettingsFromKeyVaultObservation `json:"protectedSettingsFromKeyVault,omitempty" tf:"protected_settings_from_key_vault,omitempty"`

	// The publisher of the extension, available publishers can be found by using the Azure CLI. Changing this forces a new resource to be created.
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The settings passed to the extension, these are specified as a JSON object in a string.
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of extension, available types for a publisher can be found using the Azure CLI.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	TypeHandlerVersion *string `json:"typeHandlerVersion,omitempty" tf:"type_handler_version,omitempty"`

	// The ID of the Virtual Machine. Changing this forces a new resource to be created
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type VirtualMachineExtensionParameters struct {

	// Specifies if the platform deploys the latest minor version update to the type_handler_version specified.
	// +kubebuilder:validation:Optional
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`

	// Should the Extension be automatically updated whenever the Publisher releases a new version of this VM Extension?
	// +kubebuilder:validation:Optional
	AutomaticUpgradeEnabled *bool `json:"automaticUpgradeEnabled,omitempty" tf:"automatic_upgrade_enabled,omitempty"`

	// Should failures from the extension be suppressed? Possible values are true or false. Defaults to false.
	// +kubebuilder:validation:Optional
	FailureSuppressionEnabled *bool `json:"failureSuppressionEnabled,omitempty" tf:"failure_suppression_enabled,omitempty"`

	// A protected_settings_from_key_vault block as defined below.
	// +kubebuilder:validation:Optional
	ProtectedSettingsFromKeyVault []VirtualMachineExtensionProtectedSettingsFromKeyVaultParameters `json:"protectedSettingsFromKeyVault,omitempty" tf:"protected_settings_from_key_vault,omitempty"`

	// The protected_settings passed to the extension, like settings, these are specified as a JSON object in a string.
	// +kubebuilder:validation:Optional
	ProtectedSettingsSecretRef *v1.SecretKeySelector `json:"protectedSettingsSecretRef,omitempty" tf:"-"`

	// The publisher of the extension, available publishers can be found by using the Azure CLI. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// The settings passed to the extension, these are specified as a JSON object in a string.
	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of extension, available types for a publisher can be found using the Azure CLI.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the version of the extension to use, available versions can be found using the Azure CLI.
	// +kubebuilder:validation:Optional
	TypeHandlerVersion *string `json:"typeHandlerVersion,omitempty" tf:"type_handler_version,omitempty"`

	// The ID of the Virtual Machine. Changing this forces a new resource to be created
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachine
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`

	// Reference to a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDRef *v1.Reference `json:"virtualMachineIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDSelector *v1.Selector `json:"virtualMachineIdSelector,omitempty" tf:"-"`
}

type VirtualMachineExtensionProtectedSettingsFromKeyVaultInitParameters struct {

	// The URL to the Key Vault Secret which stores the protected settings.
	SecretURL *string `json:"secretUrl,omitempty" tf:"secret_url,omitempty"`

	// The ID of the source Key Vault.
	SourceVaultID *string `json:"sourceVaultId,omitempty" tf:"source_vault_id,omitempty"`
}

type VirtualMachineExtensionProtectedSettingsFromKeyVaultObservation struct {

	// The URL to the Key Vault Secret which stores the protected settings.
	SecretURL *string `json:"secretUrl,omitempty" tf:"secret_url,omitempty"`

	// The ID of the source Key Vault.
	SourceVaultID *string `json:"sourceVaultId,omitempty" tf:"source_vault_id,omitempty"`
}

type VirtualMachineExtensionProtectedSettingsFromKeyVaultParameters struct {

	// The URL to the Key Vault Secret which stores the protected settings.
	// +kubebuilder:validation:Optional
	SecretURL *string `json:"secretUrl" tf:"secret_url,omitempty"`

	// The ID of the source Key Vault.
	// +kubebuilder:validation:Optional
	SourceVaultID *string `json:"sourceVaultId" tf:"source_vault_id,omitempty"`
}

// VirtualMachineExtensionSpec defines the desired state of VirtualMachineExtension
type VirtualMachineExtensionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualMachineExtensionParameters `json:"forProvider"`
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
	InitProvider VirtualMachineExtensionInitParameters `json:"initProvider,omitempty"`
}

// VirtualMachineExtensionStatus defines the observed state of VirtualMachineExtension.
type VirtualMachineExtensionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualMachineExtensionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineExtension is the Schema for the VirtualMachineExtensions API. Manages a Virtual Machine Extension to provide post deployment configuration and run automated tasks.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualMachineExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publisher) || (has(self.initProvider) && has(self.initProvider.publisher))",message="spec.forProvider.publisher is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.typeHandlerVersion) || (has(self.initProvider) && has(self.initProvider.typeHandlerVersion))",message="spec.forProvider.typeHandlerVersion is a required parameter"
	Spec   VirtualMachineExtensionSpec   `json:"spec"`
	Status VirtualMachineExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineExtensionList contains a list of VirtualMachineExtensions
type VirtualMachineExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineExtension `json:"items"`
}

// Repository type metadata.
var (
	VirtualMachineExtension_Kind             = "VirtualMachineExtension"
	VirtualMachineExtension_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualMachineExtension_Kind}.String()
	VirtualMachineExtension_KindAPIVersion   = VirtualMachineExtension_Kind + "." + CRDGroupVersion.String()
	VirtualMachineExtension_GroupVersionKind = CRDGroupVersion.WithKind(VirtualMachineExtension_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachineExtension{}, &VirtualMachineExtensionList{})
}
