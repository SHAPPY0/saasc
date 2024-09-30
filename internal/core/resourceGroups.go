package core

import (
	"github.com/shappy0/saasc/internal/views"
)

type ResourceGroups struct {
	*views.ResourceGroups
	App				*App
	SelectedRow		map[string]string
}

func NewResourceGroups(app *App) *ResourceGroups {
	rg := ResourceGroups{
		ResourceGroups:		views.NewResourceGroups(),
		App:				app,
	}
	rg.App.Layout.Body.AddPageX(rg.GetTitle(), rg, true, false)
	rg.SetFocusFunc(rg.OnFocus)
	rg.SetOnSelectFn(rg.OnRowSelect)
	return &rg
}

func (rg *ResourceGroups) OnFocus() {
	rg.RenderView()
}

func (rg *ResourceGroups) RenderView() {
	rg.App.Alert.Loader(true)
	data, err := rg.App.Azure.ResourceGroupsClient.List()
	rg.App.Alert.Loader(false)
	if err != nil {
		rg.App.Alert.Error(err.Error())
		rg.App.Logger.Error(err.Error())
	}
	rg.UpdateData(data)
}

func (rg *ResourceGroups) OnRowSelect(row, col int) {
	rg.SelectedRow = rg.GetSelectedItem()
	go rg.GoTo(rg.App.Primitives.Plans.GetTitle())
}

func (rg *ResourceGroups) GoTo(page string) {
	rg.App.Layout.QueueUpdateDraw(func() {
		resourceGroup := rg.SelectedRow["name"]
		rg.App.Primitives.Plans.RenderView(resourceGroup)
		rg.App.Layout.OpenPage(page,  true)
	})
}