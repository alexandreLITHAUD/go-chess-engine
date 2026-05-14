package game

import (
	"github.com/alexandreLITHAUD/go-chess-engine/engine/board"
	"github.com/alexandreLITHAUD/go-chess-engine/engine/rules"
)

type Game struct {
	// Board represents the chess board and contains the pieces on it.
	Board *board.Board

	// Rules represents the rules of the game, which determine how pieces can move and capture.
	Rules []rules.Rule

	// Players represents the players in the game, which can be human or AI.
	// Players []Player

	// CurrentTurn indicates which player's turn it is to move.
	CurrentTurn int
}
