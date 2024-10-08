package azure

import (
	"context"
	"github.com/shappy0/saasc/internal/models"
	// "github.com/shappy0/saasc/internal/utils"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

type WebApps struct {
	Azure 			*Client
	Client			*armappservice.WebAppsClient
	ClientFactory	*armappservice.ClientFactory
}

type WebAppsClient interface {
	List(string)	([]models.WebApp, error)
}

func (c *Client) NewWebApps() (*WebApps, error) {
	cf, err := armappservice.NewClientFactory(c.SubscriptionId, c.Credential, nil)
	if err != nil {
		c.Logger.Error("[Azure->NewWebApps] " + err.Error())
		return nil, err
	}
	wa := WebApps{
		Azure:		c,
		ClientFactory:	cf,
		Client:		cf.NewWebAppsClient(),
	}
	return &wa, nil
}

func (wa *WebApps) List(rg string) ([]models.WebApp, error) {
	pager := wa.Client.NewListByResourceGroupPager(rg, nil)
	ctx := context.Background()
	var data []models.WebApp
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			wa.Azure.Logger.Error("[WebAppsClient->List] " + err.Error())
			return data, err
		}
		for _, v := range page.Value {
			webapp := mapWebAppValues(v)
			data = append(data, webapp)
		}
	}
	return data, nil
}

func mapWebAppValues(data *armappservice.Site) models.WebApp {
	props := *data.Properties
	wa := models.WebApp{
		Name:						*data.Name,
		Location:					*data.Location,
		Id:							*data.ID,
		Kind:						*data.Kind,
		Identity:					mapIdentity(data.Identity),
		ClientAffinityEnabled:		*props.ClientAffinityEnabled,
		ClientCertEnabled:			*props.ClientCertEnabled,
		ContainerSize:				*props.ContainerSize,
		CustomDomainVerificationID:	*props.CustomDomainVerificationID,
		DailyMemoryTimeQuota:		*props.DailyMemoryTimeQuota,
		Enabled:					*props.Enabled,
		HTTPSOnly:					*props.HTTPSOnly,
		HostNamesDisabled:			*props.HostNamesDisabled,
		HyperV:						*props.HyperV,
		IsXenon:					*props.IsXenon,
		KeyVaultReferenceIdentity:	*props.KeyVaultReferenceIdentity,
		PublicNetworkAccess:		*props.PublicNetworkAccess,
		Reserved:					*props.Reserved,
		ScmSiteAlsoStopped:			*props.ScmSiteAlsoStopped,
		ServerFarmID:				*props.ServerFarmID,
		StorageAccountRequired:		*props.StorageAccountRequired,
		VirtualNetworkSubnetID:		*props.VirtualNetworkSubnetID,
		VnetContentShareEnabled:	*props.VnetContentShareEnabled,
		VnetImagePullEnabled:		*props.VnetImagePullEnabled,
		VnetRouteAllEnabled:		*props.VnetRouteAllEnabled,
		DefaultHostName:			*props.DefaultHostName,
		EnabledHostNames:			props.EnabledHostNames,
		HostNames:					props.HostNames,
		LastModifiedTimeUTC:		*props.LastModifiedTimeUTC,
		OutboundIPAddresses:		*props.OutboundIPAddresses,
		PossibleOutboundIPAddresses:*props.PossibleOutboundIPAddresses,
		RepositorySiteName:			*props.RepositorySiteName,
		ResourceGroup:				*props.ResourceGroup,
		State:						*props.State,
		SiteConfig:					mapSiteConfig(props.SiteConfig),
		UsageState:					string(*props.UsageState),
	}
	return wa
}

func mapIdentity(data *armappservice.ManagedServiceIdentity) models.ManagedIdentity {
	mi := models.ManagedIdentity{
		Type:					string(*data.Type),
		PrincipalID:			data.PrincipalID,
		TenantID:				data.TenantID,
		UserAssignedIdentities:	make(map[string]*models.UserAssignedIdentity),
	}
	for k, v := range data.UserAssignedIdentities{
		uai := models.UserAssignedIdentity{
			ClientID:		v.ClientID,
			PrincipalID:	v.PrincipalID,
		}
		mi.UserAssignedIdentities[k] = &uai
	}
	return mi
}

func mapSiteConfig(data *armappservice.SiteConfig) models.SiteConfig {
	sc := models.SiteConfig{
		AcrUseManagedIdentityCreds:	bool(*data.AcrUseManagedIdentityCreds),
		AcrUserManagedIdentityID:		data.AcrUserManagedIdentityID,
		AlwaysOn:						bool(*data.AlwaysOn),
		AppCommandLine:					data.AppCommandLine,
		AutoHealEnabled:				data.AutoHealEnabled,
		AutoSwapSlotName:				data.AutoSwapSlotName,
		DefaultDocuments:				data.DefaultDocuments,
		DetailedErrorLoggingEnabled:	data.DetailedErrorLoggingEnabled,
		DocumentRoot:					data.DocumentRoot,
		ElasticWebAppScaleLimit:		data.ElasticWebAppScaleLimit,
		FunctionAppScaleLimit:			data.FunctionAppScaleLimit,
		FunctionsRuntimeScaleMonitoringEnabled:	data.FunctionsRuntimeScaleMonitoringEnabled,
		HTTPLoggingEnabled:				data.HTTPLoggingEnabled,
		HealthCheckPath:				data.HealthCheckPath,
		Http20Enabled:					data.Http20Enabled,
		JavaContainer:					data.JavaContainer,
		JavaContainerVersion:			data.JavaContainerVersion,
		JavaVersion:					data.JavaVersion,
		KeyVaultReferenceIdentity:		data.KeyVaultReferenceIdentity,
		LinuxFxVersion:					string(*data.LinuxFxVersion),
		LocalMySQLEnabled:				data.LocalMySQLEnabled,
		LogsDirectorySizeLimit:			data.LogsDirectorySizeLimit,
		ManagedServiceIdentityID:		data.ManagedServiceIdentityID,
		MinimumElasticInstanceCount:	*data.MinimumElasticInstanceCount,
		NetFrameworkVersion:			data.NetFrameworkVersion,
		NodeVersion:					data.NodeVersion,
		NumberOfWorkers:				*data.NumberOfWorkers,
		PhpVersion:						data.PhpVersion,
		PowerShellVersion:				data.PowerShellVersion,
		PreWarmedInstanceCount:			data.PreWarmedInstanceCount,
		PublicNetworkAccess:			data.PublicNetworkAccess,
		PublishingUsername:				data.PublishingUsername,
		PythonVersion:					data.PythonVersion,
		RemoteDebuggingEnabled:			data.RemoteDebuggingEnabled,
		RemoteDebuggingVersion:			data.RemoteDebuggingVersion,
		RequestTracingEnabled:			data.RequestTracingEnabled,
		RequestTracingExpirationTime:	data.RequestTracingExpirationTime,
		ScmIPSecurityRestrictionsUseMain:	data.ScmIPSecurityRestrictionsUseMain,
		TracingOptions:					data.TracingOptions,
		Use32BitWorkerProcess:			data.Use32BitWorkerProcess,
		VnetName:						data.VnetName,
		VnetPrivatePortsCount:			data.VnetPrivatePortsCount,
		VnetRouteAllEnabled:			data.VnetRouteAllEnabled,
		WebSocketsEnabled:				data.WebSocketsEnabled,
		WebsiteTimeZone:				data.WebsiteTimeZone,
		WindowsFxVersion:				data.WindowsFxVersion,
		XManagedServiceIdentityID:		data.XManagedServiceIdentityID,
	}
	return sc
}