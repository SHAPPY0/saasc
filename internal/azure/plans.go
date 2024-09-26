package azure

import (
	"fmt"
	// "github.com/shappy0/saasc/internal/config"
	// "github.com/shappy0/saasc/internal/models"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// type PlansClient interface {
// 	List() ([]string, error)
// }

type Plans struct {
	Azure 			*Client
	Client			*armappservice.PlansClient
	ClientFactory	*armappservice.ClientFactory
}

type PlansClient interface {
	List()	([]string, error)
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

func (p *Plans) List() ([]string, error) {
	return make([]string, 3), nil
}