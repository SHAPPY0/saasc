package utils

import (
	"time"
	"github.com/shappy0/saasc/internal/models"
	"github.com/shappy0/saasc/internal/config"
)

const (
	DefaultFlashDelay	=	3 * time.Second
	Loader	=	"loader"
)

type Alert struct {
	Message 	models.Alert
	Duration	time.Duration
	AlertChan 	chan models.Alert
}

func NewAlert() *Alert {
	a := &Alert{
		Duration:		DefaultFlashDelay,
		AlertChan:		make(models.AlertChan, 3),
	}
	return a
}

func (a *Alert) Channel() models.AlertChan {
	return a.AlertChan
}

func (a *Alert) Info(Msg string) {
	a.SendMessage(config.Info, Msg)
}

func (a *Alert) Warning(Msg string) {
	a.SendMessage(config.Warning, Msg)
}

func (a *Alert) Error(Msg string) {
	a.SendMessage(config.Error, Msg)
}

func (a *Alert) SendMessage(Type, Msg string) {
	a.Message = models.Alert{Type: Type, Text: Msg}
	a.AlertChan <-a.Message
	go a.Hide()
}

func (a *Alert) Hide() {
	for {
		select{
		case <-time.After(a.Duration):
			a.AlertChan <-models.Alert{}
			return
		}
	}
}

func (a *Alert) Loader(loading bool) {
	if loading {
		a.Message = models.Alert{Type: Loader, Text: "Loading..."}
		a.AlertChan <-a.Message
	} else {
		a.AlertChan <-models.Alert{}
	}
}