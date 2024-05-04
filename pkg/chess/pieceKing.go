package chess

// Loop over the board and check what piesces can move to the position the king is moving to
// If the king is in check, check if the piece can move to a position that will block the check

// pieceKing represents a King chess piece
type pieceKing struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *pieceKing) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *pieceKing) canMove(b *Board, x, y int) bool {

	if !isValidKingMove(b, p.x, p.y, x, y) {
		return false
	}

	if !checkEmpty(b, x, y) && checkIsSameColor(b, p.x, p.y, x, y) {
		return false

	}
	return true
}

// position returns the current position of the piece
func (p *pieceKing) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *pieceKing) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *pieceKing) PType() PieceType {
	return p.pieceType
}

// NewKingPiece creates a new King piece
func NewKingPiece(b *Board, x, y int, color Color) error {
	king := &pieceKing{
		x:          x,
		y:          y,
		pieceColor: color,
		pieceType:  KING,
	}
	b.board[y][x] = king
	return nil
}
