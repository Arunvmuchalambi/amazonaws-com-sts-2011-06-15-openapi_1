package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetFederationTokenRequest represents the GetFederationTokenRequest schema from the OpenAPI specification
type GetFederationTokenRequest struct {
	Name interface{} `json:"Name"`
	Policy interface{} `json:"Policy,omitempty"`
	Policyarns interface{} `json:"PolicyArns,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Durationseconds interface{} `json:"DurationSeconds,omitempty"`
}

// FederatedUser represents the FederatedUser schema from the OpenAPI specification
type FederatedUser struct {
	Arn interface{} `json:"Arn"`
	Federateduserid interface{} `json:"FederatedUserId"`
}

// AssumeRoleWithSAMLRequest represents the AssumeRoleWithSAMLRequest schema from the OpenAPI specification
type AssumeRoleWithSAMLRequest struct {
	Policyarns interface{} `json:"PolicyArns,omitempty"`
	Principalarn interface{} `json:"PrincipalArn"`
	Rolearn interface{} `json:"RoleArn"`
	Samlassertion interface{} `json:"SAMLAssertion"`
	Durationseconds interface{} `json:"DurationSeconds,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
}

// DecodeAuthorizationMessageResponse represents the DecodeAuthorizationMessageResponse schema from the OpenAPI specification
type DecodeAuthorizationMessageResponse struct {
	Decodedmessage interface{} `json:"DecodedMessage,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// GetCallerIdentityRequest represents the GetCallerIdentityRequest schema from the OpenAPI specification
type GetCallerIdentityRequest struct {
}

// DecodeAuthorizationMessageRequest represents the DecodeAuthorizationMessageRequest schema from the OpenAPI specification
type DecodeAuthorizationMessageRequest struct {
	Encodedmessage interface{} `json:"EncodedMessage"`
}

// GetFederationTokenResponse represents the GetFederationTokenResponse schema from the OpenAPI specification
type GetFederationTokenResponse struct {
	Credentials interface{} `json:"Credentials,omitempty"`
	Federateduser interface{} `json:"FederatedUser,omitempty"`
	Packedpolicysize interface{} `json:"PackedPolicySize,omitempty"`
}

// ProvidedContext represents the ProvidedContext schema from the OpenAPI specification
type ProvidedContext struct {
	Contextassertion interface{} `json:"ContextAssertion,omitempty"`
	Providerarn interface{} `json:"ProviderArn,omitempty"`
}

// GetSessionTokenRequest represents the GetSessionTokenRequest schema from the OpenAPI specification
type GetSessionTokenRequest struct {
	Durationseconds interface{} `json:"DurationSeconds,omitempty"`
	Serialnumber interface{} `json:"SerialNumber,omitempty"`
	Tokencode interface{} `json:"TokenCode,omitempty"`
}

// Credentials represents the Credentials schema from the OpenAPI specification
type Credentials struct {
	Accesskeyid interface{} `json:"AccessKeyId"`
	Expiration interface{} `json:"Expiration"`
	Secretaccesskey interface{} `json:"SecretAccessKey"`
	Sessiontoken interface{} `json:"SessionToken"`
}

// GetAccessKeyInfoResponse represents the GetAccessKeyInfoResponse schema from the OpenAPI specification
type GetAccessKeyInfoResponse struct {
	Account interface{} `json:"Account,omitempty"`
}

// AssumeRoleResponse represents the AssumeRoleResponse schema from the OpenAPI specification
type AssumeRoleResponse struct {
	Assumedroleuser interface{} `json:"AssumedRoleUser,omitempty"`
	Credentials interface{} `json:"Credentials,omitempty"`
	Packedpolicysize interface{} `json:"PackedPolicySize,omitempty"`
	Sourceidentity interface{} `json:"SourceIdentity,omitempty"`
}

// AssumeRoleWithWebIdentityResponse represents the AssumeRoleWithWebIdentityResponse schema from the OpenAPI specification
type AssumeRoleWithWebIdentityResponse struct {
	Subjectfromwebidentitytoken interface{} `json:"SubjectFromWebIdentityToken,omitempty"`
	Assumedroleuser interface{} `json:"AssumedRoleUser,omitempty"`
	Audience interface{} `json:"Audience,omitempty"`
	Credentials interface{} `json:"Credentials,omitempty"`
	Packedpolicysize interface{} `json:"PackedPolicySize,omitempty"`
	Provider interface{} `json:"Provider,omitempty"`
	Sourceidentity interface{} `json:"SourceIdentity,omitempty"`
}

// GetSessionTokenResponse represents the GetSessionTokenResponse schema from the OpenAPI specification
type GetSessionTokenResponse struct {
	Credentials interface{} `json:"Credentials,omitempty"`
}

// PolicyDescriptorType represents the PolicyDescriptorType schema from the OpenAPI specification
type PolicyDescriptorType struct {
	Arn interface{} `json:"arn,omitempty"`
}

// GetAccessKeyInfoRequest represents the GetAccessKeyInfoRequest schema from the OpenAPI specification
type GetAccessKeyInfoRequest struct {
	Accesskeyid interface{} `json:"AccessKeyId"`
}

// AssumeRoleWithWebIdentityRequest represents the AssumeRoleWithWebIdentityRequest schema from the OpenAPI specification
type AssumeRoleWithWebIdentityRequest struct {
	Rolearn interface{} `json:"RoleArn"`
	Rolesessionname interface{} `json:"RoleSessionName"`
	Webidentitytoken interface{} `json:"WebIdentityToken"`
	Durationseconds interface{} `json:"DurationSeconds,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
	Policyarns interface{} `json:"PolicyArns,omitempty"`
	Providerid interface{} `json:"ProviderId,omitempty"`
}

// AssumeRoleRequest represents the AssumeRoleRequest schema from the OpenAPI specification
type AssumeRoleRequest struct {
	Sourceidentity interface{} `json:"SourceIdentity,omitempty"`
	Transitivetagkeys interface{} `json:"TransitiveTagKeys,omitempty"`
	Durationseconds interface{} `json:"DurationSeconds,omitempty"`
	Policyarns interface{} `json:"PolicyArns,omitempty"`
	Providedcontexts interface{} `json:"ProvidedContexts,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Externalid interface{} `json:"ExternalId,omitempty"`
	Policy interface{} `json:"Policy,omitempty"`
	Rolesessionname interface{} `json:"RoleSessionName"`
	Tokencode interface{} `json:"TokenCode,omitempty"`
	Rolearn interface{} `json:"RoleArn"`
	Serialnumber interface{} `json:"SerialNumber,omitempty"`
}

// AssumedRoleUser represents the AssumedRoleUser schema from the OpenAPI specification
type AssumedRoleUser struct {
	Arn interface{} `json:"Arn"`
	Assumedroleid interface{} `json:"AssumedRoleId"`
}

// GetCallerIdentityResponse represents the GetCallerIdentityResponse schema from the OpenAPI specification
type GetCallerIdentityResponse struct {
	Arn interface{} `json:"Arn,omitempty"`
	Userid interface{} `json:"UserId,omitempty"`
	Account interface{} `json:"Account,omitempty"`
}

// AssumeRoleWithSAMLResponse represents the AssumeRoleWithSAMLResponse schema from the OpenAPI specification
type AssumeRoleWithSAMLResponse struct {
	Packedpolicysize interface{} `json:"PackedPolicySize,omitempty"`
	Sourceidentity interface{} `json:"SourceIdentity,omitempty"`
	Assumedroleuser interface{} `json:"AssumedRoleUser,omitempty"`
	Audience interface{} `json:"Audience,omitempty"`
	Issuer interface{} `json:"Issuer,omitempty"`
	Namequalifier interface{} `json:"NameQualifier,omitempty"`
	Credentials interface{} `json:"Credentials,omitempty"`
	Subject interface{} `json:"Subject,omitempty"`
	Subjecttype interface{} `json:"SubjectType,omitempty"`
}
