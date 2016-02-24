package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/telyn/form"
	"github.com/telyn/form/box"
	"os"
)

var form *Form

func draw() bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	form.DrawInto(&box.TermBox{}, 0, 0)

	termbox.Flush()
	return true
}

func main() {
	field := NewTextField(12, []rune(""), func(val string) bool {
		return false
	})
	field1 := NewTextField(24, []rune(""), func(val string) bool { return true })
	form = NewForm([]Field{
		Label(field, "First text field, yeah!"),
		Label(field1, "Second text field, accompanied by a whole load more words. Like at least nine words but probably more yeah. That's how many words we're going for anyway. This ought to do"),
	})
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		fmt.Println(err)
		return
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
						form.ReceiveRune(ev.Ch)
					} else {
						form.ReceiveKey(ev.Key)
					}
				}
			case termbox.EventResize:
				form.HandleResize(ev.Width, ev.Height)
			}
		}
	}

}
