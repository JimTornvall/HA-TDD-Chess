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
