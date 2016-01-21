package main

import (
	"encoding/csv"
	"github.com/nsf/termbox-go"
	"io"
	"os"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var playerx, playery int = 20, 15
	var tiles [][]rune

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

	tiles = make([][]rune, len(records))
	for y, row := range records {
		tileRow := make([]rune, len(row))
		for x, _ := range row {
			switch row[x] {
			case "-1":
				tileRow[x] = ' '
			case "178":
				tileRow[x] = '+'
			case "210":
				tileRow[x] = '-'
			case "59":
				tileRow[x] = '|'
			case "4":
				playerx = x
				playery = y
				fallthrough
			case "226":
				tileRow[x] = '.'
			}
		}
		tiles[y] = tileRow
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
