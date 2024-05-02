package chess

import (
	"errors"
)

// PieceEmpty represents a empty chess piece
// empty pieces are used to fill the board
type PieceEmpty struct {
	x     int
	y     int
	pType PieceType
	color Color
}

// Move moves the piece to the given position
func (p *PieceEmpty) Move(b *Board, x, y int) error { return errors.New("cannot move empty piece") }

// CanMove see if a piece can move to a position
func (p *PieceEmpty) CanMove(b *Board, x, y int) bool { return false }

// Position returns the current position of the piece
func (p *PieceEmpty) Position(b *Board) (x, y int) {
	return p.x, p.y // return x, y from the piece
}

// Color returns the color of the piece
func (p *PieceEmpty) Color() Color {
	return p.color
}

// Type returns the type of the piece
func (p *PieceEmpty) Type() PieceType {
	return p.pType
}

// NewEmptyPiece creates a new empty piece
func NewEmptyPiece(b *Board, x, y int, color Color) error {
	piece := PieceEmpty{
		x:     x,
		y:     y,
		color: color,
		pType: EMPTY,
	}
	b.board[x][y] = &piece
	return nil
}
