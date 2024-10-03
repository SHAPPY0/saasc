package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
	"github.com/shappy0/saasc/internal/utils"
)

var TitlePlans = "plans"

type Plans struct {
	*widgets.Table
	Title		string
	Headers		[]string
	Data		[]models.Plan
}

func NewPlans() *Plans {
	p := Plans{
		Title:		TitlePlans,
		Table:		widgets.NewTable(TitlePlans),
		Headers:	[]string{"", "name", "os", "status", "apps", "location", "tier"},
	}
	p.Table.Headers = p.Headers
	p.Table.DrawHeader()
	return &p
}

func (p *Plans) GetTitle() string {
	return p.Title
}

func (p *Plans) UpdateData(rg string, data []models.Plan) {
	p.Data = data
	p.SetTableTitle(len(p.Data), rg, "")
	for i := 0; i < len(p.Data); i++ {
		p.Table.DrawCell(i + 1, 0, utils.IntToStr(i + 1) + ".", tcell.ColorWhite)
		p.Table.DrawCell(i + 1, 1, p.Data[i].Name, tcell.ColorWhite)
		p.Table.DrawCell(i + 1, 2, p.Data[i].Kind, tcell.ColorWhite)
		p.Table.DrawCell(i + 1, 3, p.Data[i].Properties.Status, tcell.ColorWhite)
		p.Table.DrawCell(i + 1, 4, utils.IntToStr(int(*p.Data[i].Properties.NumberOfSites)) , tcell.ColorWhite)
		p.Table.DrawCell(i + 1, 5, p.Data[i].Location, tcell.ColorWhite)
		p.Table.DrawCell(i + 1, 6, p.Data[i].SKU.Tier, tcell.ColorWhite)

	}
}