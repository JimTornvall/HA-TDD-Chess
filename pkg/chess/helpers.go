package chess

import "fmt"

// GetYCoordinate get chess y coordinate from array index
func GetYCoordinate(y int) string {
	return fmt.Sprint(8 - y)
}

// GetXCoordinate get chess x coordinate from array index
func GetXCoordinate(x int) string {
	return string(rune(x + 65))
}

// GetSquareColor given an x,y coordinate return if a square should be black or white
func GetSquareColor(x, y int) Color {
	if (x+y)%2 == 0 {
		return WHITE
	}
	return BLACK
}
