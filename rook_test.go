package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// RookSuite is a test suite for the Rook piece
type RookSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *RookSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestRockSuite runs the Rook test suite, and all the tests within the suite
func TestRookSuite(t *testing.T) {
	suite.Run(t, new(RookSuite))
}

// Test_Render_A_Rook_And_Check_Its_Variables tests that a Rook can be rendered and its variables can be checked
func (suite *RookSuite) Test_Render_A_Rook_And_Check_Its_Variables() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🏰 8
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
	expectedType := chess.PieceType(chess.ROOK)

	err := chess.NewRookPiece(&suite.board, 7, 0, chess.WHITE)
	output := suite.board.Render()
	color := suite.board.PieceAt(7, 0).Color()
	ptype := suite.board.PieceAt(7, 0).PType()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
	assert.Equal(suite.T(), expectedColor, color)
	assert.Equal(suite.T(), expectedType, ptype)
}

// Test_Move_A_Rook_To_Next_Corner tests that a Rook can be moved one step
func (suite *RookSuite) Test_Move_A_Rook_To_Next_Corner() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🏰 1
 A B C D E F G H
`
	err1 := chess.NewRookPiece(&suite.board, 7, 0, chess.WHITE)
	err2 := suite.board.Move(7, 0, 7, 7)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Rook_Diagonal tests that a Rook can not be moved diagonally
func (suite *RookSuite) Test_Move_A_Rook_Diagonal() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🏰 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewRookPiece(&suite.board, 7, 0, chess.WHITE)
	err2 := suite.board.Move(7, 0, 6, 1)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.NotNil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Rook_Through_A_Piece tests that a Rook can not be moved through another piece
func (suite *RookSuite) Test_Move_A_Rook_Through_A_Piece_Of_Its_Own() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🏰 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🏰 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewRookPiece(&suite.board, 7, 0, chess.WHITE)
	err2 := chess.NewRookPiece(&suite.board, 7, 4, chess.WHITE)
	err3 := suite.board.Move(7, 0, 7, 7)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Rook_Through_A_Piece_Of_The_Opposite_Color tests that a Rook can not be moved through a piece of the opposite color
func (suite *RookSuite) Test_Move_A_Rook_Through_A_Piece_Of_The_Opposite_Color() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🏰 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🤩 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewRookPiece(&suite.board, 7, 0, chess.WHITE)
	err2 := chess.NewRookPiece(&suite.board, 7, 4, chess.BLACK)
	err3 := suite.board.Move(7, 0, 7, 7)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Rook_To_Strike_Opponent tests that a Rook can be moved to strike an opponent
func (suite *RookSuite) Test_Move_A_Rook_To_Strike_Opponent() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🏰 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewRookPiece(&suite.board, 7, 0, chess.WHITE)
	err2 := chess.NewRookPiece(&suite.board, 7, 4, chess.BLACK)
	err3 := suite.board.Move(7, 0, 7, 4)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Nil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}
