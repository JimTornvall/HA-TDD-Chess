package chess

// PieceBishop represents a Rook chess piece
type pieceBishop struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *pieceBishop) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *pieceBishop) canMove(b *Board, x, y int) bool {
	if !isValidBishopMove(b, p.x, p.y, x, y) {
		return false
	}

	if !checkEmpty(b, x, y) && checkIsSameColor(b, p.x, p.y, x, y) {
		return false
	}
	if isKingInCheckAfterMove(b, p.x, p.y, x, y) {
		return false
	}
	return true
}

// position returns the current position of the piece
func (p *pieceBishop) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *pieceBishop) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *pieceBishop) PType() PieceType {
	return p.pieceType
}

// NewBishopPiece creates a new Bishop piece
func NewBishopPiece(b *Board, x, y int, color Color) error {
	bishop := &pieceBishop{x, y, BISHOP, color}
	b.board[y][x] = bishop
	return nil
}
