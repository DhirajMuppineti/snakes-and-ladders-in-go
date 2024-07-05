package board

import (
	"errors"
	"fmt"
)

type Board struct {
	PlayerOnePos int
	PlayerTwoPos int
}

func GetBoard() Board {
	b := Board{
		PlayerOnePos: 1,
		PlayerTwoPos: 1,
	}

	return b
}

func (b *Board) UpdateBoard(playerOnePos, playerTwoPos int) error {
	if playerOnePos <= 0 || playerTwoPos <= 0 || playerOnePos > 100 || playerTwoPos > 100 {
		return errors.New("out of bounds inputs")
	}
	b.PlayerOnePos = playerOnePos
	b.PlayerTwoPos = playerTwoPos

	return nil
}

func (b *Board) PrintBoard() {
	for i := 1; i <= 100; i++ {
		if i == b.PlayerOnePos && i == b.PlayerTwoPos {
			fmt.Print("*\t")
		} else if i == b.PlayerOnePos {
			fmt.Print("+\t")
		} else if i == b.PlayerTwoPos {
			fmt.Print("-\t")
		} else {
			fmt.Print(i, "\t")
		}

		if (i)%10 == 0 {
			fmt.Print("\n")
		}
	}
}
