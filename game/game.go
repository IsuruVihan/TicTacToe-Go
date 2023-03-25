package game

import (
	"bufio"
	"fmt"
	"github.com/IsuruVihan/TicTacToe-Go/player"
	"os"
)

var (
	CurrentPlayer   string
	CurrentPlayerId int
	CurrentChar     = "X"
)

func SwitchPlayer() {
	if CurrentPlayer == player.One {
		CurrentPlayer, CurrentPlayerId, CurrentChar = player.Two, 2, "O"
	} else {
		CurrentPlayer, CurrentPlayerId, CurrentChar = player.One, 1, "X"
	}

	fmt.Printf("\nTurn: %v (%v)\n", CurrentPlayerId, CurrentPlayer)
}

func Check(board [9]string) bool {
	patterns := [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 0}}

	for _, pattern := range patterns {
		if board[pattern[0]] == CurrentChar &&
			board[pattern[0]] == board[pattern[1]] &&
			board[pattern[1]] == board[pattern[2]] {
			fmt.Printf("WINNER: Player %v (%v)", CurrentPlayerId, CurrentPlayer)
			return true
		}
	}

	return false
}

func AvailableCells(board [9]string) []int {
	var availableCells []int

	for i, cell := range board {
		if cell == "" {
			availableCells = append(availableCells, i)
		}
	}

	fmt.Println("Available cells: ", availableCells)

	return availableCells
}

func Exit() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPress \"e\" to terminate the program...")
	_, err := reader.ReadString('e')
	if err != nil {
		return
	}
}
