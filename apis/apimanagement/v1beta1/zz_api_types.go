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

type APIObservation struct {

	// The ID of the API Management API.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Is this the current API Revision?
	IsCurrent *bool `json:"isCurrent,omitempty" tf:"is_current,omitempty"`

	// Is this API Revision online/accessible via the Gateway?
	IsOnline *bool `json:"isOnline,omitempty" tf:"is_online,omitempty"`
}

type APIParameters struct {

	// The Name of the API Management Service where this API should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// A description of the API Management API, which may include HTML formatting tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the API.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A import block as documented below.
	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// An oauth2_authorization block as documented below.
	// +kubebuilder:validation:Optional
	Oauth2Authorization []Oauth2AuthorizationParameters `json:"oauth2Authorization,omitempty" tf:"oauth2_authorization,omitempty"`

	// An openid_authentication block as documented below.
	// +kubebuilder:validation:Optional
	OpenIDAuthentication []OpenIDAuthenticationParameters `json:"openidAuthentication,omitempty" tf:"openid_authentication,omitempty"`

	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of protocols the operations in this API can be invoked. Possible values are http and https.
	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// The Name of the Resource Group where the API Management API exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Revision which used for this API.
	// +kubebuilder:validation:Required
	Revision *string `json:"revision" tf:"revision,omitempty"`

	// The description of the API Revision of the API Management API.
	// +kubebuilder:validation:Optional
	RevisionDescription *string `json:"revisionDescription,omitempty" tf:"revision_description,omitempty"`

	// Absolute URL of the backend service implementing this API.
	// +kubebuilder:validation:Optional
	ServiceURL *string `json:"serviceUrl,omitempty" tf:"service_url,omitempty"`

	// Should this API expose a SOAP frontend, rather than a HTTP frontend? Defaults to false.
	// +kubebuilder:validation:Optional
	SoapPassThrough *bool `json:"soapPassThrough,omitempty" tf:"soap_pass_through,omitempty"`

	// The API id of the source API, which could be in format azurerm_api_management_api.example.id or in format azurerm_api_management_api.example.id;rev=1
	// +kubebuilder:validation:Optional
	SourceAPIID *string `json:"sourceApiId,omitempty" tf:"source_api_id,omitempty"`

	// A subscription_key_parameter_names block as documented below.
	// +kubebuilder:validation:Optional
	SubscriptionKeyParameterNames []SubscriptionKeyParameterNamesParameters `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names,omitempty"`

	// Should this API require a subscription key?
	// +kubebuilder:validation:Optional
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty" tf:"subscription_required,omitempty"`

	// The Version number of this API, if this API is versioned.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The description of the API Version of the API Management API.
	// +kubebuilder:validation:Optional
	VersionDescription *string `json:"versionDescription,omitempty" tf:"version_description,omitempty"`

	// The ID of the Version Set which this API is associated with.
	// +kubebuilder:validation:Optional
	VersionSetID *string `json:"versionSetId,omitempty" tf:"version_set_id,omitempty"`
}

type ImportObservation struct {
}

type ImportParameters struct {

	// The format of the content from which the API Definition should be imported. Possible values are: openapi, openapi+json, openapi+json-link, openapi-link, swagger-json, swagger-link-json, wadl-link-json, wadl-xml, wsdl and wsdl-link.
	// +kubebuilder:validation:Required
	ContentFormat *string `json:"contentFormat" tf:"content_format,omitempty"`

	// The Content from which the API Definition should be imported. When a content_format of *-link-* is specified this must be a URL, otherwise this must be defined inline.
	// +kubebuilder:validation:Required
	ContentValue *string `json:"contentValue" tf:"content_value,omitempty"`

	// A wsdl_selector block as defined below, which allows you to limit the import of a WSDL to only a subset of the document. This can only be specified when content_format is wsdl or wsdl-link.
	// +kubebuilder:validation:Optional
	WsdlSelector []WsdlSelectorParameters `json:"wsdlSelector,omitempty" tf:"wsdl_selector,omitempty"`
}

type Oauth2AuthorizationObservation struct {
}

type Oauth2AuthorizationParameters struct {

	// OAuth authorization server identifier. The name of an OAuth2 Authorization Server.
	// +kubebuilder:validation:Required
	AuthorizationServerName *string `json:"authorizationServerName" tf:"authorization_server_name,omitempty"`

	// Operations scope.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type OpenIDAuthenticationObservation struct {
}

type OpenIDAuthenticationParameters struct {

	// How to send token to the server. A list of zero or more methods. Valid values are authorizationHeader and query.
	// +kubebuilder:validation:Optional
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// OpenID Connect provider identifier. The name of an OpenID Connect Provider.
	// +kubebuilder:validation:Required
	OpenIDProviderName *string `json:"openidProviderName" tf:"openid_provider_name,omitempty"`
}

type SubscriptionKeyParameterNamesObservation struct {
}

type SubscriptionKeyParameterNamesParameters struct {

	// The name of the HTTP Header which should be used for the Subscription Key.
	// +kubebuilder:validation:Required
	Header *string `json:"header" tf:"header,omitempty"`

	// The name of the QueryString parameter which should be used for the Subscription Key.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`
}

type WsdlSelectorObservation struct {
}

type WsdlSelectorParameters struct {

	// The name of endpoint (port) to import from WSDL.
	// +kubebuilder:validation:Required
	EndpointName *string `json:"endpointName" tf:"endpoint_name,omitempty"`

	// The name of service to import from WSDL.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// APISpec defines the desired state of API
type APISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIParameters `json:"forProvider"`
}

// APIStatus defines the observed state of API.
type APIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// API is the Schema for the APIs API. Manages an API within an API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APISpec   `json:"spec"`
	Status            APIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIList contains a list of APIs
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

// Repository type metadata.
var (
	API_Kind             = "API"
	API_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: API_Kind}.String()
	API_KindAPIVersion   = API_Kind + "." + CRDGroupVersion.String()
	API_GroupVersionKind = CRDGroupVersion.WithKind(API_Kind)
)

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
