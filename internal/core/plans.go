package core

import (
	"github.com/shappy0/saasc/internal/views"
)

type Plans struct {
	*views.Plans
	App				*App
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