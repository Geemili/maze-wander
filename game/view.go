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

func (g Game) Render() {
	for j := 0; j < g.WorldMap.Height; j++ {
		for i := 0; i < g.WorldMap.Width; i++ {
			termbox.SetCell(i, j, tileCharMap[g.WorldMap.GetTileAt(i, j)], termbox.ColorWhite, termbox.ColorDefault)
		}
	}
	for _, entity := range g.Entities {
		termbox.SetCell(entity.X, entity.Y, entityCharMap[entity.Kind], termbox.ColorYellow, termbox.ColorDefault)
	}
}

func (m MessageBox) Render() {
	if m.Visible {
		for i := m.X; i < m.X+m.W; i++ {
			termbox.SetCell(i, m.Y, '-', termbox.ColorWhite, termbox.ColorDefault)
			termbox.SetCell(i, m.Y+m.H, '-', termbox.ColorWhite, termbox.ColorDefault)
		}
		for i := m.Y; i < m.Y+m.H; i++ {
			termbox.SetCell(m.X, i, '|', termbox.ColorWhite, termbox.ColorDefault)
			termbox.SetCell(m.X+m.W, i, '|', termbox.ColorWhite, termbox.ColorDefault)
		}
		termbox.SetCell(m.X, m.Y, '+', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(m.X+m.W, m.Y, '+', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(m.X, m.Y+m.H, '+', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(m.X+m.W, m.Y+m.H, '+', termbox.ColorWhite, termbox.ColorDefault)

		x, y := 0, 0
		for _, char := range m.Message {
			termbox.SetCell(m.X+x+1, m.Y+y+1, char, termbox.ColorWhite, termbox.ColorDefault)
			x++
			if x > m.W-2 {
				y++
				x = 0
			}
		}
		termbox.SetCursor(m.X+x+1, m.Y+y+1)
	} else {
		termbox.HideCursor()
	}
}
