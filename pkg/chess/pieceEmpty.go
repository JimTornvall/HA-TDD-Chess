package chess

import (
	"errors"
)

// PieceEmpty represents a empty chess piece
// empty pieces are used to fill the board
type PieceEmpty struct {
	x          int
	y          int
	pieceType  PieceType
	pieceColor Color
}

// move moves the piece to the given position
func (p *PieceEmpty) move(b *Board, x, y int) error { return errors.New("cannot move empty piece") }

// canMove see if a piece can move to a position
func (p *PieceEmpty) canMove(b *Board, x, y int) bool { return false }

// position returns the current position of the piece
func (p *PieceEmpty) position(b *Board) (x, y int) {
	return p.x, p.y // return x, y from the piece
}

// Color returns the color of the piece
func (p *PieceEmpty) Color() Color {
	return p.pieceColor
}

// PType returns the type of the piece
func (p *PieceEmpty) PType() PieceType {
	return p.pieceType
}

// NewEmptyPiece creates a new empty piece
func NewEmptyPiece(b *Board, x, y int, color Color) error {
	piece := PieceEmpty{
		x:          x,
		y:          y,
		pieceColor: color,
		pieceType:  EMPTY,
	}
	b.board[x][y] = &piece
	return nil
}
