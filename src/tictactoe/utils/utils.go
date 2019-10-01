// Package utils, implements utility functions used by game TicTacToe.
package utils

import (
	"errors"
	"strings"
	"tictactoe/gen/models"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

const (
	Blank string = "---------" // Blank board state
)

// Winner return winner X or O.
func Winner(sym rune) string {
	if sym == 'X' {
		return models.GameStatusXWON
	}
	return models.GameStatusOWON
}

// GetGameStatus, takes board as input and returns game status - Running, Draw, XWins or OWins.
func GetGameStatus(board string) string {
	b := []rune(board)

	// Diagonal
	if b[0] != '-' && b[0] == b[4] && b[0] == b[8] {
		return Winner(b[0])
	}
	if b[2] != '-' && b[2] == b[4] && b[2] == b[6] {
		return Winner(b[2])
	}

	// Rows
	if b[0] != '-' && b[0] == b[1] && b[1] == b[2] {
		return Winner(b[0])
	}
	if b[3] != '-' && b[3] == b[4] && b[3] == b[5] {
		return Winner(b[3])
	}
	if b[6] != '-' && b[6] == b[7] && b[7] == b[8] {
		return Winner(b[6])
	}

	// Collumns
	if b[0] != '-' && b[0] == b[3] && b[3] == b[6] {
		return Winner(b[0])
	}
	if b[1] != '-' && b[1] == b[4] && b[1] == b[7] {
		return Winner(b[1])
	}
	if b[2] != '-' && b[2] == b[5] && b[2] == b[8] {
		return Winner(b[2])
	}

	for i := range board {
		if board[i] == '-' {
			return models.GameStatusRUNNING
		}
	}
	return models.GameStatusDRAW
}

// GetBkSym takes user symbol as input and returns robot symbol, default is X.
func GetBkSym(userSym rune) rune {
	if userSym == 'X' {
		return 'O'
	}
	return 'X'
}

// GenerateUUID, uses google/uuid package to generate a UUID.
func GenerateUUID() (strfmt.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return strfmt.UUID(""), err
	}

	return strfmt.UUID(id.String()), nil
}

// ValidateUserMove, validates User move
func ValidateUserMove(board string, before string) (rune, error) {
	userSym := rune('X')
	if len(board) != 9 {
		return userSym, errors.New("Invalid Input! board length is not 9")
	}

	for i := range board {
		if (board[i] != '-') && (board[i] != 'X') && (board[i] != 'O') {
			return userSym, errors.New("Invalid symbol in board")
		}
	}

	moves := 0
	for i := range board {
		if board[i] != before[i] {
			moves++
			if moves > 1 || before[i] != '-' {
				return userSym, errors.New("Invalid Input! You Can't play more than one move at once")
			}
			userSym = rune(board[i])
		}
	}

	if before == Blank {
		if board == Blank {
			return 'X', nil
		}
		if strings.ContainsRune(board, 'O') {
			return 'O', nil
		}
	} else {
		if moves == 0 {
			return userSym, errors.New("Please make your move")
		}
	}

	return userSym, nil
}
