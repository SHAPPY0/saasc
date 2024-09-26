package azure

import (
	"fmt"
	"github.com/shappy0/saasc/internal/config"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type Client struct {
	SubscriptionId	string
	ResourceGroup	string
	Credential		*azidentity.DefaultAzureCredential
	PlansClient		PlansClient
}

func NewClient(config *config.Conf) (*Client, error) {
	c := Client{
		SubscriptionId:		config.AzureSubscriptionId,
		ResourceGroup:		config.AzureResourceGroup,
	}
	cred, err := c.AzureCredential()
	if err != nil {
		return nil, err
	}
	c.Credential = cred
	c.PlansClient = c.NewPlans()
	return &c, nil
}

func (c *Client) AzureCredential() (*azidentity.DefaultAzureCredential, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return cred, nil
}