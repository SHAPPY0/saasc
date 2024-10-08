package widgets

import (
	"fmt"
	"strings"
	"github.com/rivo/tview"
	"github.com/shappy0/saasc/internal/utils"
)

type TextView struct {
	*tview.TextView
	Title		string
}

func NewTextView(title string) *TextView {
	tv := &TextView{
		TextView:	tview.NewTextView(),
		Title:		title,
	}
	tv.SetBorderX(true)
	tv.SetScrollable(true)
	tv.SetTextVTitle("", "")
	return tv
}

func (tv *TextView) SetTitleName(name string) {
	tv.Title = name
}

func (tv *TextView) SetTextAlignX(align string) {
	alignX := tview.AlignLeft
	if align == "AlignCenter" {
		alignX = tview.AlignCenter
	} else if align == "AlignLeft" {
		alignX = tview.AlignLeft
	} else if align == "AlignRight" {
		alignX = tview.AlignRight
	}
	tv.SetTextAlign(alignX)
}

func (tv *TextView) SetBorderX(b bool) {
	tv.SetBorder(b)
}

func (tv *TextView) SetTextVTitle(a, b string) {
	if a != "" && b != "" {
		tv.SetTitle(fmt.Sprintf(" [::b][%s]%s (%s/%s) ", utils.ColorT70d5bf, strings.ToUpper(tv.Title), a, b))
	} else {
		tv.SetTitle(fmt.Sprintf(" [::b][%s]%s ", utils.ColorT70d5bf, strings.ToUpper(tv.Title)))
	}
}

func (tv *TextView) SetTextX(text string) {
	tv.SetText(text)
}

func (tv *TextView) ClearX() {
	tv.Clear()
}