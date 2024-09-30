package core

import (
	"github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/utils"
	"github.com/shappy0/saasc/internal/azure"
)

type App struct {
	Config 		*config.Conf
	Version		string
	Layout		*Layout
	Primitives 	PrimitivesX
	Azure		*azure.Client
	Alert		*utils.Alert
	Logger 		*utils.Logger
}

type PrimitivesX struct {
	ResourceGroups	*ResourceGroups
	Plans			*Plans
}

func NewApp(c *config.Conf, logger *utils.Logger) (*App, error) {
	a := App{
		Config: 		c,
		Version: 		c.Version,
		Layout: 		NewLayout(c),
		Alert:			utils.NewAlert(),
		Logger:			logger,
	}
	azClient, err := azure.NewClient(c)
	if err != nil {
		a.Logger.Error(err.Error())
		return nil, err
	}
	a.Azure = azClient
	return &a, nil
}

func (a *App) Init() error {
	a.Primitives = PrimitivesX{
		ResourceGroups:	NewResourceGroups(a),
		Plans:			NewPlans(a),
	}
	alert := NewAlert(a)
	go alert.Listen()
	return nil
}

func (a *App) RunX() error {
	if err := a.Layout.Run(a); err != nil {
		a.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (a *App) StopX() {
	a.Layout.Stop()
}