package chess

import "fmt"

type Board struct {
	// 8x8 board
	board    [8][8]Piece
	charType StringPiece
}

func NewBoard(char StringPiece) Board {
	var b [8][8]Piece

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				b[i][j] = NewEmptyPiece(1, 2, WHITE)
			} else {
				b[i][j] = NewEmptyPiece(1, 2, BLACK)
			}
		}
	}
	return Board{board: b, charType: char}
}

func (b *Board) Render() string {
	var board string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board += b.charType.GetPiece(b.board[i][j].Color(), b.board[i][j].Type())
		}
		board += " " + b.GetYCordinate(i)
		board += "\n"
	}
	board += " A B C D E F G H\n"
	return board
}

// GetYCordinate get chess y cordinate from array index
func (b *Board) GetYCordinate(y int) string {
	return fmt.Sprint(8 - y)
}

// GetXCordinate get chess x cordinate from array index
func (b *Board) GetXCordinate(x int) string {
	return string(rune(x + 65))
}
