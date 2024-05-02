package chess

type CharLetter struct {
}

func (CharLetter) GetPiece(color Color, pieceType PieceType) string {
	var piece string
	switch pieceType {
	case EMPTY:
		piece = " O"
	case KING:
		piece = "K"
	case QUEEN:
		piece = "Q"
	case ROOK:
		piece = "R"
	case KNIGHT:
		piece = "N"
	case BISHOP:
		piece = "B"
	case PAWN:
		piece = "P"
	}
	if color == WHITE {
		piece = "\033[97m" + piece + "\033[0m"
	} else {
		piece = "\033[31m" + piece + "\033[0m"
	}
	return piece
}
