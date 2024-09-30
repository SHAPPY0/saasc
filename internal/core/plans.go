package core

import (
	// "github.com/shappy0/saasc/internal/utils"
	"github.com/shappy0/saasc/internal/views"
)

type Plans struct {
	*views.Plans
	App				*App
	ResourceGroup	string
	SelectedValue	map[string]string
}

func NewPlans(app *App) *Plans {
	p := &Plans{
		Plans:	views.NewPlans(),
		App:	app,
	}
	p.App.Layout.Body.AddPageX(p.GetTitle(), p, true, false)
	return p
}

func (p *Plans) RenderView(rg string) {
	p.ResourceGroup = rg
	p.App.Alert.Loader(true)
	data, err := p.App.Azure.PlansClient.List(p.ResourceGroup)
	p.App.Alert.Loader(false)
	if err != nil {
		p.App.Alert.Error(err.Error())
	}
	p.UpdateData(p.ResourceGroup, data)
}