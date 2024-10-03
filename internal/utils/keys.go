package utils

import (
	"github.com/gdamore/tcell/v2"
)

type Key struct {
	Key 			tcell.Key
	KeyRune			rune
	KeyLabel		string
	KeyDescription	string
}

var (
	ExistKey = Key{
		Key:			tcell.KeyCtrlC,
		KeyLabel:		"Ctrl + c",
		KeyDescription:	"Control and C",
	}
	EnterKey = Key{
		Key:			tcell.KeyEnter,
		KeyLabel:		"enter",
		KeyDescription:	"Enter Key",
	}
	EscKey = Key{
		Key:			tcell.KeyEsc,
		KeyLabel:		"Esc",
		KeyDescription:	"Esc Key",
	}
	TabKey = Key{
		Key:			tcell.KeyTAB,
		KeyLabel:		"Tab",
		KeyDescription:	"Tab Key",
	}
	RuneKey = Key{
		Key:			tcell.KeyRune,
		KeyLabel:		"rune keys",
		KeyDescription:	"Rune Keys",
	}
)

var KeyBindings = []Key{
	ExistKey,
}