package rules

import (
	"github.com/alexandreLITHAUD/go-chess-engine/engine/board"
	"github.com/alexandreLITHAUD/go-chess-engine/engine/pieces"
	"github.com/alexandreLITHAUD/go-chess-engine/engine/position"
)

type Rule interface {
	Name() string
	GetMoves(board *board.Board, piece *pieces.Piece, coordinates position.Coordinates) []position.Move
}
