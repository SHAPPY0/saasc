package core

import (
	"fmt"
	"github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/azure"
)

type App struct {
	Config 		*config.Conf
	Version		string
	Layout		*Layout
	Primitives 	PrimitivesX
	Azure		*azure.Client
}

type PrimitivesX struct {
	Plans	*Plans
}

func NewApp(c *config.Conf) (*App, error) {
	a := App{
		Config: 		c,
		Version: 		c.Version,
		Layout: 		NewLayout(c),
	}
	azClient, err := azure.NewClient(c)
	if err != nil {
		fmt.Println(err)
	}
	a.Azure = azClient
	return &a, nil
}

func (a *App) Init() error {
	a.Primitives = PrimitivesX{
		Plans:		NewPlans(a),
	}
	return nil
}

func (a *App) RunX() error {
	if err := a.Layout.Run(a); err != nil {
		return err
	}
	return nil
}

func (a *App) StopX() {
	a.Layout.Stop()
}