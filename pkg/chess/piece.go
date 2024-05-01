package chess

// Piece is an interface that represents a chess piece
type Piece interface {
	// Move moves the piece to the given position
	Move(x, y int) error
	// Position returns the current position of the piece
	Position() (x, y int)
	// Color returns the color of the piece
	Color() Color
	// Type returns the type of the piece
	Type() PieceType
}

// Color - Constants for the color of the piece
type Color int

const (
	WHITE = iota
	BLACK
)

// PieceType - Constants for the type of the piece
type PieceType int

const (
	EMPTY = iota
	KING
	QUEEN
	ROOK
	KNIGHT
	BISHOP
	PAWN
)

type StringPiece interface {
	// GetPiece returns the piece
	GetPiece(color Color, pieceType PieceType) string
}
