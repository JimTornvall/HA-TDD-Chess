package chess

// pieceKnight represents a Knight chess piece
type pieceKnight struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *pieceKnight) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *pieceKnight) canMove(b *Board, x, y int) bool {
	if !isValidKnightMove(p.x, p.y, x, y) {
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
func (p *pieceKnight) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *pieceKnight) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *pieceKnight) PType() PieceType {
	return p.pieceType
}

// NewKnightPiece creates a new Knight piece
func NewKnightPiece(b *Board, x, y int, color Color) error {
	knight := &pieceKnight{x, y, KNIGHT, color}
	b.board[y][x] = knight
	return nil
}
