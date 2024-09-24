package views

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
	"github.com/shappy0/saasc/internal/config"
)

type Header struct {
	*widgets.Flex
	Logo			*tview.TextView
	Metadata 		*widgets.MapView
	Menu 			*widgets.Menu
	
}

// var Logo = []string{
// 	` _  _  ____  __  __  ____`,
// 	`( \( )(_  _)(  )(  )(_  _)`,
// 	` )  (   )(   )(__)(  _)(_`,
// 	`(_)\_) (__) (______)(____)`,
// }

func NewHeader() *Header {
	h := &Header{
		Flex:		widgets.NewFlex(),
		Logo:		tview.NewTextView(),
		Metadata:	widgets.NewMapView(),
		Menu:		widgets.NewMenu(),
		
	}
	h.Border(false)
	h.AddItem(h.Logo, 0, 1, false)
	h.AddItem(h.Menu.Grid1, 0, 1, false)
	h.AddItem(h.Menu.Grid2, 0, 1, false)
	// h.AddItem(h.Menu.Grid3, 0, 1, false)
	h.AddItem(h.Metadata, 0, 1, false)
	return h
}

func (h *Header) Render(config *config.Conf) {
	//Logo
	h.RenderLogo()
	//Metadata
	metadata := models.Metadata{
		AzureSubscriptionId:	config.AzureSubscriptionId,
		AzureClientId:			config.AzureClientId,
		AzureTenantId:			config.AzureTenantId,
	}
	h.SetMetadata(metadata)
	//menu
	h.RenderMenu(make([]widgets.Item, 0))
}

func (h *Header) RenderLogo() error {
	h.Logo.SetDynamicColors(true)
	for I, S := range Logo {
		fmt.Fprintf(h.Logo, "[%s::b]%s", "", S)
		if I + 1 < len(Logo) {
			fmt.Fprintf(h.Logo, "\n")
		}
	}
	return nil
}

func (h *Header) RenderMenu(menus []widgets.Item) error {
	h.Menu.RenderGlobalMenus()
	h.Menu.RenderMenu(menus, true)
	return nil
}

func (h *Header) SetMetadata(metadata models.Metadata) {
	h.Metadata.Clear()
	subsKey := fmt.Sprintf("[%s]%s:", "cadetblue", "SubscriptionId")
	subsValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.AzureSubscriptionId != "" {
		subsValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.AzureSubscriptionId)
	}
	h.Metadata.SetMapKeyValue(subsKey, subsValue)

	clientIdKey := fmt.Sprintf("[%s]%s:", "cadetblue", "ClientId")
	clientIdValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.AzureClientId != "" {
		clientIdValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.AzureClientId)
	}
	h.Metadata.SetMapKeyValue(clientIdKey, clientIdValue)

	tenantIdKey := fmt.Sprintf("[%s]%s:", "cadetblue", "TenantId")
	tenantIdValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.AzureTenantId != "" {
		tenantIdValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.AzureTenantId)
	}
	h.Metadata.SetMapKeyValue(tenantIdKey, tenantIdValue)

	h.Metadata.DrawMapView()
}