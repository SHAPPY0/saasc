package views

import (
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
	"github.com/shappy0/saasc/internal/utils"
)

var TitlePlan = "plan"

type Plan struct {
	*widgets.Flex
	Title			string
	FirstRow		*widgets.Flex
	Essential		*widgets.Container
	Data			*models.Plan
}

func NewPlan() *Plan {
	p := Plan{
		Flex:		widgets.NewFlex(),
		Title:		TitlePlan,
		FirstRow:	widgets.NewFlex(),
		Essential:	widgets.NewContainer("", 0, 1),
	}
	p.SetTitleX(TitlePlan, "")
	p.AddItemX(p.FirstRow, 0, 1, false)
	return &p
}

func (p *Plan) GetTitle() string {
	return p.Title
}

func (p *Plan) UpdateData(rg string, data *models.Plan) {
	p.Data = data
	p.SetTitleX(TitlePlan, p.Data.Name)
	p.DrawFirstRow()
}

func (p *Plan) DrawFirstRow() {
	p.FirstRow.Border(false)
	p.FirstRow.ClearFlex()
	p.Essential.Clear()
	p.SetEssentialValues()
	p.Essential.Border(false)
	p.FirstRow.AddItemX(p.Essential, 0, 1, false)
}

func (p *Plan) SetEssentialValues() {
	p.Essential.SetKeyValue("Name:", p.Data.Name)
	p.Essential.SetKeyValue("Kind:", p.Data.Kind)
	p.Essential.SetKeyValue("Location:", p.Data.Location)
	p.Essential.SetKeyValue("Type:", p.Data.Type)
	p.Essential.SetKeyValue("SKU Name:", p.Data.SKU.Name)
	p.Essential.SetKeyValue("SKU Tier:", p.Data.SKU.Tier)
	p.Essential.SetKeyValue("SKU Capacity:", utils.IntToStr(int(p.Data.SKU.Capacity)))
	p.Essential.SetKeyValue("SKU Family:", p.Data.SKU.Family)
	locations := ""
	for i := 0; i < len(p.Data.SKU.Locations); i++ {
		locations += *p.Data.SKU.Locations[i]
	}
	p.Essential.SetKeyValue("SKU Locations:", locations)
	p.Essential.SetKeyValue("SKU Size:", p.Data.SKU.Size)

	p.Essential.SetKeyValue("ProvisioningState:", p.Data.Properties.ProvisioningState)
	p.Essential.SetKeyValue("ResourceGroup:", *p.Data.Properties.ResourceGroup)
	p.Essential.SetKeyValue("Status:", p.Data.Properties.Status)
	p.Essential.SetKeyValue("Subscription:", *p.Data.Properties.Subscription)
	p.Essential.SetKeyValue("ElasticScaleEnabled:", utils.BootToStr(*p.Data.Properties.ElasticScaleEnabled))
	p.Essential.SetKeyValue("PerSiteScaling:", utils.BootToStr(*p.Data.Properties.PerSiteScaling))
	p.Essential.SetKeyValue("Reserved:", utils.BootToStr(*p.Data.Properties.Reserved))
	p.Essential.SetKeyValue("TargetWorkerCount:", utils.IntToStr(int(*p.Data.Properties.TargetWorkerCount)))
	p.Essential.SetKeyValue("TargetWorkerSizeID:", utils.IntToStr(int(*p.Data.Properties.TargetWorkerSizeID)))
	p.Essential.SetKeyValue("ZoneRedundant:", utils.BootToStr(*p.Data.Properties.ZoneRedundant))
	p.Essential.SetKeyValue("GeoRegion:", *p.Data.Properties.GeoRegion)
	p.Essential.SetKeyValue("MaximumNumberOfWorkers:", utils.IntToStr(int(*p.Data.Properties.MaximumNumberOfWorkers)))
	p.Essential.SetKeyValue("NumberOfSites:", utils.IntToStr(int(*p.Data.Properties.NumberOfSites)))
	p.Essential.SetKeyValue("NumberOfWorkers:", utils.IntToStr(int(*p.Data.Properties.NumberOfWorkers)))	
}