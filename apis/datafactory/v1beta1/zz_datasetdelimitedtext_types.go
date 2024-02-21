// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type AzureBlobFsLocationInitParameters struct {

	// The storage data lake gen2 file system on the Azure Blob Storage Account hosting the file.
	FileSystem *string `json:"fileSystem,omitempty" tf:"file_system,omitempty"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type AzureBlobFsLocationObservation struct {

	// The storage data lake gen2 file system on the Azure Blob Storage Account hosting the file.
	FileSystem *string `json:"fileSystem,omitempty" tf:"file_system,omitempty"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type AzureBlobFsLocationParameters struct {

	// The storage data lake gen2 file system on the Azure Blob Storage Account hosting the file.
	// +kubebuilder:validation:Optional
	FileSystem *string `json:"fileSystem" tf:"file_system,omitempty"`

	// The filename of the file.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetDelimitedTextAzureBlobStorageLocationInitParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file. This can be an empty string.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetDelimitedTextAzureBlobStorageLocationObservation struct {

	// The container on the Azure Blob Storage Account hosting the file.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file. This can be an empty string.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetDelimitedTextAzureBlobStorageLocationParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	// +kubebuilder:validation:Optional
	Container *string `json:"container" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file. This can be an empty string.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetDelimitedTextHTTPServerLocationInitParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`
}

type DataSetDelimitedTextHTTPServerLocationObservation struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`
}

type DataSetDelimitedTextHTTPServerLocationParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	// +kubebuilder:validation:Optional
	RelativeURL *string `json:"relativeUrl" tf:"relative_url,omitempty"`
}

type DataSetDelimitedTextInitParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// An azure_blob_fs_location block as defined below.
	AzureBlobFsLocation []AzureBlobFsLocationInitParameters `json:"azureBlobFsLocation,omitempty" tf:"azure_blob_fs_location,omitempty"`

	// An azure_blob_storage_location block as defined below.
	AzureBlobStorageLocation []DataSetDelimitedTextAzureBlobStorageLocationInitParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// The column delimiter. Defaults to ,.
	ColumnDelimiter *string `json:"columnDelimiter,omitempty" tf:"column_delimiter,omitempty"`

	// The compression codec used to read/write text files. Valid values are None, bzip2, gzip, deflate, ZipDeflate, TarGzip, Tar, snappy and lz4. Please note these values are case sensitive.
	CompressionCodec *string `json:"compressionCodec,omitempty" tf:"compression_codec,omitempty"`

	// The compression ratio for the Data Factory Dataset. Valid values are Fastest or Optimal. Please note these values are case sensitive.
	CompressionLevel *string `json:"compressionLevel,omitempty" tf:"compression_level,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encoding format for the file.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The escape character. Defaults to \.
	EscapeCharacter *string `json:"escapeCharacter,omitempty" tf:"escape_character,omitempty"`

	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data. Defaults to false.
	FirstRowAsHeader *bool `json:"firstRowAsHeader,omitempty" tf:"first_row_as_header,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	HTTPServerLocation []DataSetDelimitedTextHTTPServerLocationInitParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceWeb
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// The null value string. Defaults to an empty string. Defaults to "".
	NullValue *string `json:"nullValue,omitempty" tf:"null_value,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The quote character. Defaults to ".
	QuoteCharacter *string `json:"quoteCharacter,omitempty" tf:"quote_character,omitempty"`

	// The row delimiter. Defaults to any of the following values on read: \r\n, \r, \n, and \n or \r\n on write by mapping data flow and Copy activity respectively.
	RowDelimiter *string `json:"rowDelimiter,omitempty" tf:"row_delimiter,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetDelimitedTextSchemaColumnInitParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetDelimitedTextObservation struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// An azure_blob_fs_location block as defined below.
	AzureBlobFsLocation []AzureBlobFsLocationObservation `json:"azureBlobFsLocation,omitempty" tf:"azure_blob_fs_location,omitempty"`

	// An azure_blob_storage_location block as defined below.
	AzureBlobStorageLocation []DataSetDelimitedTextAzureBlobStorageLocationObservation `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// The column delimiter. Defaults to ,.
	ColumnDelimiter *string `json:"columnDelimiter,omitempty" tf:"column_delimiter,omitempty"`

	// The compression codec used to read/write text files. Valid values are None, bzip2, gzip, deflate, ZipDeflate, TarGzip, Tar, snappy and lz4. Please note these values are case sensitive.
	CompressionCodec *string `json:"compressionCodec,omitempty" tf:"compression_codec,omitempty"`

	// The compression ratio for the Data Factory Dataset. Valid values are Fastest or Optimal. Please note these values are case sensitive.
	CompressionLevel *string `json:"compressionLevel,omitempty" tf:"compression_level,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encoding format for the file.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The escape character. Defaults to \.
	EscapeCharacter *string `json:"escapeCharacter,omitempty" tf:"escape_character,omitempty"`

	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data. Defaults to false.
	FirstRowAsHeader *bool `json:"firstRowAsHeader,omitempty" tf:"first_row_as_header,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	HTTPServerLocation []DataSetDelimitedTextHTTPServerLocationObservation `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The ID of the Data Factory Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// The null value string. Defaults to an empty string. Defaults to "".
	NullValue *string `json:"nullValue,omitempty" tf:"null_value,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The quote character. Defaults to ".
	QuoteCharacter *string `json:"quoteCharacter,omitempty" tf:"quote_character,omitempty"`

	// The row delimiter. Defaults to any of the following values on read: \r\n, \r, \n, and \n or \r\n on write by mapping data flow and Copy activity respectively.
	RowDelimiter *string `json:"rowDelimiter,omitempty" tf:"row_delimiter,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetDelimitedTextSchemaColumnObservation `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetDelimitedTextParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// An azure_blob_fs_location block as defined below.
	// +kubebuilder:validation:Optional
	AzureBlobFsLocation []AzureBlobFsLocationParameters `json:"azureBlobFsLocation,omitempty" tf:"azure_blob_fs_location,omitempty"`

	// An azure_blob_storage_location block as defined below.
	// +kubebuilder:validation:Optional
	AzureBlobStorageLocation []DataSetDelimitedTextAzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// The column delimiter. Defaults to ,.
	// +kubebuilder:validation:Optional
	ColumnDelimiter *string `json:"columnDelimiter,omitempty" tf:"column_delimiter,omitempty"`

	// The compression codec used to read/write text files. Valid values are None, bzip2, gzip, deflate, ZipDeflate, TarGzip, Tar, snappy and lz4. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	CompressionCodec *string `json:"compressionCodec,omitempty" tf:"compression_codec,omitempty"`

	// The compression ratio for the Data Factory Dataset. Valid values are Fastest or Optimal. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	CompressionLevel *string `json:"compressionLevel,omitempty" tf:"compression_level,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encoding format for the file.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The escape character. Defaults to \.
	// +kubebuilder:validation:Optional
	EscapeCharacter *string `json:"escapeCharacter,omitempty" tf:"escape_character,omitempty"`

	// When used as input, treat the first row of data as headers. When used as output, write the headers into the output as the first row of data. Defaults to false.
	// +kubebuilder:validation:Optional
	FirstRowAsHeader *bool `json:"firstRowAsHeader,omitempty" tf:"first_row_as_header,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	// +kubebuilder:validation:Optional
	HTTPServerLocation []DataSetDelimitedTextHTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceWeb
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// The null value string. Defaults to an empty string. Defaults to "".
	// +kubebuilder:validation:Optional
	NullValue *string `json:"nullValue,omitempty" tf:"null_value,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The quote character. Defaults to ".
	// +kubebuilder:validation:Optional
	QuoteCharacter *string `json:"quoteCharacter,omitempty" tf:"quote_character,omitempty"`

	// The row delimiter. Defaults to any of the following values on read: \r\n, \r, \n, and \n or \r\n on write by mapping data flow and Copy activity respectively.
	// +kubebuilder:validation:Optional
	RowDelimiter *string `json:"rowDelimiter,omitempty" tf:"row_delimiter,omitempty"`

	// A schema_column block as defined below.
	// +kubebuilder:validation:Optional
	SchemaColumn []DataSetDelimitedTextSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetDelimitedTextSchemaColumnInitParameters struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetDelimitedTextSchemaColumnObservation struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetDelimitedTextSchemaColumnParameters struct {

	// The description of the column.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetDelimitedTextSpec defines the desired state of DataSetDelimitedText
type DataSetDelimitedTextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetDelimitedTextParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DataSetDelimitedTextInitParameters `json:"initProvider,omitempty"`
}

// DataSetDelimitedTextStatus defines the observed state of DataSetDelimitedText.
type DataSetDelimitedTextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetDelimitedTextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DataSetDelimitedText is the Schema for the DataSetDelimitedTexts API. Manages an Azure Delimited Text Dataset inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetDelimitedText struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetDelimitedTextSpec   `json:"spec"`
	Status            DataSetDelimitedTextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDelimitedTextList contains a list of DataSetDelimitedTexts
type DataSetDelimitedTextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetDelimitedText `json:"items"`
}

// Repository type metadata.
var (
	DataSetDelimitedText_Kind             = "DataSetDelimitedText"
	DataSetDelimitedText_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetDelimitedText_Kind}.String()
	DataSetDelimitedText_KindAPIVersion   = DataSetDelimitedText_Kind + "." + CRDGroupVersion.String()
	DataSetDelimitedText_GroupVersionKind = CRDGroupVersion.WithKind(DataSetDelimitedText_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetDelimitedText{}, &DataSetDelimitedTextList{})
}
