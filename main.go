package main

import (
	"github.com/alexandreLITHAUD/go-chess-engine/engine/board"
	"github.com/alexandreLITHAUD/go-chess-engine/engine/pieces"
)

func main() {

	// Create an 8x8 board
	myBoard := board.NewBoard(12, 30)

	// Example: Manually placing a piece
	// Note: In your current struct, Pieces are values.
	myBoard.Squares[0][0] = &pieces.Piece{
		Type:               "Rook",
		CharRepresentation: "R",
		Color:              "Black",
	}

	myBoard.Squares[7][4] = &pieces.Piece{
		Type:               "King",
		CharRepresentation: "K",
		Color:              "White",
	}

	myBoard.Print()
}
