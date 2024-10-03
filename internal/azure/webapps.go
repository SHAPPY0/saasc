package azure

import (
	"context"
	"github.com/shappy0/saasc/internal/models"
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
		UsageState:					string(*props.UsageState),
	}
	return wa
}