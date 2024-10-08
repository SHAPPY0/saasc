package models

import (
	"time"
)

type WebApp struct {
	Id						string
	Kind					string
	Location				string
	Name					string
	Identity				ManagedIdentity
	ClientAffinityEnabled	bool
	ClientCertEnabled		bool
	ClientCertExclusionPaths	string
	ContainerSize			int32
	CustomDomainVerificationID	string
	DailyMemoryTimeQuota	int32
	Enabled					bool
	HTTPSOnly				bool
	HostNamesDisabled		bool
	HyperV					bool
	IsXenon					bool
	KeyVaultReferenceIdentity	string
	PublicNetworkAccess		string
	Reserved				bool
	ScmSiteAlsoStopped		bool
	ServerFarmID			string
	SiteConfig				SiteConfig
	StorageAccountRequired	bool
	VirtualNetworkSubnetID	string
	VnetContentShareEnabled	bool
	VnetImagePullEnabled	bool
	VnetRouteAllEnabled		bool
	DefaultHostName			string
	EnabledHostNames		[]*string
	HostNames				[]*string
	LastModifiedTimeUTC		time.Time
	OutboundIPAddresses		string
	PossibleOutboundIPAddresses	string
	RepositorySiteName		string
	ResourceGroup			string
	State					string
	UsageState				string
}

type ManagedIdentity struct {
	Type 				string
	UserAssignedIdentities	map[string]*UserAssignedIdentity
	PrincipalID			*string
	TenantID			*string
}

type UserAssignedIdentity struct {
	ClientID			*string
	PrincipalID			*string
}

type SiteConfig struct {
	// APIDefinition 				*APIDefinitionInfo
	// APIManagementConfig 			*APIManagementConfig
	AcrUseManagedIdentityCreds 		bool
	AcrUserManagedIdentityID 		*string
	AlwaysOn 						bool
	AppCommandLine 					*string
	// AppSettings 					[]*NameValuePair
	AutoHealEnabled 				*bool
	// AutoHealRules 				*AutoHealRules
	AutoSwapSlotName 				*string
	// AzureStorageAccounts 			map[string]*AzureStorageInfoValue
	// ConnectionStrings 				[]*ConnStringInfo
	// Cors 							*CorsSettings
	DefaultDocuments 				[]*string
	DetailedErrorLoggingEnabled 	*bool
	DocumentRoot 					*string
	ElasticWebAppScaleLimit 		*int32
	// Experiments 					*Experiments
	// FtpsState 						*FtpsState
	FunctionAppScaleLimit 			*int32
	FunctionsRuntimeScaleMonitoringEnabled *bool
	HTTPLoggingEnabled 				*bool
	// HandlerMappings 				[]*HandlerMapping
	HealthCheckPath 				*string
	Http20Enabled 					*bool
	// IPSecurityRestrictions 			[]*IPSecurityRestriction
	// IPSecurityRestrictionsDefaultAction *DefaultAction
	JavaContainer 					*string
	JavaContainerVersion 			*string
	JavaVersion 					*string
	KeyVaultReferenceIdentity 		*string
	// Limits 							*SiteLimits
	LinuxFxVersion 					string
	// LoadBalancing 					*SiteLoadBalancing
	LocalMySQLEnabled 				*bool
	LogsDirectorySizeLimit 			*int32
	// ManagedPipelineMode 			*ManagedPipelineMode
	ManagedServiceIdentityID 		*int32
	// Metadata 						[]*NameValuePair
	// MinTLSCipherSuite 				*TLSCipherSuites
	// MinTLSVersion 					*SupportedTLSVersions
	MinimumElasticInstanceCount 	int32
	NetFrameworkVersion 			*string
	NodeVersion 					*string
	NumberOfWorkers 				int32
	PhpVersion 						*string
	PowerShellVersion 				*string
	PreWarmedInstanceCount 			*int32
	PublicNetworkAccess 			*string
	PublishingUsername 				*string
	// Push 							*PushSettings
	PythonVersion 					*string
	RemoteDebuggingEnabled 			*bool
	RemoteDebuggingVersion 			*string
	RequestTracingEnabled 			*bool
	RequestTracingExpirationTime 	*time.Time
	// ScmIPSecurityRestrictions 		[]*IPSecurityRestriction
	// ScmIPSecurityRestrictionsDefaultAction *DefaultAction
	ScmIPSecurityRestrictionsUseMain *bool
	// ScmMinTLSVersion 				*SupportedTLSVersions
	// ScmType 						*ScmType
	TracingOptions 					*string
	Use32BitWorkerProcess 			*bool
	// VirtualApplications 			[]*VirtualApplication
	VnetName 						*string
	VnetPrivatePortsCount 			*int32
	VnetRouteAllEnabled 			*bool
	WebSocketsEnabled 				*bool
	WebsiteTimeZone 				*string
	WindowsFxVersion 				*string
	XManagedServiceIdentityID 		*int32
	// MachineKey 						*SiteMachineKey
}