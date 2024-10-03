package models

type ProvisioningState 	string
type StatusOptions		string

const (
	ProvisioningStateCanceled	ProvisioningState = "Canceled"
	ProvisioningStateDeleting	ProvisioningState = "Deleting"
	ProvisioningStateFailed		ProvisioningState = "Failed"
	ProvisioningStateInProgress	ProvisioningState = "InProgress"
	ProvisioningStateSucceeded	ProvisioningState = "Succeeded"
)

const (
	StatusOptionsCreating 	StatusOptions = "Creating"
	StatusOptionsPending 	StatusOptions = "Pending"
	StatusOptionsReady 		StatusOptions = "Ready"
)

type Plan struct {
	Location 	string
	Kind		string
	Tags		map[string]*string
	Id			string
	Name		string
	Type		string
	Properties	*PlanProperties
	SKU			*SKUDescription
}

type PlanProperties struct {
	ElasticScaleEnabled				*bool
	HyperV							*bool
	IsSpot							*bool
	IsXenon							*bool
	MaximumElasticWorkerCount		*int32
	PerSiteScaling					*bool
	Reserved						*bool
	TargetWorkerCount				*int32
	TargetWorkerSizeID				*int32
	ZoneRedundant					*bool
	GeoRegion						*string
	MaximumNumberOfWorkers			*int32
	NumberOfSites					*int32
	NumberOfWorkers					*int32
	ProvisioningState				string
	ResourceGroup					*string
	Status							string
	Subscription					*string
}

type SKUDescription struct {
	Capacity			int32
	Family				string
	Locations			[]*string
	Name				string
	Size				string
	Tier				string
}