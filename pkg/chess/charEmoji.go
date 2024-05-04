package chess

type CharEmoji struct {
}

func (CharEmoji) GetPiece(color Color, pieceType PieceType) string {
	var piece string
	if color == WHITE {
		switch pieceType {
		case EMPTY:
			piece = "🔲"
		case KING:
			piece = "👑"
		case QUEEN:
			piece = "❤️"
		case ROOK:
			piece = "🏰"
		case KNIGHT:
			piece = "🦄"
		case BISHOP:
			piece = "🗼"
		case PAWN:
			piece = "🐮"
		}
	} else {
		switch pieceType {
		case EMPTY:
			piece = "🔳"
		case KING:
			piece = "🥴"
		case QUEEN:
			piece = "😍"
		case ROOK:
			piece = "🤩"
		case KNIGHT:
			piece = "🐴"
		case BISHOP:
			piece = "😎"
		case PAWN:
			piece = "🤓"
		}
	}
	return piece
}
