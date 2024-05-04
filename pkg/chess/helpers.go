package chess

import (
	"fmt"
	"reflect"
)

// GetYCoordinate get chess y coordinate from array index
func GetYCoordinate(y int) string {
	return fmt.Sprint(8 - y)
}

// getXCoordinate get chess x coordinate from array index
func getXCoordinate(x int) string {
	return string(rune(x + 65))
}

// getSquareColor given an x,y coordinate return if a square should be black or white
func getSquareColor(x, y int) Color {
	if (x+y)%2 == 0 {
		return WHITE
	}
	return BLACK
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// checkEmpty check if the piece is moving to an empty space
func checkEmpty(b *Board, x int, y int) bool {
	if reflect.TypeOf(b.PieceAt(x, y)) == reflect.TypeOf(&PieceEmpty{}) {
		return true
	}
	return false
}

// checkIsSameColor check if the piece is moving to a square with the same color
func checkIsSameColor(b *Board, x1, y1, x2, y2 int) bool {
	if b.board[y1][x1].Color() == b.board[y2][x2].Color() {
		return true
	}
	return false
}

// isValidKnightMove checks if the move is a valid knight move
// source: (Book) Tetra Chess in Basic
func isValidKnightMove(x, y, newX, newY int) bool {
	// Calculate the absolute differences
	dx := abs(newX - x)
	dy := abs(newY - y)

	// Check if it's a valid knight move
	return (dx == 1 && dy == 2) || (dx == 2 && dy == 1)
}

// isValidBishopMove checks if the move is a valid bishop move
// source: (Book) Tetra Chess in Basic
func isValidBishopMove(b *Board, x1, y1, x2, y2 int) bool {
	if !checkDiagonal(b, x1, y1, x2, y2) {
		return false
	}

	// Calculate the absolute differences between coordinates
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)

	// If the absolute differences are equal, it's a diagonal move
	return dx == dy
}

// isValidRookMove checks if the move is a valid rook move
func isValidRookMove(b *Board, x1, y1, x2, y2 int) bool {
	// Check if the move is either horizontal or vertical
	if x1 != x2 && y1 != y2 {
		return false
	}

	// Determine the direction of movement
	dx := 0
	dy := 0
	if x1 != x2 {
		dx = (x2 - x1) / abs(x2-x1)
	} else {
		dy = (y2 - y1) / abs(y2-y1)
	}

	// Check for obstacles along the path
	for x, y := x1+dx, y1+dy; x != x2 || y != y2; x, y = x+dx, y+dy {
		if !checkEmpty(b, x, y) {
			return false // Obstacle found
		}
	}

	return true
}

// checkDiagonal checks if there are any pieces in the way of a diagonal move
func checkDiagonal(b *Board, x1, y1, x2, y2 int) bool {
	// Calculate the direction of the move
	dx := 1
	if x2 < x1 {
		dx = -1
	}
	dy := 1
	if y2 < y1 {
		dy = -1
	}

	// Check if there are any pieces in the way
	for i := 1; i < abs(x2-x1); i++ {
		if !checkEmpty(b, x1+i*dx, y1+i*dy) {
			return false
		}
	}
	return true
}

// isValidKingMove checks if the move is a valid king move
func isValidKingMove(b *Board, x, y, newX, newY int) bool {
	dx := abs(newX - x)
	dy := abs(newY - y)

	// Check if the move is within one step (horizontally, vertically, or diagonally)
	if dx > 1 || dy > 1 {
		return false
	}

	// Check if any other piece can move to the new position
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !checkEmpty(b, i, j) && (i != x || j != y) && !checkIsSameColor(b, x, y, i, j) {
				if b.board[j][i].canMove(b, newX, newY) {
					return false // Another piece can move to the target position
				}
			}
		}
	}
	return true
}

// FindKingPosition finds the position of the king
func FindKingPosition(b *Board, color Color) (int, int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if reflect.TypeOf(b.PieceAt(i, j)) == reflect.TypeOf(&pieceKing{}) && b.PieceAt(i, j).Color() == color {
				return i, j
			}
		}
	}
	return -1, -1
}

// IsKingInCheck checks if the king is in check
func IsKingInCheck(b *Board, color Color) bool {
	kingX, kingY := FindKingPosition(b, color)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !checkEmpty(b, i, j) && b.PieceAt(i, j).Color() != color {
				if b.PieceAt(i, j).canMove(b, kingX, kingY) {
					return true
				}
			}
		}
	}
	return false
}

// Copy Board and move piece to check if king is in check
func isKingInCheckAfterMove(b *Board, x1, y1, x2, y2 int) bool {
	// Copy the board
	newBoard, err := NewBoard(b.charType)
	if err != nil {
		// TODO: Handle error, send back could not create board
		return false
	}
	newBoard = *b

	// Move the piece
	newBoard.board[y2][x2] = newBoard.board[y1][x1]
	newBoard.board[y1][x1] = &PieceEmpty{x1, y1, EMPTY, getSquareColor(x1, y1)}

	// Check if the king is in check
	return IsKingInCheck(&newBoard, newBoard.Turn())
}
