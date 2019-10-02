package game

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
	"tictactoe/bot"
	"tictactoe/gen/models"
	"time"
)

const (
	URL                     string = "http://127.0.0.1:8000/api/v1/games" // Default URL
	NumberOfConcurrentUsers        = 100                                  // As Server doesn't implement throttling, Test sleeps for 1 Second after every 100th request
	UserSym                 rune   = 'X'
	BlankBoard              string = "---------"
)

// TestE2EFullGame, runs full end-to-end test
// It runs NumberOfConcurrentUsers number of || games
// As Server doesn't implement throttling, Test sleeps for 1 Second after every 100th request
// for every game,
// 1. It creates a game with POST request, extract the game ID from the response body "location" tag.
// 2. Sends GET request on URL with gameID to get the Game(board, ID, Status).
// 3. Make a move on the board, and send game in PUT
// 4. Step 3 continues, until Backend send game status as XWin, OWin or Draw
func TestE2EFullGame(t *testing.T) {
	var wg sync.WaitGroup
	players := 0

	// All the NumberOfConcurrentUsers of games will happen simultaneously.
	for i := 0; i < NumberOfConcurrentUsers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Create a Game. It will use POST method
			gameID, err := createAGame() // Test POST
			if err != nil {
				t.Error("error occurred in POST method creating a game ", err)
				return
			}

			// User must GET to get the board and game status
			game, err := getAGame(gameID)
			if err != nil || game == nil {
				t.Errorf("error occurred in GET a gameID %s err %+v", gameID, err)
				return
			}
			gameStatus := game.Status
			board := *game.Board
			// fmt.Printf("Bot played move! gameID %s board %s Status %s\n", gameID, board, gameStatus)
			if gameStatus != models.GameStatusRUNNING {
				t.Errorf("game must have been running here! gameID %s\n", gameID)
				return
			}

			// Simulate user moves util, the game is DRAW/OWIN/XWIN
			// It's PUT method testing
			// 'bot' package is used to simulate user moves. Same is used by backend also
			for gameStatus == models.GameStatusRUNNING {
				// User move
				board = bot.RobotMoveOptimum([]rune(board), UserSym)
				// fmt.Printf("User played move! gameID %s board %s Status %s\n", gameID, board, gameStatus)
				// PUT it to Robot, and receive the response
				game, err = playUserMove(gameID, board)
				if err != nil || game == nil {
					t.Errorf("error occurred in PUT a gameID %s err %+v", gameID, err)
					return
				}
				gameStatus = game.Status
				board = *game.Board
				if err != nil {
					t.Error("error occurred in PUT method playing a move ", err)
					return
				}
				// fmt.Printf("Bot played move! gameID %s board %s Status %s\n", gameID, board, gameStatus)
			}
			log.Printf("Final Game Status %+v %+v %s\n", gameID, gameStatus, board)
		}()
		players++
		if players%100 == 0 {
			time.Sleep(2 * time.Second)
		}
	}
	wg.Wait()
}

func TestGetAll(t *testing.T) {

	// Create 10 Games
	for i := 0; i < 10; i++ {
		_, err := createAGame()
		if err != nil {
			t.Error("error occurred in POST method creating a game ", err)
			return
		}
	}

	// Verify we get 10 games
	// GetAll deletes all games after getting the count. It's just for cleanup purpose
	err := getAllGames(10)
	if err != nil {
		t.Errorf("error occurred in GET err %+v", err)
		return
	}
}

func TestDelete(t *testing.T) {
	gameID, err := createAGame()
	if err != nil {
		t.Error("error occurred in POST method creating a game ", err)
		return
	}

	err = deleteAGame(gameID)
	if err != nil {
		t.Errorf("error occurred in Delete err %+v", err)
		return
	}

	// User must GET to get the board and game status
	game, err := getAGame(gameID)
	if err == nil || game != nil {
		t.Errorf("game found after delete gameID %s err %+v", gameID, err)
		return
	}
}

func getAllGames(num int) error {
	method := "GET"

	payload := strings.NewReader("")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(method, URL, payload)
	if err != nil {
		fmt.Println("error in getting a game ", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in sending a GET request ", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("GET response received err in statuscode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in reading a GET response body ", err)
		return err
	}

	games := make([]models.Game, num)
	err = json.Unmarshal(body, &games)
	if err != nil || games == nil {
		fmt.Println("error in unmarshal ", err)
		return err
	}

	count := 0
	for _, game := range games {
		if game.Status == models.GameStatusRUNNING {
			count++
		}
		err = deleteAGame(game.ID.String())
		if err != nil {
			fmt.Println("error in deleteAGame ", err)
		}
	}
	if count != num {
		return errors.New("error...invalid numbers of games found ")
	}

	return nil
}

func getGameURL(gameID string) string {
	return URL + "/" + gameID
}

func getAGame(gameID string) (*models.Game, error) {
	gameURL := getGameURL(gameID)
	method := "GET"

	payload := strings.NewReader("")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(method, gameURL, payload)
	if err != nil {
		fmt.Println("error in getting a game ", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in sending a GET request ", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("GET response received err in statuscode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in reading a GET response body ", err)
		return nil, err
	}

	var game models.Game
	err = json.Unmarshal(body, &game)
	if err != nil {
		fmt.Println("error in unmarshal ", err)
		return nil, err
	}

	return &game, nil
}

func deleteAGame(gameID string) error {
	gameURL := getGameURL(gameID)
	method := "DELETE"

	payload := strings.NewReader("")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(method, gameURL, payload)
	if err != nil {
		fmt.Println("error in DELETE a game ", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in sending a DELETE request ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Delete Failed")
	}

	return nil
}

func createAGame() (string, error) {
	method := "POST"

	// Bot will choose the step
	boardStr := bot.RobotMove([]rune(BlankBoard), UserSym)
	game := models.Game{Board: &boardStr}
	payload, _ := json.Marshal(&game)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("error in creating a POST request ", err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in sending a POST request ", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return "", errors.New("POST response received err in statuscode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in reading a POST response body ", err)
		return "", err
	}

	type locType struct {
		Location string `json:"location"`
	}
	var loc locType
	err = json.Unmarshal(body, &loc)
	if err != nil {
		fmt.Println("error in unmarshal ", err)
		return "", err
	}
	u, err := url.Parse(loc.Location)
	if err != nil {
		fmt.Println("error in parsing a POST response body ", err)
		return "", err
	}
	s := strings.Split(u.Path, "/")
	gameID := s[len(s)-1]
	// fmt.Printf("Created a game! gameID %s board %s\n", gameID, boardStr)

	return gameID, nil
}

func playUserMove(gameID string, currBoard string) (*models.Game, error) {
	gameURL := getGameURL(gameID)
	method := "PUT"

	game := models.Game{Board: &currBoard}
	payload, _ := json.Marshal(&game)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, gameURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("error in PUT request ", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in PUT request ", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("PUT response received err in statuscode")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in reading a PUT response body ", err)
		return nil, err
	}

	err = json.Unmarshal(body, &game)
	if err != nil {
		fmt.Println("error in unmarshal ", err)
		return nil, err
	}

	return &game, nil
}
