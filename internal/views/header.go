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
	MetadataContent	models.Metadata
	Menu 			*widgets.Menu
	
}

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
	h.MetadataContent = models.Metadata{
		ResourceGroup:			"-",
		AzureSubscriptionId:	config.AzureSubscriptionId,
		AzureClientId:			config.AzureClientId,
		AzureTenantId:			config.AzureTenantId,
	}
	h.RenderMetadata()
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

func (h *Header) RenderMetadata() {
	h.Metadata.Clear()
	metadata := h.MetadataContent
	rgKey := fmt.Sprintf("[%s]%s:", "cadetblue", "ResourceGroup")
	rgValue := fmt.Sprintf("[%s]%s\n", "DimGray", "-")	
	if metadata.ResourceGroup != "" {
		rgValue = fmt.Sprintf("[%s]%s\n", "DimGray", metadata.ResourceGroup)
	}
	h.Metadata.SetMapKeyValue(rgKey, rgValue)

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

	h.Metadata.DrawMapView()
}

func (h *Header) UpdateMetadata(key, value string, rerender bool) {
	if key == "ResourceGroup" {
		h.MetadataContent.ResourceGroup = value
	}
	if rerender {
		h.RenderMetadata()
	}
}