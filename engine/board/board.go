package board

import (
	"encoding/json"
	"fmt"

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

type PieceListData struct {
	Pieces []*pieces.Piece `json:"pieces"`
}

func NewBoard(height uint8, width uint8) *Board {
	// Initialize the board with empty pieces (or nil) for each square.
	squares := make([][]*pieces.Piece, height)
	for i := range squares {
		squares[i] = make([]*pieces.Piece, width)
	}
	return &Board{Squares: squares}
}

func (b *Board) LoadPiecesFromJSON(jsonData []byte) error {
	var data PieceListData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	height := len(b.Squares)

	for _, p := range data.Pieces {
		// Ensure coordinates are within board bounds
		if int(p.Position.Y) >= height || int(p.Position.X) >= len(b.Squares[0]) {
			return fmt.Errorf("piece position out of bounds: (%d, %d)", p.Position.X, p.Position.Y)
		}

		// Map JSON Y (0 = bottom) to Array Y (0 = top)
		// Formula: ArrayY = (Height - 1) - JsonY
		boardY := (height - 1) - int(p.Position.Y)
		boardX := int(p.Position.X)

		// Place the pointer on the board
		b.Squares[boardY][boardX] = p
	}

	return nil
}
