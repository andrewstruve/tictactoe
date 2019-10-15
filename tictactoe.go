package ticTacToe

const Rows = 3
const Columns = 3

var Gameboard [Rows][Columns]rune

// initialize gameboard
func InitializeGameboard() {
	var i int
	var j int
	for i = 0; i < Rows; i++ {
		for j = 0; j < Columns; j++ {
			Gameboard[i][j] = '-'
		}
	}
}

// check for win
// true for win , false for lose
func CheckForWin() (bool, rune) {
	var win bool
	var winner rune
	// check rows
	win, winner = CheckRows()
	if win {
		return win, winner
	}
	// check columns
	win, winner = CheckColumns()
	if win {
		return win, winner
	}
	//check left to right diag
	win, winner = CheckLeftToRightDiag()
	if win {
		return win, winner
	}
	//check right to left diag
	win, winner = CheckRightToLeftDiag()
	return win, winner
}
func CheckRows() (bool, rune) {
	var i uint8
	var winner rune
	winner = '-'
	var match bool
	// check rows
	for i = 0; i < Rows; i++ {
		if Gameboard[i][0] == 'x' && Gameboard[i][1] == 'x' && Gameboard[i][2] == 'x' {
			winner = 'x'
			match = true
		} else if Gameboard[i][0] == 'o' && Gameboard[i][1] == 'o' && Gameboard[i][2] == 'o' {
			winner = 'o'
			match = true
		}
		if match == true {
			break
		}
	}

	return match, winner

}
func CheckColumns() (bool, rune) {
	var i uint8
	var winner rune
	winner = '-'
	var match bool
	// check rows
	for i = 0; i < Columns; i++ {
		if Gameboard[0][i] == 'x' && Gameboard[1][i] == 'x' && Gameboard[2][i] == 'x' {
			winner = 'x'
			match = true
		} else if Gameboard[0][i] == 'o' && Gameboard[1][i] == 'o' && Gameboard[2][i] == 'o' {
			winner = 'o'
			match = true
		}
		if match == true {
			break
		}
	}

	return match, winner

}
func CheckLeftToRightDiag() (bool, rune) {
	var winner rune
	winner = '-'
	var match bool
	match = false
	// check rows

	if Gameboard[0][0] == 'x' && Gameboard[1][1] == 'x' && Gameboard[2][2] == 'x' {
		winner = 'x'
		match = true
	} else if Gameboard[0][0] == 'o' && Gameboard[1][1] == 'o' && Gameboard[2][2] == 'o' {
		winner = 'o'
		match = true
	}

	return match, winner
}
func CheckRightToLeftDiag() (bool, rune) {
	var winner rune
	winner = '-'
	var match bool
	match = false
	// check rows
	if Gameboard[0][2] == 'x' && Gameboard[1][1] == 'x' && Gameboard[2][0] == 'x' {
		winner = 'x'
		match = true
	} else if Gameboard[0][2] == 'o' && Gameboard[1][1] == 'o' && Gameboard[2][0] == 'o' {
		winner = 'o'
		match = true
	}

	return match, winner
}
