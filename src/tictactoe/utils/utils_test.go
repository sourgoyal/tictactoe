package utils

import (
	"testing"
	"tictactoe/gen/models"
)

func TestGetBkSym(t *testing.T) {
	if GetBkSym('X') != 'O' {
		t.Error("Incorrect symbol selected")
	}
	if GetBkSym('O') != 'X' {
		t.Error("Incorrect symbol selected")
	}
	if GetBkSym('-') != 'X' {
		t.Error("Incorrect symbol selected")
	}
}

func TestGetGameStation(t *testing.T) {
	type gameStatus struct {
		board  string
		status string
	}
	testCases := []gameStatus{
		{"---------", models.GameStatusRUNNING},
		{"--------X", models.GameStatusRUNNING},
		{"-------X-", models.GameStatusRUNNING},
		{"-----X---", models.GameStatusRUNNING},
		{"---X-X-X-", models.GameStatusRUNNING},
		{"--O--O-X-", models.GameStatusRUNNING},
		{"X-O-X-X-X", models.GameStatusXWON},
		{"X-O-X-X-X", models.GameStatusXWON},
		{"XXX-O-O-X", models.GameStatusXWON},
		{"X-O-O-XXX", models.GameStatusXWON},
		{"X-O-O-OOO", models.GameStatusOWON},
		{"---OOOX--", models.GameStatusOWON},
		{"OOO-O-O--", models.GameStatusOWON},
		{"XOXOOXXXO", models.GameStatusDRAW},
	}

	for _, test := range testCases {
		recvStatus := GetGameStatus(test.board)
		if test.status != recvStatus {
			t.Errorf("Board %s, Expected Status %s, Received Status %s", test.board, test.status, recvStatus)
		}
	}
}

func TestValidateUserMove(t *testing.T) {
	type game struct {
		board    string
		before   string
		sym      rune
		expcPass bool
	}

	testCases := []game{
		{"X--------", "---------", 'X', true},
		{"X--------", "---------", 'X', true},
		{"X--------", "---------", 'X', true},
		{"X--------", "X--------", 'X', false},
		{"X-------", "X--------", 'X', false},
		{"X------OO", "X--------", 'X', false},
	}

	for i, test := range testCases {
		sym, err := ValidateUserMove(test.board, test.before)
		if test.expcPass {
			if err != nil || sym != test.sym {
				t.Error("Failed to validate for test ", i+1)
				return
			}
		} else {
			if err == nil {
				t.Error("Failed to validate for test ", i+1)
				return
			}
		}
	}
}
