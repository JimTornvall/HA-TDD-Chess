package chess

// pieceRook represents a Rook chess piece
type pieceRook struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *pieceRook) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *pieceRook) canMove(b *Board, x, y int) bool {
	return false
}

// position returns the current position of the piece
func (p *pieceRook) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *pieceRook) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *pieceRook) PType() PieceType {
	return p.pieceType
}

// NewRookPiece creates a new Rock piece
func NewRookPiece(b *Board, x, y int, color Color) Piece {
	return &pieceRook{x, y, ROOK, color}
}
