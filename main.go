package main

import (
	"fmt"
	"github.com/IsuruVihan/TicTacToe-Go/board"
	"github.com/IsuruVihan/TicTacToe-Go/game"
	"github.com/IsuruVihan/TicTacToe-Go/player"
)

func main() {
	for {
		if player.Create(1) {
			break
		}
	}

	for {
		if player.Create(2) {
			break
		}
	}

	for {
		game.SwitchPlayer()

		if len(game.AvailableCells(board.Board)) == 0 {
			fmt.Println("DRAW!")
			break
		}

		board.Insert()
		board.PrintBoard()
		if game.Check(board.Board) {
			break
		}
		fmt.Printf("====================")
	}

	game.Exit()
}
