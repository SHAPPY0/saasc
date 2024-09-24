package core

import (
	// "time"
	// "github.com/rivo/tview"
	"github.com/shappy0/saasc/internal/views"
)

type Layout struct {
	*views.App
	Splash		*views.Splash
}

func NewLayout(version string) *Layout {
	l := Layout{
		App:	views.NewApp(),
		Splash:	views.NewSplash(version),
	}
	l.SetRoot(l.Splash, true)
	return &l
}

func (l *Layout) Run(app *App) error {
	if err := l.SetFocus(l.Splash).Run(); err != nil {
		return err
	}
	return nil
}