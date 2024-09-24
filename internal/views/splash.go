package views

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
	"github.com/shappy0/saasc/internal/config"
)

var TitleSplash = "splash"

type Splash struct {
	*tview.Flex
	Title 		string
	Version		string
}

var Logo = []string{
	` ___   __     __   ___  ___` ,
	`/ __) /__\   /__\ / __)/ __)`,
	`\__ \/(__)\ /(__)\\__ ( (__ `,
	`(___(__)(__|__)(__|___/\___)`,
}

func NewSplash(config *config.Conf) *Splash {
	s := Splash{
		Flex:		tview.NewFlex(),
		Title:		TitleSplash,
		Version:	config.Version,
	}
	s.SetDirection(tview.FlexRow)
	LogoV := tview.NewTextView()
	LogoV.SetDynamicColors(true)
	LogoV.SetTextAlign(tview.AlignCenter)
	s.DrawLogo(LogoV)

	Version := tview.NewTextView()
	Version.SetDynamicColors(true)
	Version.SetTextAlign(tview.AlignCenter)

	s.DrawVersion(Version, s.Version)

	s.AddItem(LogoV, 10, 1, false)
	s.AddItem(Version, 1, 1, false)
	return &s
}

func (s *Splash) GetTitle() string {
	return s.Title
}


func (s *Splash) DrawLogo(t *tview.TextView) {
	LogoV := strings.Join(Logo, fmt.Sprintf("\n[%s::b]", "#cccccc"))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
			strings.Repeat("\n", 2),
			"#cccccc",
			LogoV)
}

func (s *Splash) DrawVersion(t *tview.TextView, Version string) {
	fmt.Fprintf(t, "[%s::b]Version: [orange::b]%s", "#cccccc", Version)
}