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

type CustomerManagedKeyInitParameters struct {

	// The ID of the geo backup Key Vault Key. It can't cross region and need Customer Managed Key in same region as geo backup.
	GeoBackupKeyVaultKeyID *string `json:"geoBackupKeyVaultKeyId,omitempty" tf:"geo_backup_key_vault_key_id,omitempty"`

	// The geo backup user managed identity id for a Customer Managed Key. Should be added with identity_ids. It can't cross region and need identity in same region as geo backup.
	GeoBackupUserAssignedIdentityID *string `json:"geoBackupUserAssignedIdentityId,omitempty" tf:"geo_backup_user_assigned_identity_id,omitempty"`

	// The ID of the Key Vault Key.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the primary user managed identity id for a Customer Managed Key. Should be added with identity_ids.
	PrimaryUserAssignedIdentityID *string `json:"primaryUserAssignedIdentityId,omitempty" tf:"primary_user_assigned_identity_id,omitempty"`
}

type CustomerManagedKeyObservation struct {

	// The ID of the geo backup Key Vault Key. It can't cross region and need Customer Managed Key in same region as geo backup.
	GeoBackupKeyVaultKeyID *string `json:"geoBackupKeyVaultKeyId,omitempty" tf:"geo_backup_key_vault_key_id,omitempty"`

	// The geo backup user managed identity id for a Customer Managed Key. Should be added with identity_ids. It can't cross region and need identity in same region as geo backup.
	GeoBackupUserAssignedIdentityID *string `json:"geoBackupUserAssignedIdentityId,omitempty" tf:"geo_backup_user_assigned_identity_id,omitempty"`

	// The ID of the Key Vault Key.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the primary user managed identity id for a Customer Managed Key. Should be added with identity_ids.
	PrimaryUserAssignedIdentityID *string `json:"primaryUserAssignedIdentityId,omitempty" tf:"primary_user_assigned_identity_id,omitempty"`
}

type CustomerManagedKeyParameters struct {

	// The ID of the geo backup Key Vault Key. It can't cross region and need Customer Managed Key in same region as geo backup.
	// +kubebuilder:validation:Optional
	GeoBackupKeyVaultKeyID *string `json:"geoBackupKeyVaultKeyId,omitempty" tf:"geo_backup_key_vault_key_id,omitempty"`

	// The geo backup user managed identity id for a Customer Managed Key. Should be added with identity_ids. It can't cross region and need identity in same region as geo backup.
	// +kubebuilder:validation:Optional
	GeoBackupUserAssignedIdentityID *string `json:"geoBackupUserAssignedIdentityId,omitempty" tf:"geo_backup_user_assigned_identity_id,omitempty"`

	// The ID of the Key Vault Key.
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Specifies the primary user managed identity id for a Customer Managed Key. Should be added with identity_ids.
	// +kubebuilder:validation:Optional
	PrimaryUserAssignedIdentityID *string `json:"primaryUserAssignedIdentityId,omitempty" tf:"primary_user_assigned_identity_id,omitempty"`
}

type FlexibleServerInitParameters struct {

	// The Administrator login for the MySQL Flexible Server. Required when create_mode is Default. Changing this forces a new MySQL Flexible Server to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// The backup retention days for the MySQL Flexible Server. Possible values are between 1 and 35 days. Defaults to 7.
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// The creation mode which can be used to restore or replicate existing servers. Possible values are Default, PointInTimeRestore, GeoRestore, and Replica. Changing this forces a new MySQL Flexible Server to be created.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// A customer_managed_key block as defined below.
	CustomerManagedKey []CustomerManagedKeyInitParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Should geo redundant backup enabled? Defaults to false. Changing this forces a new MySQL Flexible Server to be created.
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// A high_availability block as defined below.
	HighAvailability []HighAvailabilityInitParameters `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A maintenance_window block as defined below.
	MaintenanceWindow []MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The point in time to restore from creation_source_server_id when create_mode is PointInTimeRestore. Changing this forces a new MySQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc,omitempty"`

	// The replication role. Possible value is None.
	ReplicationRole *string `json:"replicationRole,omitempty" tf:"replication_role,omitempty"`

	// The SKU Name for the MySQL Flexible Server.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The resource ID of the source MySQL Flexible Server to be restored. Required when create_mode is PointInTimeRestore, GeoRestore, and Replica. Changing this forces a new MySQL Flexible Server to be created.
	SourceServerID *string `json:"sourceServerId,omitempty" tf:"source_server_id,omitempty"`

	// A storage block as defined below.
	Storage []StorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags which should be assigned to the MySQL Flexible Server.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version of the MySQL Flexible Server to use. Possible values are 5.7, and 8.0.21. Changing this forces a new MySQL Flexible Server to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Specifies the Availability Zone in which this MySQL Flexible Server should be located. Possible values are 1, 2 and 3.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type FlexibleServerObservation struct {

	// The Administrator login for the MySQL Flexible Server. Required when create_mode is Default. Changing this forces a new MySQL Flexible Server to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// The backup retention days for the MySQL Flexible Server. Possible values are between 1 and 35 days. Defaults to 7.
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// The creation mode which can be used to restore or replicate existing servers. Possible values are Default, PointInTimeRestore, GeoRestore, and Replica. Changing this forces a new MySQL Flexible Server to be created.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// A customer_managed_key block as defined below.
	CustomerManagedKey []CustomerManagedKeyObservation `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// The ID of the virtual network subnet to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
	DelegatedSubnetID *string `json:"delegatedSubnetId,omitempty" tf:"delegated_subnet_id,omitempty"`

	// The fully qualified domain name of the MySQL Flexible Server.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// Should geo redundant backup enabled? Defaults to false. Changing this forces a new MySQL Flexible Server to be created.
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// A high_availability block as defined below.
	HighAvailability []HighAvailabilityObservation `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`

	// The ID of the MySQL Flexible Server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A maintenance_window block as defined below.
	MaintenanceWindow []MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The point in time to restore from creation_source_server_id when create_mode is PointInTimeRestore. Changing this forces a new MySQL Flexible Server to be created.
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc,omitempty"`

	// The ID of the private DNS zone to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	// Is the public network access enabled?
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The maximum number of replicas that a primary MySQL Flexible Server can have.
	ReplicaCapacity *float64 `json:"replicaCapacity,omitempty" tf:"replica_capacity,omitempty"`

	// The replication role. Possible value is None.
	ReplicationRole *string `json:"replicationRole,omitempty" tf:"replication_role,omitempty"`

	// The name of the Resource Group where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU Name for the MySQL Flexible Server.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The resource ID of the source MySQL Flexible Server to be restored. Required when create_mode is PointInTimeRestore, GeoRestore, and Replica. Changing this forces a new MySQL Flexible Server to be created.
	SourceServerID *string `json:"sourceServerId,omitempty" tf:"source_server_id,omitempty"`

	// A storage block as defined below.
	Storage []StorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags which should be assigned to the MySQL Flexible Server.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version of the MySQL Flexible Server to use. Possible values are 5.7, and 8.0.21. Changing this forces a new MySQL Flexible Server to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Specifies the Availability Zone in which this MySQL Flexible Server should be located. Possible values are 1, 2 and 3.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type FlexibleServerParameters struct {

	// The Administrator login for the MySQL Flexible Server. Required when create_mode is Default. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// The Password associated with the administrator_login for the MySQL Flexible Server. Required when create_mode is Default.
	// +kubebuilder:validation:Optional
	AdministratorPasswordSecretRef *v1.SecretKeySelector `json:"administratorPasswordSecretRef,omitempty" tf:"-"`

	// The backup retention days for the MySQL Flexible Server. Possible values are between 1 and 35 days. Defaults to 7.
	// +kubebuilder:validation:Optional
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// The creation mode which can be used to restore or replicate existing servers. Possible values are Default, PointInTimeRestore, GeoRestore, and Replica. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// A customer_managed_key block as defined below.
	// +kubebuilder:validation:Optional
	CustomerManagedKey []CustomerManagedKeyParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// The ID of the virtual network subnet to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DelegatedSubnetID *string `json:"delegatedSubnetId,omitempty" tf:"delegated_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate delegatedSubnetId.
	// +kubebuilder:validation:Optional
	DelegatedSubnetIDRef *v1.Reference `json:"delegatedSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate delegatedSubnetId.
	// +kubebuilder:validation:Optional
	DelegatedSubnetIDSelector *v1.Selector `json:"delegatedSubnetIdSelector,omitempty" tf:"-"`

	// Should geo redundant backup enabled? Defaults to false. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// A high_availability block as defined below.
	// +kubebuilder:validation:Optional
	HighAvailability []HighAvailabilityParameters `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A maintenance_window block as defined below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The point in time to restore from creation_source_server_id when create_mode is PointInTimeRestore. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc,omitempty"`

	// The ID of the private DNS zone to create the MySQL Flexible Server. Changing this forces a new MySQL Flexible Server to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PrivateDNSZone
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	// Reference to a PrivateDNSZone in network to populate privateDnsZoneId.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIDRef *v1.Reference `json:"privateDnsZoneIdRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone in network to populate privateDnsZoneId.
	// +kubebuilder:validation:Optional
	PrivateDNSZoneIDSelector *v1.Selector `json:"privateDnsZoneIdSelector,omitempty" tf:"-"`

	// The replication role. Possible value is None.
	// +kubebuilder:validation:Optional
	ReplicationRole *string `json:"replicationRole,omitempty" tf:"replication_role,omitempty"`

	// The name of the Resource Group where the MySQL Flexible Server should exist. Changing this forces a new MySQL Flexible Server to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU Name for the MySQL Flexible Server.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// The resource ID of the source MySQL Flexible Server to be restored. Required when create_mode is PointInTimeRestore, GeoRestore, and Replica. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	SourceServerID *string `json:"sourceServerId,omitempty" tf:"source_server_id,omitempty"`

	// A storage block as defined below.
	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags which should be assigned to the MySQL Flexible Server.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version of the MySQL Flexible Server to use. Possible values are 5.7, and 8.0.21. Changing this forces a new MySQL Flexible Server to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Specifies the Availability Zone in which this MySQL Flexible Server should be located. Possible values are 1, 2 and 3.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type HighAvailabilityInitParameters struct {

	// The high availability mode for the MySQL Flexible Server. Possibles values are SameZone and ZoneRedundant.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Specifies the Availability Zone in which the standby Flexible Server should be located. Possible values are 1, 2 and 3.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty" tf:"standby_availability_zone,omitempty"`
}

type HighAvailabilityObservation struct {

	// The high availability mode for the MySQL Flexible Server. Possibles values are SameZone and ZoneRedundant.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Specifies the Availability Zone in which the standby Flexible Server should be located. Possible values are 1, 2 and 3.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty" tf:"standby_availability_zone,omitempty"`
}

type HighAvailabilityParameters struct {

	// The high availability mode for the MySQL Flexible Server. Possibles values are SameZone and ZoneRedundant.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// Specifies the Availability Zone in which the standby Flexible Server should be located. Possible values are 1, 2 and 3.
	// +kubebuilder:validation:Optional
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty" tf:"standby_availability_zone,omitempty"`
}

type IdentityInitParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this MySQL Flexible Server.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this MySQL Flexible Server. The only possible value is UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this MySQL Flexible Server.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this MySQL Flexible Server. The only possible value is UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// A list of User Assigned Managed Identity IDs to be assigned to this MySQL Flexible Server.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this MySQL Flexible Server. The only possible value is UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type MaintenanceWindowInitParameters struct {

	// The day of week for maintenance window. Defaults to 0.
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// The start hour for maintenance window. Defaults to 0.
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// The start minute for maintenance window. Defaults to 0.
	StartMinute *float64 `json:"startMinute,omitempty" tf:"start_minute,omitempty"`
}

type MaintenanceWindowObservation struct {

	// The day of week for maintenance window. Defaults to 0.
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// The start hour for maintenance window. Defaults to 0.
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// The start minute for maintenance window. Defaults to 0.
	StartMinute *float64 `json:"startMinute,omitempty" tf:"start_minute,omitempty"`
}

type MaintenanceWindowParameters struct {

	// The day of week for maintenance window. Defaults to 0.
	// +kubebuilder:validation:Optional
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// The start hour for maintenance window. Defaults to 0.
	// +kubebuilder:validation:Optional
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// The start minute for maintenance window. Defaults to 0.
	// +kubebuilder:validation:Optional
	StartMinute *float64 `json:"startMinute,omitempty" tf:"start_minute,omitempty"`
}

type StorageInitParameters struct {

	// Should Storage Auto Grow be enabled? Defaults to true.
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`

	// The storage IOPS for the MySQL Flexible Server. Possible values are between 360 and 20000.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The max storage allowed for the MySQL Flexible Server. Possible values are between 20 and 16384.
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type StorageObservation struct {

	// Should Storage Auto Grow be enabled? Defaults to true.
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`

	// The storage IOPS for the MySQL Flexible Server. Possible values are between 360 and 20000.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The max storage allowed for the MySQL Flexible Server. Possible values are between 20 and 16384.
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

type StorageParameters struct {

	// Should Storage Auto Grow be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`

	// The storage IOPS for the MySQL Flexible Server. Possible values are between 360 and 20000.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The max storage allowed for the MySQL Flexible Server. Possible values are between 20 and 16384.
	// +kubebuilder:validation:Optional
	SizeGb *float64 `json:"sizeGb,omitempty" tf:"size_gb,omitempty"`
}

// FlexibleServerSpec defines the desired state of FlexibleServer
type FlexibleServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerParameters `json:"forProvider"`
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
	InitProvider FlexibleServerInitParameters `json:"initProvider,omitempty"`
}

// FlexibleServerStatus defines the observed state of FlexibleServer.
type FlexibleServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServer is the Schema for the FlexibleServers API. Manages a MySQL Flexible Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   FlexibleServerSpec   `json:"spec"`
	Status FlexibleServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerList contains a list of FlexibleServers
type FlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServer `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServer_Kind             = "FlexibleServer"
	FlexibleServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServer_Kind}.String()
	FlexibleServer_KindAPIVersion   = FlexibleServer_Kind + "." + CRDGroupVersion.String()
	FlexibleServer_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServer_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServer{}, &FlexibleServerList{})
}
