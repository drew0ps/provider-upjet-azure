// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this LinuxFunctionAppSlot
func (mg *LinuxFunctionAppSlot) GetTerraformResourceType() string {
	return "azurerm_linux_function_app_slot"
}

// GetConnectionDetailsMapping for this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"auth_settings[*].active_directory[*].client_secret": "spec.forProvider.authSettings[*].activeDirectory[*].clientSecretSecretRef", "auth_settings[*].facebook[*].app_secret": "spec.forProvider.authSettings[*].facebook[*].appSecretSecretRef", "auth_settings[*].github[*].client_secret": "spec.forProvider.authSettings[*].github[*].clientSecretSecretRef", "auth_settings[*].google[*].client_secret": "spec.forProvider.authSettings[*].google[*].clientSecretSecretRef", "auth_settings[*].microsoft[*].client_secret": "spec.forProvider.authSettings[*].microsoft[*].clientSecretSecretRef", "auth_settings[*].twitter[*].consumer_secret": "spec.forProvider.authSettings[*].twitter[*].consumerSecretSecretRef", "backup[*].storage_account_url": "spec.forProvider.backup[*].storageAccountUrlSecretRef", "connection_string[*].value": "spec.forProvider.connectionString[*].valueSecretRef", "custom_domain_verification_id": "status.atProvider.customDomainVerificationId", "site_config[*].application_insights_connection_string": "spec.forProvider.siteConfig[*].applicationInsightsConnectionStringSecretRef", "site_config[*].application_insights_key": "spec.forProvider.siteConfig[*].applicationInsightsKeySecretRef", "site_config[*].application_stack[*].docker[*].registry_password": "spec.forProvider.siteConfig[*].applicationStack[*].docker[*].registryPasswordSecretRef", "site_config[*].application_stack[*].docker[*].registry_username": "spec.forProvider.siteConfig[*].applicationStack[*].docker[*].registryUsernameSecretRef", "site_credential[*]": "status.atProvider.siteCredential[*]", "storage_account[*].access_key": "spec.forProvider.storageAccount[*].accessKeySecretRef", "storage_account_access_key": "spec.forProvider.storageAccountAccessKeySecretRef"}
}

// GetObservation of this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this LinuxFunctionAppSlot
func (tr *LinuxFunctionAppSlot) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this LinuxFunctionAppSlot using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LinuxFunctionAppSlot) LateInitialize(attrs []byte) (bool, error) {
	params := &LinuxFunctionAppSlotParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("KeyVaultReferenceIdentityID"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LinuxFunctionAppSlot) GetTerraformSchemaVersion() int {
	return 0
}
