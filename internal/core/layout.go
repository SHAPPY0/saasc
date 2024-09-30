package core

import (
	"time"
	"github.com/rivo/tview"
	"github.com/shappy0/saasc/internal/views"
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/config"
)

type Layout struct {
	*views.App
	Splash		*views.Splash
	Main		*widgets.Flex
	Header 		*views.Header
	Footer		*views.Footer
	Body		*widgets.Pages
}

func NewLayout(config *config.Conf) *Layout {
	l := Layout{
		App:	views.NewApp(),
		Splash:	views.NewSplash(config),
		Main:	widgets.NewFlex(),
		Header:	views.NewHeader(),
		Footer: views.NewFooter(),
		Body:	widgets.NewPages(),
	}
	l.Main.Border(false)
	l.Header.Render(config)
	l.Main.FlexRow().
			AddItem(l.Header, 5, 1, false).
			AddItem(l.Body, 0, 1, true).
			AddItem(l.Footer, 1, 1, false)

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
			l.OpenPage(app.Primitives.ResourceGroups.GetTitle(), true)
		})
	}()
	if err := l.SetFocus(l.Splash).Run(); err != nil {
		return err
	}
	return nil
}

func (l *Layout) ChangeFocusX(p tview.Primitive) {
	l.SetFocus(p)
}

func (l *Layout) OpenPage(name string, addHistory bool) {
	l.Body.OpenPageX(name, addHistory)
}


func (l *Layout) GetActivePage() string {
	return l.Body.GetActivePage()
}

func (l *Layout) GoBack() {
	l.Body.GoBack()
}