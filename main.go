package main

import (
	"encoding/csv"
	"github.com/geemili/maze-wander/game"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

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

	eventQueue := make(chan termbox.Event) // So that we can have async keyboard
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	dialog := make(chan string)
	conversation := game.Conversation{
		[]string{
			"You're gonna have a bad time...",
			"Are you sure you want to continue?",
		}, 0,
	}

	tileMap, playerx, playery := loadMap("assets/map.csv")
	g := game.Game{tileMap, []*game.Entity{}}

	player := &game.Entity{playerx, playery, 1, 1}
	g.AddEntity(player)

	messageBox := game.NewMessageBox()

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
				case ev.Ch == 'z':
					conversation.Act(dialog)
				}
				if g.WorldMap.GetTileAt(nextx, nexty) == 1 {
					player.X, player.Y = nextx, nexty
				}
			}
		case text := <-dialog:
			messageBox.Message = text
		default: // Do main loop

			g.Render()
			messageBox.Render()

			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Return the tilemap and the position of the player
func loadMap(path string) (game.TileMap, int, int) {
	tilesFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(tilesFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	playerx, playery := 0, 0

	w, h := len(records), len(records[0])

	tileMap := game.TileMap{make([]int, w*h), w, h}
	for y, row := range records {
		for x, cell := range row {
			pos := y*w + x
			switch cell {
			case "-1":
				tileMap.Tiles[pos] = 0
			case "178":
				tileMap.Tiles[pos] = 2
			case "210":
				tileMap.Tiles[pos] = 3
			case "59":
				tileMap.Tiles[pos] = 4
			case "4":
				playerx = x
				playery = y
				fallthrough
			case "226":
				tileMap.Tiles[pos] = 1
			}
		}
	}
	return tileMap, playerx, playery
}
