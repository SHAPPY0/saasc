package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
	"github.com/shappy0/saasc/internal/utils"
)

var TitleRG = "resource groups"

type ResourceGroups struct {
	*widgets.Table
	Title		string
	Headers		[]string
	Data		[]models.ResourceGroup
}

func NewResourceGroups() *ResourceGroups {
	rg := ResourceGroups{
		Title:		TitleRG,
		Table:		widgets.NewTable(TitleRG),
		Headers:	[]string{"", "name", "location"},
	}
	rg.Table.Headers = rg.Headers
	rg.Table.DrawHeader()
	return &rg
}

func (rg *ResourceGroups) GetTitle() string {
	return rg.Title
}

func (rg *ResourceGroups) UpdateData(data []models.ResourceGroup) {
	rg.Data = data
	rg.SetTableTitle(len(rg.Data), "", "")
	for i := 0; i < len(rg.Data); i++ {
		rg.Table.DrawCell(i + 1, 0, utils.IntToStr(i + 1) + ".", tcell.ColorWhite)
		rg.Table.DrawCell(i + 1, 1, rg.Data[i].Name, tcell.ColorWhite)
		rg.Table.DrawCell(i + 1, 2, rg.Data[i].Location, tcell.ColorWhite)
	}
}