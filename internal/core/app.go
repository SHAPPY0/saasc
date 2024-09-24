package core

import (
	"github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/azure"
)

type App struct {
	Config 		*config.Conf
	Version		string
	Layout		*Layout
	Primitives 	PrimitivesX
	AzureClient	*azure.AzureClient
}

type PrimitivesX struct {
	Plans	*Plans
}

func NewApp(c *config.Conf) (*App, error) {
	a := App{
		Config: 		c,
		Version: 		c.Version,
		Layout: 		NewLayout(c),
		AzureClient:	azure.NewAzureClient(),
	}
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