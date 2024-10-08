package core

import (
	"github.com/shappy0/saasc/internal/views"
)

type WebApps struct {
	*views.WebApps
	App				*App
	ResourceGroup	string
	SelectedRow		map[string]string
}

func NewWebApps(app *App) *WebApps {
	wa := WebApps{
		WebApps:		views.NewWebApps(),
		App:			app,
	}
	wa.App.Layout.Body.AddPageX(wa.GetTitle(), wa, true, false)
	wa.SetOnSelectFn(wa.OnRowSelect)
	return &wa
}

func (wa *WebApps) OnRowSelect(row, col int) {
	wa.SelectedRow = wa.GetSelectedItem()
	go func() {
		wa.App.Layout.QueueUpdateDraw(func() {
			wa.App.Primitives.WebAppDetail.RenderView(wa.SelectedRow)
			wa.App.Layout.OpenPage(wa.App.Primitives.WebAppDetail.GetTitle(), true)
		})
	}()
}

func (wa *WebApps) RenderView(rg string) {
	wa.ResourceGroup = rg
	wa.App.Alert.Loader(true)
	data, err := wa.App.Azure.WebAppsClient.List(wa.ResourceGroup)
	wa.App.Alert.Loader(false)
	if err != nil {
		wa.App.Alert.Error(err.Error())
	}
	wa.UpdateData(wa.ResourceGroup, data)
}