package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(20, 20, '@', termbox.ColorYellow, termbox.ColorDefault)
		termbox.Flush()
	}
}
