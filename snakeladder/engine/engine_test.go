package engine

import (
	"testing"

	"github.com/snakeladder/model"
	"github.com/stretchr/testify/assert"
)

type MockDice struct {
	minValue     int64
	maxvalue     int64
	currentValue int64
}


func InitmockDice(currentValue int64) *MockDice {
	return &MockDice{
		minValue:     0,
		maxvalue:     100,
		currentValue: currentValue,
	}
}

func (d *MockDice) Roll() int64 {
	return d.currentValue
}

func TestAddPlayer(t *testing.T) {
	e := &Engine{
		board:   model.InitBoard(100, 5, 5),
		dice:    &model.Dice{},
		players: []*model.Player{},
	}
	e.AddPlayer("Hari")

	assert.Equal(t, 1, len(e.players), "Player list should contain one player")
	assert.Equal(t, "Hari", e.players[0].Name, "Player name should be Hari")
}

func TestPlayerAfterPlay(t *testing.T) {
	e := &Engine{
		board:   model.InitBoard(100, 5, 5),
		dice:    model.InitDice(6),
		players: []*model.Player{},
	}

	e.Play()

	assert.Equal(t, 0, len(e.players), "Player list should be empty")

	e.AddPlayer("Ram")
	e.AddPlayer("Shyam")

	e.Play()

	assert.Equal(t, 1, len(e.players), "Player list should contain 1 player")	
}


func TestPlayWithLoosing(t *testing.T) {
	e := &Engine{
		board:   model.InitBoard(100, 0, 0),
		dice:    InitmockDice(10),
		players: []*model.Player{},
	}

	e.AddPlayer("Ram")
	e.AddPlayer("Shyam")

	e.Play()

	assert.Equal(t, e.players[0].Won, false, "Last Player must loose the game")	
}


