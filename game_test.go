package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"testing"
)

// GameSuite is a test suite for the Game
type GameSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *GameSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestGameSuite runs the Game test suite, and all the tests within the suite
func TestGameSuite(t *testing.T) {
	suite.Run(t, new(GameSuite))
}

// HelperCaptureOutput captures the stdout output of a function
func HelperCaptureOutput(f func()) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	err := w.Close()
	if err != nil {
		return "", err
	}
	out, _ := io.ReadAll(r)
	return string(out), err
}

// Test_Render_The_Board tests that the board can be rendered
func (suite *GameSuite) Test_Render_The_Board() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	output := suite.board.Render()

	assert.Equal(suite.T(), expected, output)
}
