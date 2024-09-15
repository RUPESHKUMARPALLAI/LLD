package engine

import (
	"fmt"

	"github.com/snakeladder/model"
)

func (engine *Engine) AddPlayer(name string) {
	p := model.InitPlayer(name)
	engine.players = append(engine.players, p)
}

func (engine *Engine) Play() {
	if len(engine.players) == 0 {
		fmt.Println("no players joined")
		return
	}
	totalNoOfMoves := 0
	for {

		totalNoOfMoves++

		currentPlayer := engine.players[0]
		roll := engine.dice.Roll()

		newPosition := currentPlayer.Position + roll
		if newPosition > engine.board.GetEndValue() {
			fmt.Printf("Player: %s, diceroll: %d\n", currentPlayer, roll)
			engine.players = append(engine.players[1:], engine.players[0])
			continue
		}

		currentPlayer.Position = engine.board.GetNewPosition(currentPlayer.Position + roll)
		fmt.Printf("Player: %s, diceroll: %d\n", currentPlayer, roll)
		if currentPlayer.Position == engine.board.GetEndValue() {
			currentPlayer.Won = true
			fmt.Printf("%s won", currentPlayer)
			engine.players = engine.players[1:]
		} else {
			engine.players = append(engine.players[1:], engine.players[0])
		}

		if len(engine.players) < 2 {
			fmt.Printf("Game finished! Total number of moves: %d\n", totalNoOfMoves)
			return
		}

	}

}
