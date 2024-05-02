package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"testing"
)

type GameSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the Simple calculator and logger before each test
func (suite *GameSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestSimpleSuite runs the Simple test suite, and all the tests within the suite
func TestGameSuite(t *testing.T) {
	suite.Run(t, new(GameSuite))
}

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

func (suite *GameSuite) Test_Render() {
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

	//assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
}
