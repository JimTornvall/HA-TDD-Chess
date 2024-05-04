package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// KingSuite is a test suite for the King piece
type KingSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *KingSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestKingSuite runs the King test suite, and all the tests within the suite
func TestKingSuite(t *testing.T) {
	suite.Run(t, new(KingSuite))
}

// Test_Render_A_King_And_Check_Its_Variables tests that a King can be rendered and its variables can be checked
func (suite *KingSuite) Test_Render_A_King_And_Check_Its_Variables() {
	expected := `🔲🔳🔲🔳👑🔳🔲🔳 8
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
	expectedType := chess.PieceType(chess.KING)

	err := chess.NewKingPiece(&suite.board, 4, 0, chess.WHITE)
	output := suite.board.Render()
	color := suite.board.PieceAt(4, 0).Color()
	ptype := suite.board.PieceAt(4, 0).PType()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
	assert.Equal(suite.T(), expectedColor, color)
	assert.Equal(suite.T(), expectedType, ptype)
}

// Test_Move_A_King_To_Position_One tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_One() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳👑🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 3, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Two tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Two() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲👑🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 4, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Three tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Three() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳👑🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 4, 4)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Four tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Four() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲👑🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 4, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Five tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Five() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳👑🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 3, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Six tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Six() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲👑🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 2, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Seven tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Seven() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳👑🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 2, 4)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Position_Eight tests that a King can be moved to a position
func (suite *KingSuite) Test_Move_A_King_To_Position_Eight() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲👑🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 2, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Strike_Own_Piece tests that a King cant be moved to strike its own piece
func (suite *KingSuite) Test_Move_A_King_To_Strike_Own_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲👑🔲🔳🔲🔳 4
🔳🔲🔳🔲🦄🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 4, 5, chess.WHITE)
	err3 := suite.board.Move(3, 4, 4, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Strike_Opponent_Piece tests that a King can be moved to strike an opponent piece
func (suite *KingSuite) Test_Move_A_King_To_Strike_Opponent_Piece() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲👑🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewKnightPiece(&suite.board, 4, 5, chess.BLACK)
	err3 := suite.board.Move(3, 4, 4, 5)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Nil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_King_To_Invalid_Position tests that a King cant be moved to an invalid position
func (suite *KingSuite) Test_Move_A_King_To_Invalid_Position() {
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := suite.board.Move(3, 4, 7, 7)
	err3 := suite.board.Move(3, 4, 0, 0)
	err4 := suite.board.Move(3, 4, 14, 4)
	err5 := suite.board.Move(3, 4, -1, -1)
	err6 := suite.board.Move(3, 4, 3, 6)
	err7 := suite.board.Move(3, 4, 5, 4)
	err8 := suite.board.Move(3, 4, 8, -1)

	assert.Nil(suite.T(), err1)
	assert.NotNil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.NotNil(suite.T(), err4)
	assert.NotNil(suite.T(), err5)
	assert.NotNil(suite.T(), err6)
	assert.NotNil(suite.T(), err7)
	assert.NotNil(suite.T(), err8)
}

// Test_Move_A_King_Into_Check tests that a King cant be moved into check
func (suite *KingSuite) Test_Move_A_King_Into_Check() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲👑🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🤩🔲🤩🔲🔳😎 1
 A B C D E F G H
`
	err1 := chess.NewKingPiece(&suite.board, 3, 4, chess.WHITE)
	err2 := chess.NewRookPiece(&suite.board, 2, 7, chess.BLACK)
	err3 := chess.NewRookPiece(&suite.board, 4, 7, chess.BLACK)
	err4 := chess.NewBishopPiece(&suite.board, 7, 7, chess.BLACK)
	err5 := suite.board.Move(3, 4, 2, 4)
	err6 := suite.board.Move(3, 4, 4, 5)
	err7 := suite.board.Move(3, 4, 3, 3)
	output := suite.board.Render()

	assert.Nil(suite.T(), err1, "Failed to create White King")
	assert.Nil(suite.T(), err2, "Failed to create Black Rook1")
	assert.Nil(suite.T(), err3, "Failed to create Black Rook2")
	assert.Nil(suite.T(), err4, "Failed to create Black Bishop")
	assert.NotNil(suite.T(), err5, "King tried to move into check by Rook1")
	assert.NotNil(suite.T(), err6, "King tried to move into check by Rook2")
	assert.NotNil(suite.T(), err7, "King tried to move into check by Bishop")
	assert.Equal(suite.T(), expected, output, "Board render should be the same as expected")
}
