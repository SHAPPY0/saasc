package core

import (
	// "github.com/shappy0/saasc/internal/utils"
	"github.com/shappy0/saasc/internal/views"
)

type Plans struct {
	*views.Plans
	App				*App
	ResourceGroup	string
	SelectedRow		map[string]string
}

func NewPlans(app *App) *Plans {
	p := &Plans{
		Plans:	views.NewPlans(),
		App:	app,
	}
	p.App.Layout.Body.AddPageX(p.GetTitle(), p, true, false)
	p.SetOnSelectFn(p.OnRowSelect)
	p.SetFocusFunc(p.OnFocus)
	return p
}

func (p *Plans) OnRowSelect(row, col int) {
	p.SelectedRow = p.GetSelectedItem()
	go func() {
		p.App.Layout.QueueUpdateDraw(func() {
			p.App.Primitives.WebApps.RenderView(p.ResourceGroup)
			p.App.Layout.OpenPage(p.App.Primitives.WebApps.GetTitle(), true)
		})
	}()
}

func (p *Plans) OnFocus() {
	p.RenderView(p.App.Config.GetResourceGroup())
}

func (p *Plans) RenderView(rg string) {
	if rg == "" {
		p.App.Alert.Error("Please select resource group")
	} else {
		p.ResourceGroup = rg
		p.App.Alert.Loader(true)
		data, err := p.App.Azure.PlansClient.List(p.ResourceGroup)
		p.App.Alert.Loader(false)
		if err != nil {
			p.App.Alert.Error(err.Error())
		}
		p.UpdateData(p.ResourceGroup, data)
	}
}