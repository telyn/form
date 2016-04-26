package form

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/telyn/form/box"
	//"log"
)

// TextField is a form field for a single line of text
type MaskedTextField struct {
	viewOffsetX int
	oldCursorX  int
	cursorX     int
	width       int
	value       []rune
	validFn     func(string) (string, bool)
}

// NewTextField creates a new text field (quel surprise). Width is absolute -
// come hell or high water, the displayed width will be width. value is the
// initial value and need not be capacious
func NewMaskedTextField(width int, value []rune, validateFn func(string) (string, bool)) *MaskedTextField {
	return &MaskedTextField{
		width:   width,
		cursorX: -1,
		validFn: validateFn,
	}
}

func (f *MaskedTextField) Draw() *box.CellsBox {
	//log.Printf("Drawing. len(value):%v viewOffsetX:%v cursorX:%v", len(f.value), f.viewOffsetX, f.cursorX)
	box := box.New(f.width, 1)
	fg := termbox.ColorRed | termbox.AttrUnderline
	if _, ok := f.Validate(); ok {
		fg = termbox.ColorGreen | termbox.AttrUnderline
	}
	for i := 0; i < f.width; i++ {
		idx := f.viewOffsetX + i
		ch := ' '

		if idx < len(f.value) {
			ch = '*'
		}
		bg := termbox.ColorDefault
		if i == f.cursorX {
			bg = termbox.ColorWhite
		}
		box.SetCell(i, 0, ch, fg, bg)
	}
	return box
}

func (f *MaskedTextField) DrawInto(box box.Box, x, y int) {
	fieldBox := f.Draw()
	fieldBox.DrawInto(box, x, y)
}

func (f *MaskedTextField) GetCursor() (x, y int) {
	return f.viewOffsetX, 0
}

func (f *MaskedTextField) GetValue() string {
	return string(f.value)
}

func (f *MaskedTextField) Focus(hasFocus bool) {
	//log.Printf("focus(%v)", hasFocus)
	if hasFocus {
		f.cursorX = f.oldCursorX
	} else {
		f.oldCursorX = f.cursorX
		f.cursorX = -1
	}
}

func (f *MaskedTextField) HandleResize(x, y int) {
	return // leave me alone I am a fixed size
}

func (f *MaskedTextField) SetCursor(x, y int) {
	if x < 0 && f.viewOffsetX > 0 {
		f.cursorX = 0
		f.viewOffsetX--
	} else if x >= f.width {
		f.cursorX = f.width - 1
		if x+f.viewOffsetX < len(f.value)+1 {
			f.viewOffsetX++
		}
	} else if x >= 0 {
		if x+f.viewOffsetX < len(f.value)+1 {
			f.cursorX = x
		}
	} else {
		f.cursorX = 0
	}
}

func (f *MaskedTextField) removeChar(offset int) {
	if offset > len(f.value)-1 {
		return
	}
	if offset < len(f.value)-1 {
		//log.Printf("Moving %d (%c) to %d", offset+1, f.value[offset+1], offset)
		copy(f.value[offset:], f.value[offset+1:])
	}
	//log.Printf("Taking %d (%c) off the end", offset, f.value[offset])
	f.value = f.value[:len(f.value)-1]
}

func (f *MaskedTextField) ReceiveKey(key termbox.Key) {
	switch key {
	case termbox.KeyArrowLeft:
		f.SetCursor(f.cursorX-1, 0)
	case termbox.KeyArrowRight:
		f.SetCursor(f.cursorX+1, 0)
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		f.SetCursor(f.cursorX-1, 0)
		f.removeChar(f.viewOffsetX + f.cursorX)
	case termbox.KeySpace:
		f.ReceiveRune(' ')
	case termbox.KeyDelete:
		f.removeChar(f.viewOffsetX + f.cursorX)

	}
}

func (f *MaskedTextField) Size() (int, int) {
	return f.width, 1
}

func (f *MaskedTextField) ReceiveRune(ch rune) {

	pos := f.viewOffsetX + f.cursorX
	runes := make([]rune, len(f.value)+1)
	if len(f.value) > 0 {
		// I *could* use copy here, but this works and I don't want to
		// mess with it
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
	//log.Printf("runes:%v pos:%v cursor:%v", len(f.value), pos, f.cursorX)
}

func (f *MaskedTextField) Validate() (string, bool) {
	return f.validFn(string(f.value))
}

func (f *MaskedTextField) Value() string {
	return string(f.value)
}
