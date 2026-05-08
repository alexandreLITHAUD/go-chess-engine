package pieces

/**
 * Pieces represents a chess piece and its properties.
 * It contains information about the type of the piece, its color, movement capabilities, and other attributes.
 */
type Piece struct {
	Type        string   // Type of the piece (e.g., "Pawn", "Knight", "Bishop", "Rook", "Queen", "King", ...)
	Color       string   // Color of the piece (e.g., "White", "Black", "Red", "Blue", ...)
	IsFirstMove bool     // Indicates if the piece has moved before (useful for pawns and castling)
	IsSliding   bool     // Indicates if the piece can move multiple squares in a direction (e.g., bishops, rooks, queens)
	IsImmortal  bool     // Indicates if the piece cannot be captured (e.g., kings in some variants)
	Range       int8     // Maximum range of movement for the piece (e.g., 1 for knights, 8 for queens)
	CanCapture  bool     // Indicates if the piece can capture other pieces (e.g., pawns can capture diagonally, but not straight ahead)
	CanPromote  bool     // Indicates if the piece can be promoted (e.g., pawns can be promoted to queens, rooks, bishops, or knights)
	MoveDirs    [][]int8 // Possible move directions for the piece (e.g., for a knight: [[2, 1], [2, -1], [-2, 1], [-2, -1], [1, 2], [1, -2], [-1, 2], [-1, -2]])
	CaptureDirs [][]int8 // Possible capture directions for the piece (e.g., for a pawn: [[1, 0], [1, -1], [1, 1]] for white pawns)
}
