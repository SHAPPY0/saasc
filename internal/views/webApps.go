package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
)

var TitleWebApp = "web app"

type WebApps struct {
	*widgets.Table
	Title 		string
	Headers		[]string
	Data		[]models.WebApp
}

func NewWebApps() *WebApps {
	wa := WebApps{
		Title:		TitleWebApp,
		Table:		widgets.NewTable(TitleWebApp),
		Headers:	[]string{"name", "status", "kind", "usage state", "host", "public", "location"},
	}
	wa.Table.Headers = wa.Headers
	wa.Table.DrawHeader()
	return &wa
}

func (wa *WebApps) GetTitle() string {
	return wa.Title
}

func (wa *WebApps) UpdateData(rg string, data []models.WebApp) {
	wa.Data = data
	wa.SetTableTitle(len(wa.Data), rg, "")
	for i := 0; i < len(wa.Data); i++ {
		wa.Table.DrawCell(i + 1, 0, wa.Data[i].Name, tcell.ColorWhite)
		wa.Table.DrawCell(i + 1, 1, wa.Data[i].State, tcell.ColorWhite)
		wa.Table.DrawCell(i + 1, 2, wa.Data[i].Kind, tcell.ColorWhite)
		wa.Table.DrawCell(i + 1, 3, wa.Data[i].UsageState, tcell.ColorWhite)
		wa.Table.DrawCell(i + 1, 4, wa.Data[i].DefaultHostName, tcell.ColorWhite)
		wa.Table.DrawCell(i + 1, 5, wa.Data[i].PublicNetworkAccess, tcell.ColorWhite)
		wa.Table.DrawCell(i + 1, 6, wa.Data[i].Location, tcell.ColorWhite)

	}
}