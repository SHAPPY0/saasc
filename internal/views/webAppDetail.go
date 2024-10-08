package views

import (
	"fmt"
	"strings"
	"github.com/shappy0/saasc/internal/widgets"
	"github.com/shappy0/saasc/internal/models"
	"github.com/shappy0/saasc/internal/utils"
)

var TitleWebAppDetail = "web app detail"

type WebAppDetail struct {
	*widgets.Flex
	Title			string
	FirstRow		*widgets.Flex
	SecondRow		*widgets.Flex
	Essential		*widgets.Container
	SiteConfig		*widgets.Container
	Networking		*widgets.Container
	Data			models.WebApp
}

func NewWebAppDetail() *WebAppDetail {
	wad := WebAppDetail{
		Flex:		widgets.NewFlex(),
		Title:		TitleWebAppDetail,
		FirstRow:	widgets.NewFlex(),
		SecondRow:	widgets.NewFlex(),
		Essential:	widgets.NewContainer("Essential", 0, 1),
		SiteConfig:	widgets.NewContainer("siteConfig", 0, 1),
		Networking:	widgets.NewContainer("Networking", 0, 2),
	}
	wad.FlexRow()
	wad.SetTitle(TitleWebAppDetail)
	wad.AddItemX(wad.FirstRow, 0, 1, false)
	wad.AddItemX(wad.SecondRow, 0, 1, false)
	return &wad
}

func (wad *WebAppDetail) GetTitle() string {
	return wad.Title
}

func (wad *WebAppDetail) UpdateData(data models.WebApp) {
	wad.Data = data
	wad.SetTitleX(TitleWebAppDetail, wad.Data.Name)
	wad.DrawFirstRow()
	wad.DrawSecondRow()
}

func (wad *WebAppDetail) DrawFirstRow() {
	wad.FirstRow.Border(false)
	wad.FirstRow.ClearFlex()
	wad.Essential.Clear()
	wad.SetEssentialValues()
	wad.FirstRow.AddItemX(wad.Essential, 0, 1, false)
	wad.SiteConfig.Clear()
	wad.SetSiteConfigValues()
	wad.FirstRow.AddItemX(wad.SiteConfig, 0, 1, false)
}

func (wad *WebAppDetail) SetEssentialValues() {
	wad.Essential.SetKeyValue("Name:", wad.Data.Name)
	wad.Essential.SetKeyValue("Status:", wad.Data.State)
	wad.Essential.SetKeyValue("AvailabilityState:", wad.Data.UsageState)
	wad.Essential.SetKeyValue("ResourceGroup:", wad.Data.ResourceGroup)
	wad.Essential.SetKeyValue("Location:", wad.Data.Location)
	wad.Essential.SetKeyValue("DefaultHost:", wad.Data.DefaultHostName)
	wad.Essential.SetKeyValue("AppServicePlan:", utils.ParseResourceId(wad.Data.ServerFarmID, "serverfarms"))
	wad.Essential.SetKeyValue("Kind:", wad.Data.Kind)
	wad.Essential.SetKeyValue("LasModifiedAt:", utils.DateTimeToStr(wad.Data.LastModifiedTimeUTC))
}

func (wad *WebAppDetail) SetSiteConfigValues() {
	alwaysOn := "false"
	if wad.Data.SiteConfig.AlwaysOn {
		alwaysOn = "true"
	}
	wad.SiteConfig.SetKeyValue("AlwaysOn:", alwaysOn)
	acrUseMI := "false"
	if wad.Data.SiteConfig.AcrUseManagedIdentityCreds {
		acrUseMI = "true"
	}
	wad.SiteConfig.SetKeyValue("AcrUseManagedIdentityCreds:", acrUseMI)
	imageInfo := getImageInfo(wad.Data.SiteConfig.LinuxFxVersion)
	wad.SiteConfig.SetKeyValue("ImageType:", imageInfo[0])
	wad.SiteConfig.SetKeyValue("Image:", imageInfo[1])
	wad.SiteConfig.SetKeyValue("MinElasticInstaneCount:", utils.IntToStr(int(wad.Data.SiteConfig.MinimumElasticInstanceCount)))
	wad.SiteConfig.SetKeyValue("NumberOfWorkers:", utils.IntToStr(int(wad.Data.SiteConfig.NumberOfWorkers)))
}

func (wad *WebAppDetail) DrawSecondRow() {
	wad.SecondRow.Border(false)
	wad.SecondRow.ClearFlex()
	wad.Networking.Clear()
	wad.SecondRow.FlexRow()
	wad.SetNetworkingValues()
	wad.SecondRow.AddItemX(wad.Networking, 0, 1, false)
}

func (wad *WebAppDetail) SetNetworkingValues() {
	wad.Networking.SetKeyValue("Vnet/Subnet:", fmt.Sprintf("%s/%s", utils.ParseResourceId(wad.Data.VirtualNetworkSubnetID, "virtaulNetworks"), utils.ParseResourceId(wad.Data.VirtualNetworkSubnetID, "subnets")))
	enabledHostNames := ""
	for i := 0; i < len(wad.Data.EnabledHostNames); i++ {
		enabledHostNames += (string(*wad.Data.EnabledHostNames[i]) + " ")
	}
	wad.Networking.SetKeyValue("PublicNetworkAccess:", wad.Data.PublicNetworkAccess)
	vnetImagePullEnabled := "false"
	if wad.Data.VnetImagePullEnabled {
		vnetImagePullEnabled = "true"
	}
	wad.Networking.SetKeyValue("VnetImagePullEnabled:", vnetImagePullEnabled)
	wad.Networking.SetKeyValue("Identity:", wad.Data.Identity.Type)
	if wad.Data.Identity.Type == "UserAssigned" {
		umiName := ""
		for k, _ := range wad.Data.Identity.UserAssignedIdentities {
			umiName = utils.ParseResourceId(k, "userAssignedIdentites")
		}
		wad.Networking.SetKeyValue("UserManagedIdentit:", umiName)
	}
	wad.Networking.SetKeyValue("EnabledHostNames:", enabledHostNames)
	wad.Networking.SetKeyValue("KeyVaultReferenceIdentity:", utils.ParseResourceId(wad.Data.KeyVaultReferenceIdentity, "userAssignedIdentites"))
	wad.Networking.SetKeyValue("OutboundIPAddresses:", wad.Data.OutboundIPAddresses)
	wad.Networking.SetKeyValue("PossibleOutboundIPAddresses:", wad.Data.PossibleOutboundIPAddresses)
}

func getImageInfo(image string) []string {
	var imageInfo []string
	if image != "" {
		imageInfo = strings.Split(image, "|")
	}
	return imageInfo
}