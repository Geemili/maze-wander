package game

type TileMap struct {
	Tiles         []int
	Width, Height int
}

func (t TileMap) GetTileAt(x, y int) int {
	if x < 0 || y < 0 || x > t.Width || y > t.Height {
		return -1
	}
	return t.Tiles[(y*t.Width)+x]
}

func (t *TileMap) SetTileAt(x, y int, tile int) {
	if x < 0 || y < 0 || x > t.Width || y > t.Height {
		return
	}
	t.Tiles[(y*t.Width)+x] = tile
}
