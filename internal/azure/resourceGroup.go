package azure

import (
	"fmt"
	"log"
	"sort"
	"context"
	"github.com/shappy0/saasc/internal/models"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type ResourceGroups struct {
	Azure 			*Client
	Client			*armresources.ResourceGroupsClient
	ClientFactory	*armresources.ClientFactory
}

type ResourceGroupsClient interface {
	List()	([]models.ResourceGroup, error)
}

func (c *Client) NewResourceGroups() *ResourceGroups {
	cf, err := armresources.NewClientFactory(c.SubscriptionId, c.Credential, nil)
	if err != nil {
		fmt.Println(err)
	}
	rg := ResourceGroups{
		Azure:			c,
		ClientFactory:	cf,
		Client:			cf.NewResourceGroupsClient(),
	}
	return &rg
}

func (rg *ResourceGroups) List() ([]models.ResourceGroup, error) {
	pager := rg.Client.NewListPager(nil)
	ctx := context.Background()
	var data []models.ResourceGroup
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("Failed %v", err)
		}
		for _, v := range page.Value {
			rg := models.ResourceGroup{
				Name:		*v.Name,
				Location:	*v.Location,
			}
			data = append(data, rg)
		}
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name < data[j].Name
	})
	return data, nil
}