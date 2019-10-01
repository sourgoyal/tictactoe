// Package bot implements robot functionality for ticTocToe game.
// Currently, it makes random moves.
// Yet to implement AI to move optimally
package bot

import (
	"math/rand"
	"time"
)

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
