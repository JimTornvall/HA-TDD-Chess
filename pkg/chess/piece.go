package chess

// Piece is an interface that represents a chess piece
type Piece interface {
	// Move moves the piece to the given position
	move(board *Board, x, y int) error
	// CanMove see if a piece can move to a position
	canMove(board *Board, x, y int) bool
	// Position returns the current position of the piece
	position(board *Board) (x, y int)
	// Color returns the color of the piece
	color() Color
	// Type returns the type of the piece
	pType() PieceType
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
