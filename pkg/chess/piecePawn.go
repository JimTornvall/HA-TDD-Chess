package chess

import (
	"reflect"
)

type PiecePawn struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
	hasMoved   bool
}

// Move moves the piece to the given position
func (p *PiecePawn) move(b *Board, x, y int) error {
	return nil
}

// CanMove see if a piece can move to a position
func (p *PiecePawn) canMove(b *Board, x, y int) bool {

	// check if the piece is moving to the same position
	// as an other piece of the same other color
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

	// check if the piece is moving to an empty space
	if reflect.TypeOf(b.PieceAt(x, y)) != reflect.TypeOf(&PieceEmpty{}) {
		return false
	}

	return true
}

func (p *PiecePawn) canMove2(b *Board, x, y int) bool {
	dx, dy := x-p.x, y-p.y
	if p.pieceColor == WHITE {
		if dx == 1 && dy == 0 && b.PieceAt(x, y) == nil {
			return true
		}
		if dx == 1 && abs(dy) == 1 && b.PieceAt(x, y) != nil {
			return true
		}
	} else {
		if dx == -1 && dy == 0 && b.PieceAt(x, y) == nil {
			return true
		}
		if dx == -1 && abs(dy) == 1 && b.PieceAt(x, y) != nil {
			return true
		}
	}
	// Add your current functionality here
	// ...
	return false
}

// Position returns the current position of the piece
func (p *PiecePawn) position(b *Board) (x, y int) {
	return p.x, p.y
}

// Color returns the color of the piece
func (p *PiecePawn) Color() Color {
	return p.pieceColor
}

// Type returns the type of the piece
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
