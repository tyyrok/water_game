package routes

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleMessage(msg *Message, board *GameBoard) *Message{
	switch msg.Type {
	case "move":
		return handleTurn(msg, board)
	case "end_game":
		return &Message{Type: "end_game"}
	default:
		return &Message{
			Type: "error",
		}
	}
}

func handleTurn(msg *Message, board *GameBoard) *Message {
	turns, err := validateTurn(msg)
	if err != nil {
		return &Message{Type: "turn validation error",}
	}
	if len(turns) == 2 {
		board.swapElements(&turns[0], &turns[1])
		return &Message{Type: "move",
			Data: map[string]any{
				"status": "swap",
				"turns": turns,
				"waterLevel": board.WaterLevel}}
	}

	return &Message{}
}


func saveGameResult(board *GameBoard, ctx *gin.Context) {

}

func validateTurn(msg *Message) ([]Turn, error) {
	fail := false
	var turns []Turn
	rawTurns, ok := msg.Data["turns"]
	if !ok {
		fail = true
	}
	rTurns, ok := rawTurns.([]any)
	if !ok {
		fail = true
	}
	for _, elem := range rTurns {
		turn, ok := elem.(map[string]any)
		if !ok {
			fail = true
			break
		}
		row, ok := turn["row"]
		if !ok {
			fail = true
		}
		r, err := toInt(row)
		if err != nil {
			fail = true
		}
		col, ok := turn["col"]
		if !ok {
			fail = true
		}
		c, err := toInt(col)
		if err != nil {
			fail = true
		}
		if fail {
			break
		}

		turns = append(turns, Turn{Row: r-1, Col: c-1})
	}
	if fail{
		return []Turn{}, errors.New("turn validation error")
	}
	return turns, nil
}

func toInt(value any) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		r, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return r, nil
	default:
		return 0, errors.New("undefined type")
	}
}