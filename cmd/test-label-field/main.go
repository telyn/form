package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/telyn/form"
	"github.com/telyn/form/box"
	"os"
)

var labelField Field

func draw() bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	labelField.DrawInto(&box.TermBox{}, 0, 0)

	termbox.Flush()
	return true
}

func main() {
	textField := NewTextField(12, []rune("jean-michel jarre was a fraud"), func(val string) bool {
		return true
	})
	labelField = Label(textField,
		"was jean-michel jarre a fraud? discuss in one line or else. YOU MUST DISCUSS IT OR ELSE ELEVEN YEARS DUNGEON AAAAAA")

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
