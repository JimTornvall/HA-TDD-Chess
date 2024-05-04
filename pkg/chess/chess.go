package chess

import (
	"errors"
	"fmt"
)

// Chess game
// expectations:
// - 8x8 board
// - 2 players
// - 6 different pieces
// - 2 colors
// White always starts, and is always at the top of the board

type Board struct {
	// 8x8 board
	board    [8][8]Piece
	charType StringPiece
	turn     Color
}

func NewBoard(char StringPiece) (Board, error) {
	//var b [8][8]Piece

	var b Board = Board{}
	b.charType = char
	b.turn = WHITE

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				err := NewEmptyPiece(&b, i, j, WHITE)
				if err != nil {
					return b, err
				}
			} else {
				err := NewEmptyPiece(&b, i, j, BLACK)
				if err != nil {
					return b, err
				}
			}
		}
	}
	return b, nil
}

// Init initializes the board, by placing the pieces in their starting positions
func (b *Board) Init() error {
	// ----- WHITE -----
	err := NewRookPiece(b, 0, 0, WHITE)
	if err != nil {
		return err
	}
	err = NewRookPiece(b, 7, 0, WHITE)
	if err != nil {
		return err
	}
	err = NewKnightPiece(b, 1, 0, WHITE)
	if err != nil {
		return err
	}
	err = NewKnightPiece(b, 6, 0, WHITE)
	if err != nil {
		return err
	}
	err = NewBishopPiece(b, 2, 0, WHITE)
	if err != nil {
		return err
	}
	err = NewBishopPiece(b, 5, 0, WHITE)
	if err != nil {
		return err
	}
	err = NewQueenPiece(b, 3, 0, WHITE)
	if err != nil {
		return err
	}
	//b.board[0][4] = NewKingPiece(0, 4, WHITE, KING)
	//
	for i := 0; i < 8; i++ {
		err := NewPawnPiece(b, i, 1, WHITE)
		if err != nil {
			return err
		}
	}
	// ----- BLACK -----
	err = NewRookPiece(b, 0, 7, BLACK)
	if err != nil {
		return err
	}
	err = NewRookPiece(b, 7, 7, BLACK)
	if err != nil {
		return err
	}
	err = NewKnightPiece(b, 1, 7, BLACK)
	if err != nil {
		return err
	}
	err = NewKnightPiece(b, 6, 7, BLACK)
	if err != nil {
		return err
	}
	err = NewBishopPiece(b, 2, 7, BLACK)
	if err != nil {
		return err
	}
	err = NewBishopPiece(b, 5, 7, BLACK)
	if err != nil {
		return err
	}
	err = NewQueenPiece(b, 3, 7, BLACK)
	if err != nil {
		return err
	}
	//b.board[7][4] = NewKingPiece(7, 4, BLACK, KING)
	//
	for i := 0; i < 8; i++ {
		err := NewPawnPiece(b, i, 6, BLACK)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Board) Move(x1, y1, x2, y2 int) error {
	if x1 < 0 || x1 > 7 || y1 < 0 || y1 > 7 {
		return errors.New("invalid piece")
	}
	if x2 < 0 || x2 > 7 || y2 < 0 || y2 > 7 {
		return errors.New("invalid position")
	}

	// DEBUG
	//println("Moving piece from", x1, y1, "to", x2, y2)
	//println("Piece at", x1, y1, "is", b.board[y1][x1].PType())
	//println("Piece at", x2, y2, "is", b.board[y2][x2].PType())
	//println("Turn is", b.turn)
	//println("Piece color is", b.board[y1][x1].Color())
	//println("Piece can move", b.board[y1][x1].canMove(b, x2, y2))
	// DEBUG

	if b.board[y1][x1].canMove(b, x2, y2) && b.board[y1][x1].Color() == b.turn {
		b.board[y2][x2] = b.board[y1][x1]
		b.SwitchTurn()
		b.board[y1][x1] = &PieceEmpty{x1, y1, EMPTY, getSquareColor(x1, y1)}
		return nil
	}
	return fmt.Errorf("cannot move piece from %d,%d to %d,%d", x1, y1, x2, y2)
}

func (b *Board) PieceAt(x, y int) Piece {
	return b.board[y][x]
}

func (b *Board) SwitchTurn() {
	if b.turn == WHITE {
		b.turn = BLACK
	} else {
		b.turn = WHITE
	}
}

func (b *Board) Turn() Color {
	return b.turn
}

func (b *Board) Render() string {
	var board string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board += b.charType.GetPiece(b.board[i][j].Color(), b.board[i][j].PType())
		}
		board += " " + GetYCoordinate(i)
		board += "\n"
	}
	board += " A B C D E F G H\n"
	return board
}
