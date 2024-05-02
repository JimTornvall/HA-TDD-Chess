package chess

type Board struct {
	// 8x8 board
	board    [8][8]Piece
	charType StringPiece
}

func NewBoard(char StringPiece) (Board, error) {
	//var b [8][8]Piece

	var b Board = Board{}
	b.charType = char

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
func (b *Board) Init() {
	//b.board[0][0] = NewRookPiece(0, 0, WHITE, ROOK)
	//b.board[0][1] = NewKnightPiece(0, 1, WHITE, KNIGHT)
	//b.board[0][2] = NewBishopPiece(0, 2, WHITE, BISHOP)
	//b.board[0][3] = NewQueenPiece(0, 3, WHITE, QUEEN)
	//b.board[0][4] = NewKingPiece(0, 4, WHITE, KING)
	//b.board[0][5] = NewBishopPiece(0, 5, WHITE, BISHOP)
	//b.board[0][6] = NewKnightPiece(0, 6, WHITE, KNIGHT)
	//b.board[0][7] = NewRookPiece(0, 7, WHITE, ROOK)
	//
	//for i := 0; i < 8; i++ {
	//	b.board[1][i] = NewPawnPiece(1, i, WHITE, PAWN)
	//}
	//
	//b.board[7][0] = NewRookPiece(7, 0, BLACK, ROOK)
	//b.board[7][1] = NewKnightPiece(7, 1, BLACK, KNIGHT)
	//b.board[7][2] = NewBishopPiece(7, 2, BLACK, BISHOP)
	//b.board[7][3] = NewQueenPiece(7, 3, BLACK, QUEEN)
	//b.board[7][4] = NewKingPiece(7, 4, BLACK, KING)
	//b.board[7][5] = NewBishopPiece(7, 5, BLACK, BISHOP)
	//b.board[7][6] = NewKnightPiece(7, 6, BLACK, KNIGHT)
	//b.board[7][7] = NewRookPiece(7, 7, BLACK, ROOK)
	//
	//for i := 0; i < 8; i++ {
	//	b.board[6][i] = NewPawnPiece(6, i, BLACK, PAWN)
	//}
}

func (b *Board) Render() string {
	var board string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board += b.charType.GetPiece(b.board[i][j].Color(), b.board[i][j].Type())
		}
		board += " " + GetYCoordinate(i)
		board += "\n"
	}
	board += " A B C D E F G H\n"
	return board
}
