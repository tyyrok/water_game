package routes

import "math/rand/v2"

type Message struct {
	Type string `json:"type"`
	Data map[string]any `json:"data"`
}

type Turn struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func (t *Turn) isEqual(other *Turn) bool {
	if (t.Row == other.Row && t.Col == other.Col) {
		return true
	}
	return false
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
	b.Cells[e1.Row][e1.Col], b.Cells[e2.Row][e2.Col] = b.Cells[e2.Row][e2.Col], b.Cells[e1.Row][e1.Col]
}

func (b *GameBoard) addWaterElement(turns []Turn) []int {
	b.Cells[turns[2].Row][turns[2].Col] = 3
	b.Cells[turns[0].Row][turns[0].Col] = rand.IntN(3)
	b.Cells[turns[1].Row][turns[1].Col] = rand.IntN(3)
	return []int{
		b.Cells[turns[0].Row][turns[0].Col],
		b.Cells[turns[1].Row][turns[1].Col]}
}