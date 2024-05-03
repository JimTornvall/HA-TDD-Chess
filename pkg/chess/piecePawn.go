package chess

// PiecePawn represents a Pawn chess piece
type PiecePawn struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
	hasMoved   bool
}

// move moves the piece to the given position
func (p *PiecePawn) move(b *Board, x, y int) error {
	return nil
}

// canMove see if a piece can move to a position
func (p *PiecePawn) canMove(b *Board, x, y int) bool {

	// check if the piece is moving to the same position
	// as another piece of the same other color
	dx, dy := x-p.x, y-p.y
	if p.pieceColor == WHITE {
		if dx == 1 && dy == 0 && b.PieceAt(x, y) == nil {
			if b.PieceAt(x, y).Color() == BLACK {
				return true
			}
			return false
		}
		if dx == 1 && abs(dy) == 1 && b.PieceAt(x, y) != nil {
			if b.PieceAt(x, y).Color() == BLACK {
				return true
			}
			return false
		}
	} else {
		if dx == -1 && dy == 0 && b.PieceAt(x, y) == nil {
			if b.PieceAt(x, y).Color() == WHITE {
				return true
			}
			return false
		}
		if dx == -1 && abs(dy) == 1 && b.PieceAt(x, y) != nil {
			if b.PieceAt(x, y).Color() == WHITE {
				return true
			}
			return false
		}
	}

	// check if pawn can move two steps (hasMoved==false)
	if !p.hasMoved {
		// check if the piece is moving forward
		if p.pieceColor == WHITE {
			if p.x != x || (p.y+1 != y && p.y+2 != y) {
				return false
			}
		} else {
			if p.x != x || (p.y-1 != y && p.y-2 != y) {
				return false
			}
		}
	} else {
		// check if the piece is moving forward
		if p.pieceColor == WHITE {
			if p.x != x || p.y+1 != y {
				return false
			}
		} else {
			if p.x != x || p.y-1 != y {
				return false
			}
		}
	}

	if !checkEmpty(b, x, y) {
		return false
	}

	return true
}

// position returns the current position of the piece
func (p *PiecePawn) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *PiecePawn) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *PiecePawn) PType() PieceType {
	return p.pieceType
}

// NewPawnPiece creates a new pawn piece
func NewPawnPiece(b *Board, x, y int, color Color) error {
	pawn := &PiecePawn{
		x:          x,
		y:          y,
		pieceColor: color,
		pieceType:  PAWN,
	}
	b.board[y][x] = pawn
	return nil
}
