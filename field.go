package form

import (
	termbox "github.com/nsf/termbox-go"
)

type Field interface {
	Widget
	Focus(hasFocus bool)
	ReceiveKey(termbox.Key)
	ReceiveRune(ch rune)
	Validate() bool
}
