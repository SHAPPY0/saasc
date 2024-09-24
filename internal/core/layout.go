package core

import (
	"time"
	// "github.com/rivo/tview"
	"github.com/shappy0/saasc/internal/views"
	"github.com/shappy0/saasc/internal/widgets"
)

type Layout struct {
	*views.App
	Splash		*views.Splash
	Main		*widgets.Flex
	Header 		*views.Header
	// Footer		*views.Footer
}

func NewLayout(version string) *Layout {
	l := Layout{
		App:	views.NewApp(),
		Splash:	views.NewSplash(version),
		Main:	widgets.NewFlex(),
		Header:	views.NewHeader(),
		// Footer: views.NewFooter(),
	}
	l.Main.AddItem(l.Header, 5, 1, false)
			// AddItem(l.Body, 0, 1, true).
			// AddItem(l.Footer, 1, 1, false)

	l.SetRoot(l.Splash, true)
	return &l
}

func (l *Layout) Run(app *App) error {
	go func() {
		<- time.After(1 * time.Second)
		l.QueueUpdateDraw(func() {
			l.SetRoot(l.Main, true)
			l.SetFocus(l.Main)
			// if app.Config.IsRegionInConfig() {
			// 	app.Primitives.Jobs.UpdateTable()
			// 	l.OpenPage("jobs", true)
			// } else {
			// 	l.OpenPage("main", true)
			// }
			l.OpenPage("main", true)
		})
	}()
	if err := l.SetFocus(l.Splash).Run(); err != nil {
		return err
	}
	return nil
}