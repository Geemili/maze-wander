package game

type Game struct {
	WorldMap TileMap
	Entities []*Entity
}

func NewGame(width, height int) *Game {
	g := Game{
		TileMap{
			make([]int, width*height),
			width, height,
		},
		[]*Entity{},
	}
	return &g
}

func (g *Game) AddEntity(entity *Entity) {
	g.Entities = append(g.Entities, entity)
}
