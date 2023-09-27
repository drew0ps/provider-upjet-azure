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

type ServerInitParameters struct {

	// The Administrator login for the MariaDB Server. Changing this forces a new resource to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is true.
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`

	// Backup retention days for the server, supported values are between 7 and 35 days.
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// The creation mode. Can be used to restore or replicate existing servers. Possible values are Default, Replica, GeoRestore, and PointInTimeRestore. Defaults to Default.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether or not public network access is allowed for this server. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// When create_mode is PointInTimeRestore, specifies the point in time to restore from creation_source_server_id. It should be provided in RFC3339 format, e.g. 2013-11-08T22:00:40Z.
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// Specifies if SSL should be enforced on connections. Possible values are true and false.
	SSLEnforcementEnabled *bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled,omitempty"`

	// The minimum TLS version to support on the sever. Possible values are TLSEnforcementDisabled, TLS1_0, TLS1_1, and TLS1_2. Defaults to TLS1_2.
	SSLMinimalTLSVersionEnforced *string `json:"sslMinimalTlsVersionEnforced,omitempty" tf:"ssl_minimal_tls_version_enforced,omitempty"`

	// Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the tier + family + cores pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the product documentation. Possible values are B_Gen5_1, B_Gen5_2, GP_Gen5_2, GP_Gen5_4, GP_Gen5_8, GP_Gen5_16, GP_Gen5_32, MO_Gen5_2, MO_Gen5_4, MO_Gen5_8 and MO_Gen5_16.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Max storage allowed for a server. Possible values are between 5120 MB (5GB) and 1024000MB (1TB) for the Basic SKU and between 5120 MB (5GB) and 4096000 MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the product documentation.
	StorageMb *float64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the version of MariaDB to use. Possible values are 10.2 and 10.3. Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ServerObservation struct {

	// The Administrator login for the MariaDB Server. Changing this forces a new resource to be created.
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is true.
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`

	// Backup retention days for the server, supported values are between 7 and 35 days.
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// The creation mode. Can be used to restore or replicate existing servers. Possible values are Default, Replica, GeoRestore, and PointInTimeRestore. Defaults to Default.
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// For creation modes other than Default, the source server ID to use.
	CreationSourceServerID *string `json:"creationSourceServerId,omitempty" tf:"creation_source_server_id,omitempty"`

	// The FQDN of the MariaDB Server.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// The ID of the MariaDB Server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether or not public network access is allowed for this server. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// When create_mode is PointInTimeRestore, specifies the point in time to restore from creation_source_server_id. It should be provided in RFC3339 format, e.g. 2013-11-08T22:00:40Z.
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// Specifies if SSL should be enforced on connections. Possible values are true and false.
	SSLEnforcementEnabled *bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled,omitempty"`

	// The minimum TLS version to support on the sever. Possible values are TLSEnforcementDisabled, TLS1_0, TLS1_1, and TLS1_2. Defaults to TLS1_2.
	SSLMinimalTLSVersionEnforced *string `json:"sslMinimalTlsVersionEnforced,omitempty" tf:"ssl_minimal_tls_version_enforced,omitempty"`

	// Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the tier + family + cores pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the product documentation. Possible values are B_Gen5_1, B_Gen5_2, GP_Gen5_2, GP_Gen5_4, GP_Gen5_8, GP_Gen5_16, GP_Gen5_32, MO_Gen5_2, MO_Gen5_4, MO_Gen5_8 and MO_Gen5_16.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Max storage allowed for a server. Possible values are between 5120 MB (5GB) and 1024000MB (1TB) for the Basic SKU and between 5120 MB (5GB) and 4096000 MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the product documentation.
	StorageMb *float64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the version of MariaDB to use. Possible values are 10.2 and 10.3. Changing this forces a new resource to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ServerParameters struct {

	// The Administrator login for the MariaDB Server. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// The Password associated with the administrator_login for the MariaDB Server.
	// +kubebuilder:validation:Optional
	AdministratorLoginPasswordSecretRef *v1.SecretKeySelector `json:"administratorLoginPasswordSecretRef,omitempty" tf:"-"`

	// Enable/Disable auto-growing of the storage. Storage auto-grow prevents your server from running out of storage and becoming read-only. If storage auto grow is enabled, the storage automatically grows without impacting the workload. The default value if not explicitly specified is true.
	// +kubebuilder:validation:Optional
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`

	// Backup retention days for the server, supported values are between 7 and 35 days.
	// +kubebuilder:validation:Optional
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// The creation mode. Can be used to restore or replicate existing servers. Possible values are Default, Replica, GeoRestore, and PointInTimeRestore. Defaults to Default.
	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// For creation modes other than Default, the source server ID to use.
	// +crossplane:generate:reference:type=Server
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CreationSourceServerID *string `json:"creationSourceServerId,omitempty" tf:"creation_source_server_id,omitempty"`

	// Reference to a Server to populate creationSourceServerId.
	// +kubebuilder:validation:Optional
	CreationSourceServerIDRef *v1.Reference `json:"creationSourceServerIdRef,omitempty" tf:"-"`

	// Selector for a Server to populate creationSourceServerId.
	// +kubebuilder:validation:Optional
	CreationSourceServerIDSelector *v1.Selector `json:"creationSourceServerIdSelector,omitempty" tf:"-"`

	// Turn Geo-redundant server backups on/off. This allows you to choose between locally redundant or geo-redundant backup storage in the General Purpose and Memory Optimized tiers. When the backups are stored in geo-redundant backup storage, they are not only stored within the region in which your server is hosted, but are also replicated to a paired data center. This provides better protection and ability to restore your server in a different region in the event of a disaster. This is not supported for the Basic tier.
	// +kubebuilder:validation:Optional
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether or not public network access is allowed for this server. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// When create_mode is PointInTimeRestore, specifies the point in time to restore from creation_source_server_id. It should be provided in RFC3339 format, e.g. 2013-11-08T22:00:40Z.
	// +kubebuilder:validation:Optional
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`

	// Specifies if SSL should be enforced on connections. Possible values are true and false.
	// +kubebuilder:validation:Optional
	SSLEnforcementEnabled *bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled,omitempty"`

	// The minimum TLS version to support on the sever. Possible values are TLSEnforcementDisabled, TLS1_0, TLS1_1, and TLS1_2. Defaults to TLS1_2.
	// +kubebuilder:validation:Optional
	SSLMinimalTLSVersionEnforced *string `json:"sslMinimalTlsVersionEnforced,omitempty" tf:"ssl_minimal_tls_version_enforced,omitempty"`

	// Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the tier + family + cores pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the product documentation. Possible values are B_Gen5_1, B_Gen5_2, GP_Gen5_2, GP_Gen5_4, GP_Gen5_8, GP_Gen5_16, GP_Gen5_32, MO_Gen5_2, MO_Gen5_4, MO_Gen5_8 and MO_Gen5_16.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// Max storage allowed for a server. Possible values are between 5120 MB (5GB) and 1024000MB (1TB) for the Basic SKU and between 5120 MB (5GB) and 4096000 MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the product documentation.
	// +kubebuilder:validation:Optional
	StorageMb *float64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the version of MariaDB to use. Possible values are 10.2 and 10.3. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
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
	InitProvider ServerInitParameters `json:"initProvider,omitempty"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Manages a MariaDB Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sslEnforcementEnabled) || (has(self.initProvider) && has(self.initProvider.sslEnforcementEnabled))",message="spec.forProvider.sslEnforcementEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   ServerSpec   `json:"spec"`
	Status ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
