package views

import (
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
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
		Headers:	[]string{"id", "name", "description"},
	}
	p.Table.Headers = p.Headers
	p.Table.DrawHeader()
	return &p
}

func (p *Plans) GetTitle() string {
	return p.Title
}