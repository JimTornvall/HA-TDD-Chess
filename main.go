package main

import (
	"fmt"

	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
)

func main() {
	b, _ := chess.NewBoard(chess.CharEmoji{})
	println("Empty Board before init")
	fmt.Println(b.Render())
	println()
	println("Board after init")
	b.Init()
	fmt.Println(b.Render())
}
