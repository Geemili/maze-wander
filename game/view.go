package game

import (
	"github.com/nsf/termbox-go"
)

var entityCharMap map[int]rune = map[int]rune{
	1: '@',
}

var tileCharMap map[int]rune = map[int]rune{
	-1: '}', // Error
	0:  ' ',
	1:  '.',
	2:  '+',
	3:  '-',
	4:  '|',
}

func Render(g Game) {
	for j := 0; j < g.WorldMap.Height; j++ {
		for i := 0; i < g.WorldMap.Width; i++ {
			termbox.SetCell(i, j, tileCharMap[g.WorldMap.GetTileAt(i, j)], termbox.ColorWhite, termbox.ColorDefault)
		}
	}
	for _, entity := range g.Entities {
		termbox.SetCell(entity.X, entity.Y, entityCharMap[entity.Kind], termbox.ColorYellow, termbox.ColorDefault)
	}
}
