package board

import (
	"github.com/alexandreLITHAUD/go-chess-engine/engine/pieces"
)

/**
 * Board represents the chess board and contains the pieces on it.
 * It is a 2D array of Pieces, where each element represents a square on the board.
 * The first dimension represents the ranks (rows) and the second dimension represents the files (columns).
 */
type Board struct {
	Squares [][]*pieces.Piece
}

func NewBoard(height uint8, width uint8) *Board {
	// Initialize the board with empty pieces (or nil) for each square.
	squares := make([][]*pieces.Piece, height)
	for i := range squares {
		squares[i] = make([]*pieces.Piece, width)
	}
	return &Board{Squares: squares}
}
