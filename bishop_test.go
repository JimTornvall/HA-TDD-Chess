package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// BishopSuite is a test suite for the Bishop piece
type BishopSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *BishopSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestBishopSuite runs the Bishop test suite, and all the tests within the suite
func TestBishopSuite(t *testing.T) {
	suite.Run(t, new(BishopSuite))
}

// Test_Render_A_Bishop_And_Check_Its_Variables tests that a Bishop can be rendered and its variables can be checked
func (suite *BishopSuite) Test_Render_A_Bishop_And_Check_Its_Variables() {
	expected := `🔲🔳🗼🔳🔲🔳🔲🔳 8
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
	expectedType := chess.PieceType(chess.BISHOP)

	err := chess.NewBishopPiece(&suite.board, 2, 0, chess.WHITE)
	output := suite.board.Render()
	color := suite.board.PieceAt(2, 0).Color()
	ptype := suite.board.PieceAt(2, 0).PType()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
	assert.Equal(suite.T(), expectedColor, color)
	assert.Equal(suite.T(), expectedType, ptype)
}

// Test_Move_A_Bishop_To_Position_One tests that a Bishop can be moved to a position
func (suite *BishopSuite) Test_Move_A_Bishop_To_Position_One() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🗼🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 0, 1)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Bishop_To_Position_Two tests that a Bishop can be moved to a position
func (suite *BishopSuite) Test_Move_A_Bishop_To_Position_Two() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🗼 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 7, 0)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Bishop_To_Position_Three tests that a Bishop can be moved to a position
func (suite *BishopSuite) Test_Move_A_Bishop_To_Position_Three() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🗼🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 0, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Bishop_To_Position_Four tests that a Bishop can be moved to a position
func (suite *BishopSuite) Test_Move_A_Bishop_To_Position_Four() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🗼🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 6, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Bishop_Through_Another_Piece tests that a Bishop can't move through other pieces
func (suite *BishopSuite) Test_Move_A_Bishop_Through_Another_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🗼🔲🔳🔲🔳 4
🔳🔲🔳🔲🤓🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewPawnPiece(&suite.board, 4, 5, chess.BLACK)
	err3 := suite.board.Move(3, 4, 6, 7)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Bishop_Strike_Own_Piece tests that a Bishop can't move to a square with the same color
func (suite *BishopSuite) Test_Move_A_Bishop_Strike_Own_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🐮🔲🔳🔲 5
🔲🔳🔲🗼🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewPawnPiece(&suite.board, 4, 3, chess.WHITE)
	err3 := suite.board.Move(3, 4, 4, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Bishop_To_Strike_Opponent tests that a Bishop can be moved to strike an opponent
func (suite *BishopSuite) Test_Move_A_Bishop_To_Strike_Opponent() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🗼🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewBishopPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewPawnPiece(&suite.board, 4, 5, chess.BLACK)
	err3 := suite.board.Move(3, 4, 4, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Nil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}
