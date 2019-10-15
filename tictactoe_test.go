package ticTacToe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializedGameBoardSuccess(t *testing.T) {
	var initChar rune
	var comparison bool
	var i uint8
	var j uint8
	initChar = '-'
	InitializeGameboard()

	for i = 0; i < Rows; i++ {
		for j = 0; j < Columns; j++ {
			comparison = (initChar == Gameboard[i][j])
			assert.True(t, comparison, "rune should be '-'")
		}
	}

}
func TestInitializedGameBoardFail(t *testing.T) {
	var initChar rune
	var comparison bool
	var i uint8
	var j uint8
	initChar = 'x'
	InitializeGameboard()

	for i = 0; i < Rows; i++ {
		for j = 0; j < Columns; j++ {
			comparison = (initChar == Gameboard[i][j])
			assert.False(t, comparison, "Wrong Character")
		}
	}

}
func TestRow1Win_x(t *testing.T) {
	var win bool
	var winner rune
	var correctWinner bool
	InitializeGameboard()
	Gameboard[0][0] = 'x'
	Gameboard[0][1] = 'x'
	Gameboard[0][2] = 'x'
	win, winner = CheckForWin()
	correctWinner = (winner == 'x')
	assert.True(t, win, "should see a win")
	assert.True(t, correctWinner, "should see a true")
}
func TestRow2Win_x(t *testing.T) {
	var win bool
	var winner rune
	var correctWinner bool
	InitializeGameboard()
	Gameboard[1][0] = 'x'
	Gameboard[1][1] = 'x'
	Gameboard[1][2] = 'x'
	win, winner = CheckForWin()
	correctWinner = (winner == 'x')
	assert.True(t, win, "should see a win")
	assert.True(t, correctWinner, "should see a true")
}

func TestRow3Win_x(t *testing.T) {
	var win bool
	var winner rune
	var correctWinner bool
	InitializeGameboard()
	Gameboard[2][0] = 'x'
	Gameboard[2][1] = 'x'
	Gameboard[2][2] = 'x'
	win, winner = CheckForWin()
	correctWinner = (winner == 'x')
	assert.True(t, win, "should see a win")
	assert.True(t, correctWinner, "should see a true")
}
func TestRow1Loss_x(t *testing.T) {
	var win bool
	var winner rune
	var correctWinner bool
	InitializeGameboard()
	Gameboard[0][0] = 'o'
	Gameboard[0][1] = 'x'
	Gameboard[0][2] = 'o'
	win, winner = CheckForWin()
	correctWinner = (winner == '-')
	assert.False(t, win, "should see a win")
	assert.True(t, correctWinner, "should see a -")
}
func TestRow2Loss_x(t *testing.T) {
	var win bool
	var winner rune
	var correctWinner bool
	InitializeGameboard()
	Gameboard[1][0] = 'x'
	Gameboard[1][1] = 'o'
	Gameboard[1][2] = 'o'
	win, winner = CheckForWin()
	correctWinner = (winner == '-')
	assert.False(t, win, "should see a loss")
	assert.True(t, correctWinner, "should see a -")
}

func TestRow3Loss_x(t *testing.T) {
	var win bool
	var winner rune
	var correctWinner bool
	InitializeGameboard()
	Gameboard[2][0] = 'x'
	Gameboard[2][1] = 'x'
	Gameboard[2][2] = 'o'
	win, winner = CheckForWin()
	correctWinner = (winner == '-')
	assert.False(t, win, "should see a loss")
	assert.True(t, correctWinner, "should see a -")
}
