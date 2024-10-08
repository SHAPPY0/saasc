package core

import (
	"github.com/shappy0/saasc/internal/views"
	"github.com/shappy0/saasc/internal/models"
	// "github.com/shappy0/saasc/internal/utils"
)

type WebAppDetail struct {
	*views.WebAppDetail
	App				*App
	ResourceGroup	string
	SelectedRow		map[string]string
}

func NewWebAppDetail(app *App) *WebAppDetail {
	wad := WebAppDetail{
		WebAppDetail:	views.NewWebAppDetail(),
		App:			app,
	}
	wad.App.Layout.Body.AddPageX(wad.GetTitle(), wad, true, false)
	return &wad
}

func (wad *WebAppDetail) RenderView(data map[string]string) {
	var selectedWAData models.WebApp
	webAppData := wad.App.Primitives.WebApps.Data
	for i := 0; i < len(webAppData); i++ {
		if webAppData[i].Name == data["name"] {
			selectedWAData = webAppData[i]
		}
	}
	wad.UpdateData(selectedWAData)
}