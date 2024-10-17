package azure

import (
	"fmt"
	// "log"
	"context"
	// "github.com/shappy0/saasc/internal/utils"
	"github.com/shappy0/saasc/internal/models"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

type Plans struct {
	Azure 			*Client
	Client			*armappservice.PlansClient
	ClientFactory	*armappservice.ClientFactory
}

type PlansClient interface {
	List(string)		([]models.Plan, error)
	Get(string, string) (*models.Plan, error)
}

func (c *Client) NewPlans() (*Plans, error) {
	cf, err := armappservice.NewClientFactory(c.SubscriptionId, c.Credential, nil)
	if err != nil {
		c.Logger.Error("[Azure->NewPlans] " + err.Error())
		return nil, err
	}
	p := Plans{
		Azure:			c,
		ClientFactory:	cf,
		Client:			cf.NewPlansClient(),
	}
	return &p, nil
}

func (p *Plans) List(rg string) ([]models.Plan, error) {
	pager := p.Client.NewListByResourceGroupPager(rg, nil)
	ctx := context.Background()
	var data []models.Plan
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			p.Azure.Logger.Error("[PlanClient->List] " + err.Error())
			return data, err
		}
		for _, v := range page.Value {
			plan := models.Plan{
				Name:		*v.Name,
				Location:	*v.Location,
				Type:		*v.Type,
				Id:			*v.ID,
				Kind:		*v.Kind,
				Properties:	mapPlanProperties(v.Properties),
				SKU:		&models.SKUDescription{
					Capacity:		*v.SKU.Capacity,
					Family:			*v.SKU.Family,
					Locations:		v.SKU.Locations,
					Name:			*v.SKU.Size,
					Tier:			*v.SKU.Tier,
				},
			}
			data = append(data, plan)
		}
	}
	return data, nil
}

func mapPlanProperties(pp *armappservice.PlanProperties) *models.PlanProperties {
	planProperties := models.PlanProperties{
		ElasticScaleEnabled:		pp.ElasticScaleEnabled,
		HyperV:						pp.HyperV,
		IsXenon:					pp.IsXenon,
		MaximumElasticWorkerCount:	pp.MaximumElasticWorkerCount,
		PerSiteScaling:				pp.PerSiteScaling,
		Reserved:					pp.Reserved,
		TargetWorkerCount:			pp.TargetWorkerCount,
		TargetWorkerSizeID:			pp.TargetWorkerSizeID,
		ZoneRedundant:				pp.ZoneRedundant,
		GeoRegion:					pp.GeoRegion,
		MaximumNumberOfWorkers:		pp.MaximumNumberOfWorkers,
		NumberOfSites:				pp.NumberOfSites,
		NumberOfWorkers:			pp.NumberOfWorkers,
		ProvisioningState:			string(*pp.ProvisioningState),
		ResourceGroup:				pp.ResourceGroup,
		Status:						string(*pp.Status),
		Subscription:				pp.Subscription,
	}
	return &planProperties
}

func (p *Plans) Get(rg, name string) (*models.Plan, error) {
	var data *models.Plan
	if rg == "" || name == "" {
		msg := "[Plan:Get] rg and plan name required"
		p.Azure.Logger.Error(msg)
		return nil, fmt.Errorf(msg)
	}
	ctx := context.Background()
	resp, err := p.Client.Get(ctx, rg, name, nil)
	if err != nil {
		return nil, err
	}
	data = &models.Plan{
		Name:		*resp.Name,
		Location:	*resp.Location,
		Type:		*resp.Type,
		Id:			*resp.ID,
		Kind:		*resp.Kind,
		Properties:	mapPlanProperties(resp.Properties),
		SKU:		&models.SKUDescription{
			Capacity:		*resp.SKU.Capacity,
			Family:			*resp.SKU.Family,
			Locations:		resp.SKU.Locations,
			Name:			*resp.SKU.Size,
			Tier:			*resp.SKU.Tier,
		},
	}
	return data, nil
}
