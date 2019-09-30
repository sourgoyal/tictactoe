package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"tictactoe/gen/models"
	"time"

	"github.com/go-openapi/strfmt"
)

func TestCreateGame(t *testing.T) {

	url := "http://127.0.0.1:3000/api/v1/games"
	method := "POST"

	payload := strings.NewReader("{\n\"board\": \"-------X-\"\n}")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestGetAllGame(t *testing.T) {

	url := "http://127.0.0.1:3000/api/v1/games"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	games := make([]models.Game, 100)
	json.Unmarshal(body, &games)
	var id strfmt.UUID
	var board string
	for _, game := range games {
		fmt.Println(game)
		id = game.ID
		board = *game.Board
	}
	fmt.Println(id)

	getURL := url + "/" + id.String()
	if err := getAGame(getURL); err != nil {
		t.Errorf("Test Failed Error: %+v", err.Error())
	}

	if err := putAGame(getURL, []rune(board)); err != nil {
		t.Errorf("Test Failed Error: %+v", err.Error())
	}

	if err := deleleAGame(getURL); err != nil {
		t.Errorf("Test Failed Error: %+v", err.Error())
	}

	if err := getAGame(getURL); err != nil {
		t.Errorf("Test Failed Error: %+v", err.Error())
	}
}

func getAGame(url string) error {
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return nil
}

func deleleAGame(url string) error {
	method := "DELETE"

	payload := strings.NewReader("")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return nil
}

func putAGame(url string, currBoard []rune) error {

	method := "PUT"

	bkBoard := make([]rune, 9)
	userSym := 'X'
	copy(bkBoard, currBoard)

	var playSlice []int
	for i := range currBoard {
		if currBoard[i] == '-' {
			playSlice = append(playSlice, i)
		}
	}
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(playSlice) - 1)
	bkBoard[randIndex] = userSym
	boardStr := string(bkBoard)
	fmt.Printf("After User move %s\n", boardStr)

	game := models.Game{Board: &boardStr}
	payload, _ := json.Marshal(&game)

	// payload := strings.NewReader("{\n        \"board\": \"O------XX\"\n}")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return nil
}
