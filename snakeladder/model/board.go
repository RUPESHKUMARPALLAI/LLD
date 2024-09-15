package model

import (
	"fmt"
	"math/rand"
)

type Board struct {
	size    int64
	snakes  []*snake
	ladders []*ladder
}

func InitBoard(size, noOfSnakes, noOfLadders int64) *Board {
	b := &Board{
		size: size,
	}

	snakeLadderMap := make(map[string]bool)

	for i := 0; i < int(noOfSnakes); i++ {
		for {
			start := rand.Int63n(b.size) + 1
			end := rand.Int63n(b.size) + 1
			if start == end || start == 1 {
				continue
			}
			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
				b.snakes = append(b.snakes, InitSnake(start, end))
				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}

		}
	}
	for i := 0; i < int(noOfLadders); i++ {
		for {
			start := rand.Int63n(b.size) + 1
			end := rand.Int63n(b.size) + 1
			if start == end || start == 1 {
				continue
			}
			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
				b.ladders = append(b.ladders, InitLadder(start, end))
				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}

		}
	}



	return b

}


func (b *Board) isLadder(pos int64) (bool, int64) {

	for _, val := range b.ladders {
		if val.startPoint == pos {
			return true, val.endPoint
		}
	}

	return false, -1
}


func (b *Board) isSnake(pos int64) (bool, int64) {
	for _, val := range b.snakes {
		if val.startPoint == pos {
			return true, val.endPoint
		}
	}
	return false, -1
}

func (b *Board) GetEndValue() int64 {
	return b.size
}

func (b *Board) GetNewPosition(pos int64) int64 {
	ok, val := b.isLadder(pos)
	if ok {
		fmt.Printf("is ladder. start: %d, end: %d\n", pos, val)
		return val
	}
	
	ok, val = b.isSnake(pos)
	if ok {
		fmt.Printf("is snake. start: %d, end: %d\n", pos, val)
		return val
	}

	return pos
}
