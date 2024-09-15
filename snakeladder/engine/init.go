package engine

import "github.com/snakeladder/model"

func InitEngine(numberOfSnakes, numberOfLadders, boardSize int64) *Engine {
	return &Engine{
		board: model.InitBoard(boardSize, numberOfSnakes, numberOfLadders),
		dice: model.InitDice(6),
		players: []*model.Player{},
	}
}
