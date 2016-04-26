package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
)

type Button struct {
	Action func()
	Text   string
}

func (b *Button) Width() int {
	return len(b.Text)
}

type ButtonsField struct {
	buttons           []Button
	currentlySelected int
}

func NewButtonsField(bs []Button) (bf *ButtonsField) {
	bf = new(ButtonsField)

	bf.buttons = bs
	bf.currentlySelected = -1
	return
}

func (bf *ButtonsField) GetCursor() (x, y int) {
	return -1, 0
}

func (bf *ButtonsField) SetCursor(x, y int) {
}

func (bf *ButtonsField) DrawInto(box box.Box, x, y int) {
	// do something to centre it.
	curX := x
	for b, butt := range bf.buttons {
		for _, ch := range butt.Text {
			curX++
			if b == bf.currentlySelected {
				box.SetCell(curX, y, ch, termbox.ColorWhite, termbox.ColorGreen)
			} else {
				box.SetCell(curX, y, ch, termbox.ColorWhite, termbox.ColorBlue)
			}
		}
		curX += 2
	}
}

func (bf *ButtonsField) Focus(hasFocus bool) {
	if hasFocus {
		bf.currentlySelected = 0
	} else {
		bf.currentlySelected = -1
	}
}

func (bf *ButtonsField) GetSelected() Button {
	return bf.buttons[bf.currentlySelected]
}

func (bf *ButtonsField) HandleResize(w, h int) {
	// no action necessary.
}

func (bf *ButtonsField) ReceiveKey(key termbox.Key) {
	switch key {
	case termbox.KeyEnter:
		bf.GetSelected().Action()
	case termbox.KeyArrowLeft:
		if bf.currentlySelected > 0 {
			bf.currentlySelected--
		}
	case termbox.KeyArrowRight:
		if bf.currentlySelected < len(bf.buttons)-1 {
			bf.currentlySelected++
		}

	}
}

func (bf *ButtonsField) ReceiveRune(ch rune) {
	// i just don't care about runes
}

func (bf *ButtonsField) Size() (x, y int) {
	//map strlen(.Text) +1 over Buttons, add 1
	length := 0
	for _, b := range bf.buttons {
		length += len(b.Text) + 1
	}
	return length + 1, 1
}

func (bf *ButtonsField) Validate() (string, bool) {
	return "", true
}
func (bf *ButtonsField) Value() string {
	return ""
}
