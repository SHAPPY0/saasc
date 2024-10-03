package azure

import (
	// "fmt"
	"github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/utils"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type Client struct {
	Logger					*utils.Logger
	SubscriptionId			string
	ResourceGroup			string
	Credential				*azidentity.DefaultAzureCredential
	ResourceGroupsClient 	ResourceGroupsClient
	PlansClient				PlansClient
	WebAppsClient			WebAppsClient
}

func NewClient(config *config.Conf, logger *utils.Logger) (*Client, error) {
	c := Client{
		Logger:				logger,
		SubscriptionId:		config.AzureSubscriptionId,
		ResourceGroup:		config.AzureResourceGroup,
	}
	cred, err := c.AzureCredential()
	if err != nil {
		logger.Error("[Azure->NewClient] " + err.Error())
		return nil, err
	}
	c.Credential = cred
	rgClient, err := c.NewResourceGroups()
	if err != nil {
		return nil, err
	}
	c.ResourceGroupsClient = rgClient

	pc, err := c.NewPlans()
	if err != nil {
		return nil, err
	}
	c.PlansClient = pc

	wa, err := c.NewWebApps()
	if err != nil {
		return nil, err
	}
	c.WebAppsClient = wa
	return &c, nil
}

func (c *Client) AzureCredential() (*azidentity.DefaultAzureCredential, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}
	return cred, nil
}