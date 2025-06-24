package routes

import "math/rand/v2"

type Message struct {
	Type string `json:"type"`
	Data map[string]any `json:"data"`
}

type Turn struct {
	Row, Col int
}

type GameBoard struct {
	Cells [10][8]int `json:"cells"`
	WaterLevel int `json:"waterLevel"`
}

func (b *GameBoard) initiate() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			b.Cells[i][j] = rand.IntN(3)
		} 
	}
	b.WaterLevel = 0
}

func (b *GameBoard) swapElements(e1 *Turn, e2 *Turn) {
	b.Cells[e1.Row][e2.Col], b.Cells[e2.Row][e2.Col] = b.Cells[e2.Row][e2.Col], b.Cells[e1.Row][e2.Col]
}