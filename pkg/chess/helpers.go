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
// souce: (Book) Tetra Chess in Basic
func isValidKnightMove(x, y, newX, newY int) bool {
	// Calculate the absolute differences
	xdiff := abs(newX - x)
	ydiff := abs(newY - y)

	// Check if it's a valid knight move
	return (xdiff == 1 && ydiff == 2) || (xdiff == 2 && ydiff == 1)
}
