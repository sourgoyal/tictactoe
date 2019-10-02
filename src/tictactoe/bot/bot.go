// Package bot implements robot functionality for ticTocToe game.
// Currently, it makes random moves.
// Yet to implement AI to move optimally
package bot

import (
	"math/rand"
	"tictactoe/utils"
	"time"
)

var WinningCombo = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

// RobotMove, makes a move on the board.
// Currently, it makes a random moves.
func RobotMove(currBoard []rune, botSym rune) string {
	botBoard := make([]rune, 9)
	copy(botBoard, currBoard)

	var playSlice []int
	for i := range currBoard {
		if currBoard[i] == '-' {
			playSlice = append(playSlice, i)
		}
	}

	if len(playSlice) > 0 {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(len(playSlice))
		botBoard[playSlice[randIndex]] = botSym
		return string(botBoard)
	}

	return string(currBoard)
}

// RobotMoveOptimum make a move on board little optimally.
// First, it makes such a move where it can win.
// if it can't win with one move, it checks, can opponent win with one move,
// if yes, then it makes a move on that board index to spoil his chance of winning
// If nothing above is possible, it make a random move
func RobotMoveOptimum(b []rune, sym rune) string {

	// Win possiblity for self
	index := winPossiblity(b, sym)
	if index >= 9 {
		// Win possibility for opponent
		oppSym := utils.GetBkSym(sym)
		index = winPossiblity(b, oppSym)
	}

	if index < 9 {
		b[index] = sym
		return string(b)
	}

	return RobotMove(b, sym)
}

func winPossiblity(b []rune, sym rune) int {
	index := 10
	for i := range WinningCombo {
		count := 0
		index = 10
		for j := range WinningCombo[i] {
			if b[WinningCombo[i][j]] == '-' {
				index = WinningCombo[i][j]
			} else if b[WinningCombo[i][j]] == sym {
				count++
			}
		}
		if count == 2 && index < 9 {
			return index
		}
	}

	return 10
}
