package pieces

import (
	"github.com/alexandreLITHAUD/go-chess-engine/engine/position"
)

/**
 * Pieces represents a chess piece and its properties.
 * It contains information about the type of the piece, its color, movement capabilities, and other attributes.
 */
type Piece struct {
	Position           position.Coordinates `json:"position"`           // Current position of the piece on the board
	Type               string               `json:"type"`               // Type of the piece (e.g., "Pawn", "Knight", "Bishop", "Rook", "Queen", "King", ...)
	CharRepresentation string               `json:"CharRepresentation"` // Character representation of the piece (e.g., 'P' for Pawn, 'N' for Knight, 'B' for Bishop, 'R' for Rook, 'Q' for Queen, 'K' for King)
	Color              string               `json:"color"`              // Color of the piece (e.g., "White", "Black", "Red", "Blue", ...)
	IsFirstMove        bool                 `json:"isFirstMove"`        // Indicates if the piece has moved before (useful for pawns and castling)
	IsSliding          bool                 `json:"isSliding"`          // Indicates if the piece can move multiple squares in a direction (e.g., bishops, rooks, queens)
	IsImmortal         bool                 `json:"isImmortal"`         // Indicates if the piece cannot be captured (e.g., kings in some variants)
	Range              int8                 `json:"range"`              // Maximum range of movement for the piece (e.g., 1 for knights, 8 for queens)
	CanCapture         bool                 `json:"canCapture"`         // Indicates if the piece can capture other pieces (e.g., pawns can capture diagonally, but not straight ahead)
	CanPromote         bool                 `json:"canPromote"`         // Indicates if the piece can be promoted (e.g., pawns can be promoted to queens, rooks, bishops, or knights)
	MoveDirs           []position.Offset    `json:"moveDirs"`           // Possible move directions for the piece (e.g., for a knight: [[2, 1], [2, -1], [-2, 1], [-2, -1], [1, 2], [1, -2], [-1, 2], [-1, -2]])
	CaptureDirs        []position.Offset    `json:"captureDirs"`        // Possible capture directions for the piece (e.g., for a pawn: [[1, 0], [1, -1], [1, 1]] for white pawns)
}

func newPiece(pieceType string, color string, isFirstMove bool, isSliding bool, isImmortal bool, moveRange int8, canCapture bool, canPromote bool, moveDirs []position.Offset, captureDirs []position.Offset, position position.Coordinates) *Piece {
	return &Piece{
		Position:    position,
		Type:        pieceType,
		Color:       color,
		IsFirstMove: isFirstMove,
		IsSliding:   isSliding,
		IsImmortal:  isImmortal,
		Range:       moveRange,
		CanCapture:  canCapture,
		CanPromote:  canPromote,
		MoveDirs:    moveDirs,
		CaptureDirs: captureDirs,
	}
}

func (p *Piece) Char() rune {
	return rune(p.CharRepresentation[0])
}
