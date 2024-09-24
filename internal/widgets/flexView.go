package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
)

type Flex struct {
	*tview.Flex
}

func NewFlex() *Flex {
	f := &Flex{
		Flex:	tview.NewFlex(),
	}
	f.SetBorder(true)
	return f
}

func (f *Flex) AddItemX(primitive tview.Primitive, fixedSize, proportion int, focus bool) {
	f.AddItemc(primitive, fixedSize, proportion, focus)
}

func (f *Flex) FullScreen(on bool) {
	f.SetFullScreen(on)
}

func (f *Flex) SetTitle(title, a string) {
	if a != "" {
		f.SetTitle(fmt.Sprintf(" [::b][%s]%s(%s) ", utils.ColorT7, strings.ToUpper(title), a))
	} else {
		f.SetTitle(fmt.Sprintf(" [::b][%s]%s ", utils.ColorT7, strings.ToUpper(title)))
	}
}

func (f *Flex) ClearFlex() {
	f.Clear()
}