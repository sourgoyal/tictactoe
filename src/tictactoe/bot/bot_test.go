package bot

import (
	"testing"
	"tictactoe/utils"
)

func TestRobotMove(t *testing.T) {

	testCaseSuccess := []string{
		"---------",
		"-------X-",
		"----X----",
		"X--------",
		"X-------X",
		"-XXXXXXXX",
	}
	testCaseFail := []string{
		"XXXXXXXXX",
	}

	for _, test := range testCaseSuccess {
		botMove := RobotMove([]rune(test), 'X')
		if _, err := utils.ValidateUserMove(botMove, test); err != nil {
			t.Error("Invalid Move made by Robot")
		}
	}

	for _, test := range testCaseFail {
		botMove := RobotMove([]rune(test), 'X')
		if _, err := utils.ValidateUserMove(botMove, test); err == nil {
			t.Error("Invalid Move made by Robot")
		}
	}
}
