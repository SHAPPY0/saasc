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
	p.SetBlurFunc(p.OnBlur)
	return p
}

func (p *Plans) OnRowSelect(row, col int) {
	p.SelectedRow = p.GetSelectedItem()
	go func() {
		p.App.Layout.QueueUpdateDraw(func() {
			p.App.Primitives.Plan.RenderView(p.ResourceGroup, p.SelectedRow)
			p.App.Layout.OpenPage(p.App.Primitives.Plan.GetTitle(), true)
		})
	}()
}

func (p *Plans) OnFocus() {
	p.App.Layout.Header.Menu.AddGlobalMenus(p.GlobalMenus, true)
	p.RenderView(p.App.Config.GetResourceGroup())
}

func (p *Plans) OnBlur() {
	p.App.Layout.Header.Menu.RemoveGlobalMenus(p.GlobalMenus)
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