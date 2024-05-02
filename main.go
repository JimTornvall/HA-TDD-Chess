package main

import (
	"fmt"

	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
)

func main() {
	b, _ := chess.NewBoard(chess.CharEmoji{})
	fmt.Println(b.Render())
}
