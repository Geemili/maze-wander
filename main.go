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

	var playerx, playery int = 20, 15
	tiles := [][]rune{
		{'+', '-', '-', '-', '-', '+'},
		{'|', '.', '.', '.', '.', '|'},
		{'|', '.', '.', '.', '.', '|'},
		{'|', '.', '.', '.', '.', '|'},
		{'|', '.', '.', '.', '.', '|'},
		{'|', '.', '.', '.', '.', '|'},
		{'+', '-', '-', '-', '-', '+'},
	}

loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		for y, row := range tiles {
			for x, tile := range row {
				termbox.SetCell(x, y, tile, termbox.ColorGreen, termbox.ColorDefault)
			}
		}
		termbox.SetCell(playerx, playery, '@', termbox.ColorYellow, termbox.ColorDefault)
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowLeft:
				playerx -= 1
			case termbox.KeyArrowRight:
				playerx += 1
			case termbox.KeyArrowUp:
				playery -= 1
			case termbox.KeyArrowDown:
				playery += 1
			}
		}
	}
}
