package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// QueenSuite is a test suite for the Queen piece
type QueenSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *QueenSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestQueenSuite runs the Queen test suite, and all the tests within the suite
func TestQueenSuite(t *testing.T) {
	suite.Run(t, new(QueenSuite))
}

// Test_Render_A_Queen_And_Check_Its_Variables tests that a Queen can be rendered and its variables can be checked
func (suite *QueenSuite) Test_Render_A_Queen_And_Check_Its_Variables() {
	expected := `🔲🔳🔲❤️🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	expectedColor := chess.Color(chess.WHITE)
	expectedType := chess.PieceType(chess.QUEEN)

	err := chess.NewQueenPiece(&suite.board, 3, 0, chess.WHITE)
	output := suite.board.Render()
	color := suite.board.PieceAt(3, 0).Color()
	ptype := suite.board.PieceAt(3, 0).PType()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
	assert.Equal(suite.T(), expectedColor, color)
	assert.Equal(suite.T(), expectedType, ptype)
}

// Test_Move_A_Queen_To_Position_One tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_One() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
❤️🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 0, 1)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Two tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Two() {
	expected := `🔲🔳🔲🔳🔲🔳🔲❤️ 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 7, 0)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Three tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Three() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
❤️🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 0, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Four tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Four() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲❤️🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 6, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Five tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Five() {
	expected := `🔲🔳🔲❤️🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 3, 0)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Six tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Six() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
❤️🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 0, 4)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Seven tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Seven() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲❤️ 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 7, 4)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Position_Eight tests that a Queen can be moved to a position
func (suite *QueenSuite) Test_Move_A_Queen_To_Position_Eight() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳❤️🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 3, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_Through_Another_Piece tests that a Queen cant be moved through another piece
func (suite *QueenSuite) Test_Move_A_Queen_Through_Another_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲❤️🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲😍🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewQueenPiece(&suite.board, 4, 6, chess.BLACK)
	err3 := suite.board.Move(3, 4, 4, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Strike_Own_Piece tests that a Queen cant be moved to a square with the same color
func (suite *QueenSuite) Test_Move_A_Queen_To_Strike_Own_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲❤️🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🦄🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 4, 6, chess.WHITE)
	err3 := suite.board.Move(3, 4, 4, 6)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_To_Strike_Opponent_Piece tests that a Queen can be moved to a square with an opponent piece
func (suite *QueenSuite) Test_Move_A_Queen_To_Strike_Opponent_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲❤️🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 4, 6, chess.BLACK)
	err3 := suite.board.Move(3, 4, 4, 6)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Queen_Illegal_Move tests that a Queen cant be moved to an illegal position
func (suite *QueenSuite) Test_Move_A_Queen_Illegal_Move() {
	err1 := chess.NewQueenPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 7, 7)
	err3 := suite.board.Move(3, 4, 4, 1)
	err4 := suite.board.Move(3, 4, 7, 4)

	assert.Nil(suite.T(), err1)
	assert.NotNil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.NotNil(suite.T(), err4)
}
