package form

import (
	termbox "github.com/nsf/termbox-go"
)

type Field interface {
	Widget
	Focus(bool)
	ReceiveKey(termbox.Key)
	ReceiveRune(rune)
	Validate() (string, bool)
	Value() string
}
