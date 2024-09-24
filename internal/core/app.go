package core

import (
	"github.com/shappy0/saasc/internal/config"
)

type App struct {
	Config 		*config.Conf
	Version		string
	Layout		*Layout
}

type PrimitivesX struct {

}

func NewApp(c *config.Conf) (*App, error) {
	a := App{
		Config: c,
		Version: c.Version,
		Layout: NewLayout(c.Version),
	}
	return &a, nil
}

func (a *App) Init() error {
	return nil
}

func (app *App) RunX() error {
	if err := app.Layout.Run(app); err != nil {
		return err
	}
	return nil
}

func (app *App) StopX() {
	app.Layout.Stop()
}