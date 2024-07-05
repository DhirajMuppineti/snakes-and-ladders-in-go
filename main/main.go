package main

import (
	"fmt"
	"main/board"
	"math/rand"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	b := board.GetBoard()
	turn := true
	fmt.Println("Welcome to Snakes and Ladders !")
	fmt.Println("Currently we're under constructions, so there are no snakes or ladders, but you can enjoy the game nevertheless.")
	for {
		if turn {
			fmt.Println("Player One's turn, roll the dice!")
		} else {
			fmt.Println("Player Two's turn, roll the dice!")
		}

		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeyEnter {
				break
			} else {
				fmt.Printf("Press Enter to roll the dice.")
			}
		}

		dice := rand.Intn(6) + 1
		fmt.Printf("You rolled %d! \n", dice)

		if turn {
			newPlayerOnePos := b.PlayerOnePos + dice
			if newPlayerOnePos > 100 {
				fmt.Println("Sorry, can't move this turn.")
				turn = !turn
				continue
			}
			b.UpdateBoard(newPlayerOnePos, b.PlayerTwoPos)
		} else {
			newPlayerTwoPos := b.PlayerTwoPos + dice
			if newPlayerTwoPos > 100 {
				fmt.Println("Sorry, can't move this turn.")
				turn = !turn
				continue
			}
			b.UpdateBoard(b.PlayerOnePos, newPlayerTwoPos)
		}

		b.PrintBoard()
		fmt.Println(b.PlayerOnePos, b.PlayerTwoPos)
		if b.PlayerOnePos == 100 {
			fmt.Println("Player 1 won the game!")
			break
		}
		if b.PlayerTwoPos == 100 {
			fmt.Println("Player 2 won the game!")
			break
		}
		turn = !turn
	}

}
