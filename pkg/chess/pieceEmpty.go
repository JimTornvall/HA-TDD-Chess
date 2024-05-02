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

// Move moves the piece to the given position
func (p *PieceEmpty) move(b *Board, x, y int) error { return errors.New("cannot move empty piece") }

// CanMove see if a piece can move to a position
func (p *PieceEmpty) canMove(b *Board, x, y int) bool { return false }

// Position returns the current position of the piece
func (p *PieceEmpty) position(b *Board) (x, y int) {
	return p.x, p.y // return x, y from the piece
}

// Color returns the color of the piece
func (p *PieceEmpty) color() Color {
	return p.pieceColor
}

// Type returns the type of the piece
func (p *PieceEmpty) pType() PieceType {
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
