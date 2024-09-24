package azure

import (
	"fmt"
	"github.com/shappy0/saasc/internal/config"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// type PlansClient interface {
// 	List() ([]string, error)
// }

type PlansClient struct {
	AsClientFactory		*armappservice.ClientFactory
}

func NewPlansClient(config *config.Conf) *PlansClient {
	pc := PlansClient{}
	pc.AsClientFactory = pc.CreateClientFactory(config)
	return &pc
}

func (pc *PlansClient) CreateClientFactory(config *config.Conf) *armappservice.ClientFactory {
	asClientFactory, err := armappservice.NewClientFactory(config.AzureSubscriptionId, nil, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return asClientFactory
}