package board

import (
	"fmt"
	"github.com/IsuruVihan/TicTacToe-Go/game"
)

var (
	Board [9]string
)

func Insert() {
	var cell int

	for {
		fmt.Print("Enter cell index: ")
		_, err := fmt.Scan(&cell)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if cell < 0 || cell > 8 {
			fmt.Println("Cell index must be 0 - 8!")
			continue
		}

		available := false
		for _, availableCell := range game.AvailableCells(Board) {
			if availableCell == cell {
				available = true
				break
			}
		}

		if !available {
			fmt.Printf("Cell %v is already filled!\n", cell)
			continue
		}

		Board[cell] = game.CurrentChar
		break
	}
}

func PrintBoard() {
	for i, char := range Board {
		tempChar := char
		if tempChar == "" {
			tempChar = "_"
		}

		if (i+1)%3 == 0 {
			fmt.Printf(" %v \n", tempChar)
		} else {
			fmt.Printf(" %v ", tempChar)
		}
	}
}
