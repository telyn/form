package box

import (
	"github.com/cheekybits/is"
	termbox "github.com/nsf/termbox-go"
	"testing"
)

func TestFill(t *testing.T) {
	is := is.New(t)

	box := New(4, 4)
	defaultCell := termbox.Cell{
		Ch: ' ',
		Fg: termbox.ColorDefault,
		Bg: termbox.ColorDefault,
	}
	box.Fill(&defaultCell)

	for _, c := range box.CellBuffer() {
		is.Equal(' ', c.Ch)
	}
}
