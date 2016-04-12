package main

import (
	termbox "github.com/nsf/termbox-go"
	. "github.com/telyn/form"
	"github.com/telyn/form/box"
	"log"
)

var buttons *ButtonsField

func draw() bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	buttons.DrawInto(&box.TermBox{}, 0, 0)
	termbox.Flush()
	return true

}

func main() {
	buttons = NewButtonsField([]Button{{
		Action: func() {
			log.Printf("Yeah!")
		},
		Text: "yeah!",
	}, {
		Action: func() {
			log.Printf("No!")
		},
		Text: "no!",
	}})

	buttons.Focus(true)
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	if draw() {
		log.Printf("drew ok\r\n")
	loop:
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc, termbox.KeyCtrlC:
					break loop
				default:
					if ev.Ch != '\x00' {
						buttons.ReceiveRune(ev.Ch)
					} else {
						buttons.ReceiveKey(ev.Key)
					}
					draw()
				}
			case termbox.EventResize:
				draw()
			}
		}
		termbox.Close()
	} else {
		log.Printf("Didn't draw correctly. weird\r\n")
	}

}
