package models

import (
	"time"
)

type WebApp struct {
	Id						string
	Kind					string
	Location				string
	Name					string
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