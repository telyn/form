package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/telyn/form"
	"github.com/telyn/form/box"
	"log"
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
	field2 := NewTextField(18, []rune(""), func(val string) bool { return true })
	form = NewForm([]Field{
		Label(field, "First text field, yeah!"),
		Label(field1, "Second text field, accompanied by a whole load more words. Like at least nine words but probably more yeah. That's how many words we're going for anyway. This ought to do"),
		Label(field2, "Third text field oh boy"),
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
			ev := termbox.PollEvent()
			if !form.HandleEvent(&ev) {
				log.Print("Exiting because HandleEvent said so")
				break loop
			}
			draw()
		}
	}

}
