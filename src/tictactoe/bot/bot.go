package bot

import (
	"math/rand"
	"time"
)

// Currently Robot moves randomly.
// Yet to implement AI to move optimally
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
