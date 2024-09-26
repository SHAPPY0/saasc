package models

type Metadata struct {
	AzureSubscriptionId		string
	AzureClientId			string
	AzureTenantId			string
}

type Plan struct {
	Location 	string
	Kind		string
	Tags		map[string]*string
	ID			string
	Name		string
	Type		string
}