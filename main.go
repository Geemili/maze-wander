package main

import (
	"encoding/csv"
	"github.com/nsf/termbox-go"
	"io"
	"os"
)

func checkCollision(tiles [][]rune, x, y int) bool {
	if x < 0 || y < 0 || y > len(tiles) || x > len(tiles[0]) {
		return true
	}
	switch tiles[y][x] {
	case '.':
		return false
	default:
		return true
	}
}

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
			nextx, nexty := playerx, playery
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
			if !checkCollision(tiles, nextx, nexty) {
				playerx, playery = nextx, nexty
			}
		}
	}
}
