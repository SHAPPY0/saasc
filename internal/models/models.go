package models

type Alert struct {
	Type 		string
	Text 		string
}

type AlertChan chan Alert

type Metadata struct {
	ResourceGroup			string
	AzureSubscriptionId		string
	AzureClientId			string
	AzureTenantId			string
}

type ResourceGroup struct {
	Name		string
	Location	string
}
