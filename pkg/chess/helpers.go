package chess

import "fmt"

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
