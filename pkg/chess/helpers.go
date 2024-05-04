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
func isValidBishopMove(x1, y1, x2, y2 int) bool {
	// Calculate the absolute differences between coordinates
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)

	// If the absolute differences are equal, it's a diagonal move
	return dx == dy
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
