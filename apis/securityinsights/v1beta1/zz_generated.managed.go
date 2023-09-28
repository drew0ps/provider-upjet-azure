/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelAlertRuleFusion.
func (mg *SentinelAlertRuleFusion) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelAlertRuleMSSecurityIncident.
func (mg *SentinelAlertRuleMSSecurityIncident) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelAlertRuleMachineLearningBehaviorAnalytics.
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelAutomationRule.
func (mg *SentinelAutomationRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelDataConnectorIOT.
func (mg *SentinelDataConnectorIOT) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelLogAnalyticsWorkspaceOnboarding.
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SentinelWatchlist.
func (mg *SentinelWatchlist) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SentinelWatchlist.
func (mg *SentinelWatchlist) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SentinelWatchlist.
func (mg *SentinelWatchlist) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SentinelWatchlist.
func (mg *SentinelWatchlist) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SentinelWatchlist.
func (mg *SentinelWatchlist) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SentinelWatchlist.
func (mg *SentinelWatchlist) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SentinelWatchlist.
func (mg *SentinelWatchlist) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SentinelWatchlist.
func (mg *SentinelWatchlist) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SentinelWatchlist.
func (mg *SentinelWatchlist) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SentinelWatchlist.
func (mg *SentinelWatchlist) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SentinelWatchlist.
func (mg *SentinelWatchlist) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SentinelWatchlist.
func (mg *SentinelWatchlist) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
