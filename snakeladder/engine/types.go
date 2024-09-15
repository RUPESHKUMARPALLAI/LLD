package engine

import "github.com/snakeladder/model"

type Engine struct {
	board *model.Board
	dice  model.DiceInterface
	players []*model.Player
}
