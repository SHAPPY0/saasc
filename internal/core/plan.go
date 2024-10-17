package core

import (
	// "github.com/shappy0/saasc/internal/utils"
	"github.com/shappy0/saasc/internal/views"
)

type Plan struct {
	*views.Plan
	App				*App
	ResourceGroup	string
	PlanName		string
	SelectedRow		map[string]string
}

func NewPlan(app *App) *Plan {
	p := &Plan{
		Plan:	views.NewPlan(),
		App:	app,
	}
	p.App.Layout.Body.AddPageX(p.GetTitle(), p, true, false)
	// p.SetOnSelectFn(p.OnRowSelect)
	// p.SetFocusFunc(p.OnFocus)
	return p
}

// func (p *Plans) OnRowSelect(row, col int) {
// 	p.SelectedRow = p.GetSelectedItem()
// 	go func() {
// 		p.App.Layout.QueueUpdateDraw(func() {
// 			p.App.Primitives.WebApps.RenderView(p.ResourceGroup)
// 			p.App.Layout.OpenPage(p.App.Primitives.WebApps.GetTitle(), true)
// 		})
// 	}()
// }

// func (p *Plans) OnFocus() {
// 	p.RenderView(p.App.Config.GetResourceGroup())
// }

func (p *Plan) RenderView(rg string, SelectedRow map[string]string) {
	p.ResourceGroup = rg
	p.PlanName = SelectedRow["name"]
	if p.PlanName == "" {
		p.App.Alert.Error("[Plan::RenderView] Invalid PlanName")
	} else {
		p.App.Alert.Loader(true)
		data, err := p.App.Azure.PlansClient.Get(p.ResourceGroup, p.PlanName)
		p.App.Alert.Loader(false)
		if err != nil {
			p.App.Alert.Error(err.Error())
		} else {
			p.UpdateData(p.ResourceGroup, data)
		}
	}
}