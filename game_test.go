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

	assert.Equal(suite.T(), expected, output, "The board should be rendered empty")
}

// Test_Render_The_Board_With_Pieces tests that the board can be rendered with pieces
func (suite *GameSuite) Test_Render_The_Board_With_Pieces() {
	expected := `🏰🦄🗼❤️👑🗼🦄🏰 8
🐮🐮🐮🐮🐮🐮🐮🐮 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🤓🤓🤓🤓🤓🤓🤓🤓 2
🤩🐴😎😍🥴😎🐴🤩 1
 A B C D E F G H
`
	err := suite.board.Init()
	output := suite.board.Render()

	assert.Nil(suite.T(), err, "The board should be initialized")
	assert.Equal(suite.T(), expected, output, "The board should be rendered with pieces")
}

// Test_Own_King_Is_In_Check_After_Move tests that the king is in check after a move
func (suite *GameSuite) Test_Own_King_Is_In_Check_After_Move() {
	expected := `🔲🔳🔲🔳👑🔳🔲🔳 8
🔳🔲🔳🔲🏰🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🤩🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 4, 0, chess.WHITE)
	err2 := chess.NewRookPiece(&suite.board, 4, 1, chess.WHITE)
	err3 := chess.NewRookPiece(&suite.board, 4, 7, chess.BLACK)
	err4 := suite.board.Move(4, 1, 5, 1)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1, "The white king should be placed")
	assert.Nil(suite.T(), err2, "The white rook should be placed")
	assert.Nil(suite.T(), err3, "The black rook should be placed")
	assert.NotNil(suite.T(), err4, "The white rook should not be able to move")
	assert.Equal(suite.T(), expected, output, "The board should be rendered with white rook not moved")

}
