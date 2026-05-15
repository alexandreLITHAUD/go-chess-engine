package main

import (
	"os"

	"github.com/alexandreLITHAUD/go-chess-engine/engine/board"
)

func main() {

	// Create an 8x8 board
	myBoard := board.NewBoard(12, 30)

	// Example: Manually placing a piece
	// Note: In your current struct, Pieces are values.
	// myBoard.Squares[0][0] = &pieces.Piece{
	// 	Type:               "Rook",
	// 	CharRepresentation: "R",
	// 	Color:              "Black",
	// }
	//
	// myBoard.Squares[7][4] = &pieces.Piece{
	// 	Type:               "King",
	// 	CharRepresentation: "K",
	// 	Color:              "White",
	// }

	jsonBytes, err := os.ReadFile("./config/pieces.json")
	if err != nil {
		panic(err)
	}

	err = myBoard.LoadPiecesFromJSON(jsonBytes)
	if err != nil {
		panic(err)
	}

	myBoard.Print()
}
