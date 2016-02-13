package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/telyn/form"
	"os"
)

var dot, greenPlus termbox.Cell

func draw() bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	dots := NewCellsBox([]termbox.Cell{dot, dot, dot, dot, dot,
		dot, dot, dot, dot, dot,
		dot, dot, dot, dot, dot,
		dot, dot, dot, dot, dot,
	}, 5)

	greenPlus := NewCellsBox([]termbox.Cell{greenPlus, greenPlus, greenPlus, greenPlus}, 2)

	if err := greenPlus.DrawInto(dots, 1, 2); err != nil {
		termbox.Close()
		fmt.Fprintf(os.Stderr, "%v1\r\n", err)
		return false
	}
	if err := dots.DrawInto(&TermBox{}, 0, 0); err != nil {
		termbox.Close()
		fmt.Fprintf(os.Stderr, "%v2\r\n", err)
		return false
	}

	termbox.SetCell(0, 0, '+', termbox.ColorRed, termbox.ColorDefault)
	termbox.SetCell(1, 0, '+', termbox.ColorRed, termbox.ColorDefault)
	termbox.SetCell(7, 0, '+', termbox.ColorRed, termbox.ColorDefault)

	termbox.Flush()
	return true
}

func main() {
	dot = termbox.Cell{Ch: '.', Fg: termbox.ColorDefault, Bg: termbox.ColorDefault}
	greenPlus = termbox.Cell{Ch: '+', Fg: termbox.ColorGreen, Bg: termbox.ColorDefault}
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	if draw() {
		fmt.Fprintf(os.Stderr, "drew ok\r\n")
	loop:
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					break loop
				}
			case termbox.EventResize:
				draw()
			}
		}
		termbox.Close()
	} else {
		fmt.Printf("Didn't draw correctly\r\n")
	}
}
