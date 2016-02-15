// TextField is a form field for a single line of text
package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	"log"
)

type TextField struct {
	viewOffsetX int
	oldCursorX  int
	cursorX     int
	width       int
	value       []rune
}

func NewTextField(width int, value []rune) *TextField {
	return &TextField{
		width:   width,
		cursorX: -1,
	}
}

func (f *TextField) Draw() *box.CellsBox {
	log.Printf("Drawing. len(value):%v viewOffsetX:%v cursorX:%v", len(f.value), f.viewOffsetX, f.cursorX)
	box := box.New(f.width, 1)
	fg := termbox.ColorRed
	if f.Validate() {
		fg = termbox.ColorGreen
	}
	for i := 0; i < f.width; i++ {
		idx := f.viewOffsetX + i
		ch := ' '

		if idx < len(f.value) {
			ch = f.value[idx]
		}
		bg := termbox.ColorBlue
		if i == f.cursorX {
			bg = termbox.ColorWhite
		}
		box.SetCell(i, 0, ch, fg, bg)
	}
	return box
}

func (f *TextField) DrawInto(box box.Box, x, y int) {
	fieldBox := f.Draw()
	fieldBox.DrawInto(box, x, y)
}

func (f *TextField) GetCursor() (x, y int) {
	return f.viewOffsetX, 0
}

func (f *TextField) GetValue() string {
	return string(f.value)
}

func (f *TextField) Focus(hasFocus bool) {
	if hasFocus {
		f.cursorX = f.oldCursorX
	} else {
		f.oldCursorX = f.cursorX
		f.cursorX = -1

	}
}

func (f *TextField) SetCursor(x, y int) {
	if x < 0 && f.viewOffsetX > 0 {
		f.cursorX = 0
		f.viewOffsetX--
	} else if x > f.width {
		f.cursorX = f.width - 1
		f.viewOffsetX++
	} else {
		f.cursorX = x
	}
}

func (f *TextField) removeChar(offset int) {
	copy(f.value[offset+1:], f.value[offset:])
	f.value = f.value[:len(f.value)-1]
}

func (f *TextField) ReceiveKey(key termbox.Key) {
	switch key {
	case termbox.KeyArrowLeft:
		f.SetCursor(f.cursorX-1, 0)
	case termbox.KeyArrowRight:
		f.SetCursor(f.cursorX+1, 0)
	case termbox.KeyBackspace:
		f.removeChar(f.viewOffsetX + f.cursorX - 1)
		f.SetCursor(f.cursorX-1, 0)
	case termbox.KeyDelete:
		f.removeChar(f.viewOffsetX + f.cursorX)

	}
}

func (f *TextField) Size() (int, int) {
	return f.width, 1
}

func (f *TextField) ReceiveRune(ch rune) {
	pos := f.viewOffsetX + f.cursorX
	runes := make([]rune, len(f.value)+1)
	if len(f.value) > 0 {
		for i, _ := range runes {
			if i < pos {
				runes[i] = f.value[i]
			} else if i > pos {
				runes[i] = f.value[i-1]
			}
		}

	}
	runes[pos] = ch
	f.value = runes
	f.SetCursor(f.cursorX+1, 0)
	log.Printf("runes:%v pos:%v cursor:%v", len(f.value), pos, f.cursorX)
}

func (f *TextField) Validate() bool {
	return true
}