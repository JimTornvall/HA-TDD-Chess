package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// KnightSuite is a test suite for the Knight piece
type KnightSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *KnightSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestKnightSuite runs the Knight test suite, and all the tests within the suite
func TestKnightSuite(t *testing.T) {
	suite.Run(t, new(KnightSuite))
}

// Test_Render_A_Knight_And_Check_Its_Variables tests that a Knight can be rendered and its variables can be checked
func (suite *KnightSuite) Test_Render_A_Knight_And_Check_Its_Variables() {
	expected := `🔲🦄🔲🔳🔲🔳🔲🔳 8
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
	expectedType := chess.PieceType(chess.KNIGHT)

	err := chess.NewKnightPiece(&suite.board, 1, 0, chess.WHITE)
	output := suite.board.Render()
	color := suite.board.PieceAt(1, 0).Color()
	ptype := suite.board.PieceAt(1, 0).PType()

	suite.Nil(err)
	suite.Equal(expected, output)
	suite.Equal(expectedColor, color)
	suite.Equal(expectedType, ptype)
}

// Knight Moves
//🔲🔳🔲🔳🔲🔳🔲🔳 8
//🔳🔲🔳🔲🔳🔲🔳🔲 7
//🔲🔳⭕🔳⭕🔳🔲🔳 6
//🔳⭕🔳🔲🔳⭕🔳🔲 5
//🔲🔳🔲🦄🔲🔳🔲🔳 4
//🔳⭕🔳🔲🔳⭕🔳🔲 3
//🔲🔳⭕🔳⭕🔳🔲🔳 2
//🔳🔲🔳🔲🔳🔲🔳🔲 1
// A B C D E F G H

// Test_Move_A_Knight_To_Position_One tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_One() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🦄🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 2, 2)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Two tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Two() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🦄🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 4, 2)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Three tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Three() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🦄🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 1, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Four tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Four() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🦄🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 5, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Five tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Five() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🦄🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 1, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Six tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Six() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🦄🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 5, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Seven tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Seven() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🦄🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 2, 6)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Position_Eight tests that a Knight can be moved to a position
func (suite *KnightSuite) Test_Move_A_Knight_To_Position_Eight() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🦄🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 4, 6)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_Over_Another_Piece tests that a Knight can be moved over another piece
func (suite *KnightSuite) Test_Move_A_Knight_Over_Another_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🦄🔳🔲🔳🔲🔳 4
🔳🦄🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 2, 4, chess.WHITE)
	err3 := suite.board.Move(3, 4, 1, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Nil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Invalid_Position tests that a Knight can't be moved to an invalid position
func (suite *KnightSuite) Test_Move_A_Knight_To_Invalid_Position() {
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)

	// Invalid position's to test
	err2 := suite.board.Move(3, 4, 0, 0)
	err3 := suite.board.Move(3, 4, 1, 1)
	err4 := suite.board.Move(3, 4, 7, 0)
	err5 := suite.board.Move(3, 4, 0, 7)
	err6 := suite.board.Move(3, 4, 7, 7)
	err7 := suite.board.Move(3, 4, 3, 4)

	assert.Nil(suite.T(), err1)
	assert.NotNil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.NotNil(suite.T(), err4)
	assert.NotNil(suite.T(), err5)
	assert.NotNil(suite.T(), err6)
	assert.NotNil(suite.T(), err7)
}

// Test_Move_A_Knight_To_Strike_Opponent tests that a Knight can be moved to strike an opponent
func (suite *KnightSuite) Test_Move_A_Knight_To_Strike_Opponent() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🦄🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 2, 2, chess.BLACK)
	err3 := suite.board.Move(3, 4, 2, 2)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Nil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Knight_To_Strike_Own_Piece tests that a Knight can't be moved to strike a piece of its own color
func (suite *KnightSuite) Test_Move_A_Knight_To_Strike_Own_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🦄🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🦄🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKnightPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 2, 2, chess.WHITE)
	err3 := suite.board.Move(3, 4, 2, 2)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}
