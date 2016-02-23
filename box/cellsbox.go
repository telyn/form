package box

import (
	termbox "github.com/nsf/termbox-go"
)

func New(width int, height int) *CellsBox {
	return &CellsBox{
		cells: make([]termbox.Cell, width*height),
		width: width,
	}

}

// cellsBox is a Box which can Draw its contents into Boxes
type CellsBox struct {
	cells []termbox.Cell
	width int
}

func (c *CellsBox) CellBuffer() []termbox.Cell {
	return c.cells
}

func (c *CellsBox) Fill(cell *termbox.Cell) {
	for i, _ := range c.cells {
		c.cells[i] = *cell
	}
}

func (c *CellsBox) Size() (w, h int) {
	return c.width, len(c.cells) / c.width
}

func (c *CellsBox) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	offset := (y * c.width) + x
	c.cells[offset] = termbox.Cell{Ch: ch, Fg: fg, Bg: bg}
}

func (c *CellsBox) GetCell(x, y int) *termbox.Cell {
	return &c.cells[y*c.width+x]
}

// DrawInto inserts one box inside another at a given point
func (c *CellsBox) DrawInto(c1 Box, x, y int) error {
	w, h := c.Size()
	w1, h1 := c1.Size()

	// if the Box we're drawing is going to go outside the bounds of the one we're drawing into, fail!
	if x+w > w1 || y+h > h1 {
		return OutOfBoundsError{x + w, y + h, w1, h1}
	}

	counter := 0
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			cell := c.cells[counter]
			c1.SetCell(x+i, y+j, cell.Ch, cell.Fg, cell.Bg)
			counter++
		}
	}
	return nil

}
