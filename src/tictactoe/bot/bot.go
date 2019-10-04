// Package bot implements robot functionality for ticTocToe game.
// Currently, it makes random moves.
// Yet to implement AI to move optimally
package bot

import (
	"math/rand"
	"tictactoe/utils"
	"time"
)

// WinningCombo has list of combos. To win, one of the below combos must have same symbol (X or O)
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

// RobotMoveOptimum returns the best possible move on the board
// It ensures that player using this API never looses
// It is implemented using minimax backtracking algorithm.
// Alpha pruning algorithm is used to optimize the minimax algorithm.
// This function always consideres input 'sym' as Maximizer
func RobotMoveOptimum(b []rune, sym rune) string {
	bestMove := 0
	bestVal := -1000
	for i := range b {
		if b[i] == '-' {
			b[i] = sym
			moveVal := miniMax(b, 0, false, sym, -1000, 1000)
			b[i] = '-'
			if moveVal > bestVal {
				bestVal = moveVal
				bestMove = i
			}
		}
	}

	// Update the board with best move and return it to the user
	b[bestMove] = sym
	return string(b)
}

// evaluate is heuristic function to return score for a possible move made by minimax algo
// This function always consideres input 'sym' as Maximizer
func evaluate(board string, sym rune) int {
	b := []rune(board)
	for i := range WinningCombo {
		maxCount := 0
		minCount := 0
		for j := range WinningCombo[i] {
			if b[WinningCombo[i][j]] == '-' {
				break
			} else if b[WinningCombo[i][j]] == sym {
				maxCount++
			} else if b[WinningCombo[i][j]] == utils.GetBkSym(sym) {
				minCount++
			}
		}
		if maxCount == 3 {
			return 10
		} else if minCount == 3 {
			return -10
		}
	}

	return 0
}

// isMoveLeft returns if there are any blanks on the board.
func isMoveLeft(board string) bool {
	b := []rune(board)
	for i := range b {
		if b[i] == '-' {
			return true
		}
	}
	return false
}

// miniMax algo is used to find the best move for the 'sym' player.
// Therefore, input 'sym' is always considered as Maximizer
func miniMax(b []rune, depth int, isMax bool, sym rune, alpha, beta int) int {

	score := evaluate(string(b), sym)
	if score == 10 || score == -10 {
		return score
	}
	if !isMoveLeft(string(b)) {
		return 0
	}

	if isMax {
		best := -1000
		for i := range b {
			if b[i] == '-' {
				b[i] = sym
				moveVal := miniMax(b, depth+1, !isMax, sym, alpha, beta)
				b[i] = '-'
				if moveVal > best {
					best = moveVal
				}
				if moveVal > alpha {
					alpha = moveVal
				}
				if beta <= alpha {
					break
				}
			}
		}
		return best
	} else {
		best := 1000

		for i := range b {
			if b[i] == '-' {
				opp := utils.GetBkSym(sym)
				b[i] = opp
				moveVal := miniMax(b, depth+1, !isMax, sym, alpha, beta)
				b[i] = '-'
				if moveVal < best {
					best = moveVal
				}
				if moveVal < beta {
					beta = moveVal
				}
				if beta <= alpha {
					break
				}
			}
		}
		return best
	}
}

// RobotMove, makes a move on the board.
// Tt makes only random moves.
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

// RobotMoveOptimumTest make a move on board little optimally.
// It is used by E2E gotest player. It gives little bit chance for backend player to win.
// If E2E gotest player also uses RobotMoveOptimum(minimax), then all games are draw
// First, it makes such a move where it can win.
// if it can't win with one move, it checks, can opponent win with one move,
// if yes, then it makes a move on that board index to spoil his chance of winning
// If nothing above is possible, it make a random move
func RobotMoveOptimumTest(b []rune, sym rune) string {

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
	var index int
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
