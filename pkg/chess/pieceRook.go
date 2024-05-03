package chess

// PieceRook represents a Rook chess piece
type PieceRook struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *PieceRook) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *PieceRook) canMove(b *Board, x, y int) bool {

	// only move horizontally or vertically
	if p.x != x && p.y != y {
		return false
	}

	// check if the piece move intersects with another piece
	dx, dy := x-p.x, y-p.y
	if dx != 0 {
		for i := 1; i < abs(dx); i++ {
			if dx > 0 {
				if !checkEmpty(b, p.x+i, p.y) {
					return false
				}
			} else {
				if !checkEmpty(b, p.x-i, p.y) {
					return false
				}
			}
		}
	} else {
		for i := 1; i < abs(dy); i++ {
			if dy > 0 {
				if !checkEmpty(b, p.x, p.y+i) {
					return false
				}
			} else {
				if !checkEmpty(b, p.x, p.y-i) {
					return false
				}
			}
		}
	}
	if !checkEmpty(b, x, y) && checkIsSameColor(b, p.x, p.y, x, y) {
		return false
	}
	return true
}

// position returns the current position of the piece
func (p *PieceRook) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *PieceRook) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *PieceRook) PType() PieceType {
	return p.pieceType
}

// NewRookPiece creates a new Rock piece
func NewRookPiece(b *Board, x, y int, color Color) error {
	rook := &PieceRook{
		x:          x,
		y:          y,
		pieceColor: color,
		pieceType:  ROOK,
	}
	b.board[y][x] = rook
	return nil
}
