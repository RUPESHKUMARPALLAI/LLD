package model

import (
	"fmt"
)

type Player struct {
	Name     string
	Position int64
	Won      bool
}

func InitPlayer(name string) *Player {
	return &Player{
		Name:     name,
		Position: 0,
		Won:      false,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("name: %s, position: %d", p.Name, p.Position)
}