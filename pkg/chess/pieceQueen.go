package chess

// PieceQueen represents a Rook chess piece
type pieceQueen struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *pieceQueen) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *pieceQueen) canMove(b *Board, x, y int) bool {
	if !isValidRookMove(b, p.x, p.y, x, y) && !isValidBishopMove(b, p.x, p.y, x, y) {
		return false
	}

	if !checkEmpty(b, x, y) && checkIsSameColor(b, p.x, p.y, x, y) {
		return false
	}
	return true
}

// position returns the current position of the piece
func (p *pieceQueen) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *pieceQueen) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *pieceQueen) PType() PieceType {
	return p.pieceType
}

// NewQueenPiece creates a new Queen piece
func NewQueenPiece(b *Board, x, y int, color Color) error {
	queen := &pieceQueen{x, y, QUEEN, color}
	b.board[y][x] = queen
	return nil
}
