package main

import (
	"github.com/JimTornvall/HA-TDD-Chess/pkg/chess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// PawnSuite is a test suite for the Pawn piece
type PawnSuite struct {
	suite.Suite
	board chess.Board
}

// SetupTest initializes the board before each test
func (suite *PawnSuite) SetupTest() {
	suite.board, _ = chess.NewBoard(chess.CharEmoji{})
}

// TestPawnSuite runs the Pawn test suite, and all the tests within the suite
func TestPawnSuite(t *testing.T) {
	suite.Run(t, new(PawnSuite))
}

// Test_Render_A_Pawn_And_Check_Its_Variables tests that a pawn can be rendered and its variables can be checked
func (suite *PawnSuite) Test_Render_A_Pawn_And_Check_Its_Variables() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🐮🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`
	expectedColor := chess.Color(chess.WHITE)
	expectedType := chess.PieceType(chess.PAWN)

	err := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	output := suite.board.Render()
	color := suite.board.PieceAt(0, 1).Color()
	ptype := suite.board.PieceAt(0, 1).PType()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
	assert.Equal(suite.T(), expectedColor, color)
	assert.Equal(suite.T(), expectedType, ptype)
}

// Test_Move_A_Pawn_One_Step_WHITE tests that a pawn can be moved one step
func (suite *PawnSuite) Test_Move_A_Pawn_One_Step_WHITE() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🐮🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	err2 := suite.board.Move(0, 1, 0, 2)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Pawn_One_Step_BLACK tests that a pawn can be moved one step
func (suite *PawnSuite) Test_Move_A_Pawn_One_Step_BLACK() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🐮🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🤓🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	// Move a white pawn to make room for the black pawn
	_ = chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	_ = suite.board.Move(0, 1, 0, 2)

	err1 := chess.NewPawnPiece(&suite.board, 2, 6, chess.BLACK)
	err2 := suite.board.Move(2, 6, 2, 5)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_That_A_Blocked_Pawn_Cant_Move tests that a blocked pawn can't move
func (suite *PawnSuite) Test_That_A_Blocked_Pawn_Cant_Move() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🐮🔲🔳🔲🔳🔲🔳🔲 7
🤓🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	err2 := chess.NewPawnPiece(&suite.board, 0, 2, chess.BLACK)
	err3 := suite.board.Move(0, 1, 0, 2)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Pawn_Two_Steps tests that a pawn can be moved two steps
func (suite *PawnSuite) Test_Move_A_Pawn_Two_Steps() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🐮🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	err2 := suite.board.Move(0, 1, 0, 3)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Pawn_Three_Steps_Should_Fail tests that a pawn can't be moved three steps
func (suite *PawnSuite) Test_Move_A_Pawn_Three_Steps_Should_Fail() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🐮🔲🔳🔲🔳🔲🔳🔲 7
🔲🔳🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	err2 := suite.board.Move(0, 1, 0, 4)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.NotNil(suite.T(), err2)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Pawn_To_A_Invalid_Position tests that a pawn can't be moved to an invalid position
func (suite *PawnSuite) Test_Move_A_Pawn_To_A_Invalid_Position() {
	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)

	// Invalid position's to test
	err2 := suite.board.Move(0, 1, 0, 8)
	err3 := suite.board.Move(0, 1, 1, 1)
	err4 := suite.board.Move(0, 1, 0, 0)

	assert.Nil(suite.T(), err1)
	assert.NotNil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.NotNil(suite.T(), err4)
}

// Test_Move_A_Pawn_To_Strike_Opponent tests that a pawn can be moved to strike an opponent
func (suite *PawnSuite) Test_Move_A_Pawn_To_Strike_Opponent() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🔳🔲🔳🔲🔳🔲🔳🔲 7
🔲🐮🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	err2 := chess.NewPawnPiece(&suite.board, 1, 2, chess.BLACK)
	err3 := suite.board.Move(0, 1, 1, 2)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Nil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}

// Test_Move_A_Pawn_To_Strike_Yourself tests that a pawn can't be moved to strike yourself
func (suite *PawnSuite) Test_Move_A_Pawn_To_Strike_Yourself() {
	expected := `🔲🔳🔲🔳🔲🔳🔲🔳 8
🐮🔲🔳🔲🔳🔲🔳🔲 7
🔲🐮🔲🔳🔲🔳🔲🔳 6
🔳🔲🔳🔲🔳🔲🔳🔲 5
🔲🔳🔲🔳🔲🔳🔲🔳 4
🔳🔲🔳🔲🔳🔲🔳🔲 3
🔲🔳🔲🔳🔲🔳🔲🔳 2
🔳🔲🔳🔲🔳🔲🔳🔲 1
 A B C D E F G H
`

	err1 := chess.NewPawnPiece(&suite.board, 0, 1, chess.WHITE)
	err2 := chess.NewPawnPiece(&suite.board, 1, 2, chess.WHITE)
	err3 := suite.board.Move(0, 1, 1, 2)

	output := suite.board.Render()

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.NotNil(suite.T(), err3)
	assert.Equal(suite.T(), expected, output)
}
