package main

import (
	"fmt"

	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
)

func main() {
	b, _ := chess.NewBoard(chess.CharEmoji{})
	println("Empty Board before init")
	fmt.Println(b.Render())
	println("Board after init")
	_ = b.Init()
	fmt.Println(b.Render())
}
