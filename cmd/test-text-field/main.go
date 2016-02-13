package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/telyn/form"
	"github.com/telyn/form/box"
	"os"
)

var textField *TextField

func draw() bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCell(0, 0, 'H', 0, 0)
	termbox.SetCell(0, 0, 'e', 0, 0)
	termbox.SetCell(0, 0, 'y', 0, 0)
	textField.DrawInto(&box.TermBox{}, 4, 0)

	termbox.Flush()
	return true
}

func main() {
	textField = NewTextField(16, make([]rune, 0))
	textField.Focus(true)

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
				case termbox.KeyEsc, termbox.KeyCtrlC:
					break loop
				default:
					if ev.Ch != '\x00' {
						textField.ReceiveRune(ev.Ch)
					} else {
						textField.ReceiveKey(ev.Key)
					}
					draw()
				}

			case termbox.EventResize:
				draw()
			}
		}
		termbox.Close()
	} else {
		fmt.Printf("Didn't draw correctly\r\n")
	}

	fmt.Printf("Hey %s\r\n", textField.GetValue())
}
