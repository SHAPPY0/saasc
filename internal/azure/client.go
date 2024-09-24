package azure

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type AzureClient struct {
	Credential		*azidentity.DefaultAzureCredential
}

func NewAzureClient() *AzureClient {
	ac := AzureClient{}
	ac.Credential = ac.AzureCredential()
	return &ac
}

func (ac *AzureClient) AzureCredential() *azidentity.DefaultAzureCredential {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return cred
}