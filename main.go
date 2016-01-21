package main

import (
	"encoding/csv"
	"github.com/geemili/maze-wander/game"
	"github.com/nsf/termbox-go"
	"io"
	"os"
	"time"
)

func dialogBox(words string) {
	for i := 1; i < 80; i++ {
		termbox.SetCell(i, 30, '-', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(i, 36, '-', termbox.ColorWhite, termbox.ColorDefault)
	}
	for i := 31; i < 36; i++ {
		termbox.SetCell(0, i, '|', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(80, i, '|', termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.SetCell(0, 30, '+', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(80, 30, '+', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(0, 36, '+', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(80, 36, '+', termbox.ColorWhite, termbox.ColorDefault)

	for idx, char := range words {
		termbox.SetCell(1+idx, 31, char, termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.SetCursor(1+len(words), 31)
}

func typeDialog(dialogOutput chan string, text string) {
	for idx, _ := range text {
		dialogOutput <- text[:idx+1]
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Stuff for the dialog box
	var dialogDisplay string = ""

	eventQueue := make(chan termbox.Event) // So that we can have async keyboard
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	dialogQueue := make(chan string)
	go typeDialog(dialogQueue, "Somebody wanted to have a bad time...")

	var tilesFile io.Reader
	tilesFile, err = os.Open("assets/map.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(tilesFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	playerx, playery := 0, 0

	tiles := make([][]int, len(records))
	for y, row := range records {
		tileRow := make([]int, len(row))
		for x, _ := range row {
			switch row[x] {
			case "-1":
				tileRow[x] = 0
			case "178":
				tileRow[x] = 2
			case "210":
				tileRow[x] = 3
			case "59":
				tileRow[x] = 4
			case "4":
				playerx = x
				playery = y
				fallthrough
			case "226":
				tileRow[x] = 1
			}
		}
		tiles[y] = tileRow
	}

	g := game.NewGame(len(tiles), len(tiles[0]))
	for y, row := range tiles {
		for x, tile := range row {
			g.WorldMap.SetTileAt(x, y, tile)
		}
	}

	player := &game.Entity{playerx, playery, 1, 1}
	g.AddEntity(player)

loop:
	for {
		select { // Multiplex between channells
		case ev := <-eventQueue: // Handle termbox events
			if ev.Type == termbox.EventKey {
				// Handle keyboard input
				nextx, nexty := player.X, player.Y
				switch {
				case ev.Key == termbox.KeyEsc:
					break loop
				case ev.Key == termbox.KeyArrowLeft, ev.Ch == 'a':
					nextx -= 1
				case ev.Key == termbox.KeyArrowRight, ev.Ch == 'd':
					nextx += 1
				case ev.Key == termbox.KeyArrowUp, ev.Ch == 'w':
					nexty -= 1
				case ev.Key == termbox.KeyArrowDown, ev.Ch == 's':
					nexty += 1
				}
				if g.WorldMap.GetTileAt(nextx, nexty) == 1 {
					player.X, player.Y = nextx, nexty
				}
			}
		case text := <-dialogQueue:
			dialogDisplay = text
		default: // Do main loop
			game.Render(*g)
			dialogBox(dialogDisplay)
			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
