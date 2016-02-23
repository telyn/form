package box

import (
	termbox "github.com/nsf/termbox-go"
)

type OutOfBoundsError struct {
	X         int
	Y         int
	BoxWidth  int
	BoxHeight int
}

func (e OutOfBoundsError) Error() string {
	return "Went out of bounds!"
}

type Box interface {
	Fill(*termbox.Cell)
	Size() (int, int)
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	GetCell(x, y int) *termbox.Cell
}
