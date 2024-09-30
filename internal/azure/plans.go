package azure

import (
	"fmt"
	"log"
	"context"
	// "github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/models"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

type Plans struct {
	Azure 			*Client
	Client			*armappservice.PlansClient
	ClientFactory	*armappservice.ClientFactory
}

type PlansClient interface {
	List(string)	([]models.Plan, error)
}

func (c *Client) NewPlans() *Plans {
	cf, err := armappservice.NewClientFactory(c.SubscriptionId, c.Credential, nil)
	if err != nil {
		fmt.Println(err)
	}
	p := Plans{
		Azure:			c,
		ClientFactory:	cf,
		Client:			cf.NewPlansClient(),
	}

	return &p
}

func (p *Plans) List(rg string) ([]models.Plan, error) {
	pager := p.Client.NewListByResourceGroupPager(rg, nil)
	ctx := context.Background()
	var data []models.Plan
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("Failed %v", err)
		}
		for _, v := range page.Value {
			plan := models.Plan{
				Name:	*v.Name,
				Location:	*v.Location,
				Type:		*v.Type,
				Id:			*v.ID,
				Kind:		*v.Kind,
			}
			data = append(data, plan)
		}
	}
	return data, nil
}