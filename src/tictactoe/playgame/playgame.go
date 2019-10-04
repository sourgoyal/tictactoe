// Package game implements the tictoctoe handler functions.
package game

import (
	"log"
	"net/url"
	"tictactoe/bot"
	"tictactoe/db"
	"tictactoe/gen/models"
	"tictactoe/gen/restapi/operations"
	"tictactoe/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// GameInfo contains db and handler functions interface
type GameInfo struct {
	db   *db.GamesDB
	Info GameOper
}

// GameOper implements interface for all handler functions
type GameOper interface {
	GetAllGames() middleware.Responder
	GetGameDetails(strfmt.UUID) middleware.Responder
	DeleteGame(strfmt.UUID) middleware.Responder
	CreateGame(models.Game, *url.URL) middleware.Responder
	PlayGame(strfmt.UUID, models.Game) middleware.Responder
}

// New create a new instance of tictoctoe games backend
func New() *GameInfo {
	game := &GameInfo{db: db.NewGameDB()}
	go game.db.Cleanup()
	return game
}

// CreateGame is called on POST request from HTTP client.
// It create a new game and stores it into db.
func (g *GameInfo) CreateGame(game *models.Game, reqURL *url.URL) middleware.Responder {
	// Create a UUID to represent as game ID.
	gameID, err := utils.GenerateUUID()
	if err != nil {
		return operations.NewPostAPIV1GamesInternalServerError()
	}

	// It's newly created gameID and must not be present in game db. If exists return with internal server error.
	if g.db.GameExists(gameID) {
		return operations.NewPostAPIV1GamesInternalServerError()
	}

	// board received in POST request.
	// User may move her step and POST also
	board := *game.Board

	// Validate user input board and get the symbol used by user.
	// It returns defailt symbol as 'X' if user has not made any move
	userSym, err := utils.ValidateUserMove(board, utils.Blank)
	if err != nil {
		// Return HTTP BAD request error with reason
		return operations.NewPostAPIV1GamesBadRequest().WithPayload(&operations.PostAPIV1GamesBadRequestBody{Reason: err.Error()})
	}

	// Get backend symbol
	bkSym := utils.GetBkSym(userSym)
	// RobotMove implements a robot functionality. It decides on which move to take on the board for backend
	bkBoard := bot.RobotMoveOptimum([]rune(board), bkSym)

	// Add game to db
	backendGame := models.Game{Board: &bkBoard, ID: gameID, Status: models.GameStatusRUNNING}
	gameD := db.GameData{UserSymbol: userSym, UserTurn: true, Game: backendGame}
	if errCode := g.db.AddGame(gameD, gameID); errCode != db.Success {
		return operations.NewPostAPIV1GamesInternalServerError()
	}

	// As per requiement, send game URL in header 'location' and body 'location'
	respURL := reqURL
	respURL.Path = respURL.Path + "/" + string(gameID)
	respPost := operations.NewPostAPIV1GamesCreated()
	respPost.SetLocation(reqURL.String())
	respPost.SetPayload(&operations.PostAPIV1GamesCreatedBody{Location: respURL.String()})
	return respPost
}

// GetGameDetails is called on '/api/v1/games/{game_id}:' GET request.
// it returns a requested game ID details.
func (g *GameInfo) GetGameDetails(ID strfmt.UUID) middleware.Responder {
	var errCode db.GameDBErrorCode
	var game models.Game
	if game, errCode = g.db.GetGame(ID); errCode != db.Success {
		return operations.NewGetAPIV1GamesGameIDNotFound()
	}

	return operations.NewGetAPIV1GamesGameIDOK().WithPayload(&game)
}

// GetAllGames is called on '/api/v1/games/' GET request.
// Successful response, returns an array of games, returns an empty array if no users found.
func (g *GameInfo) GetAllGames() middleware.Responder {
	var payload []*models.Game

	g.db.Range(func(game db.GameData) {
		payload = append(payload, &game.Game)
	})
	return operations.NewGetAPIV1GamesOK().WithPayload(payload)
}

// DeleteGame is called on '/api/v1/games/{game_id}:' DELETE request.
// It deletes a requested game from db.
func (g *GameInfo) DeleteGame(ID strfmt.UUID) middleware.Responder {
	if errCode := g.db.DeleteGame(ID); errCode != db.Success {
		return operations.NewDeleteAPIV1GamesGameIDNotFound()
	}

	return operations.NewDeleteAPIV1GamesGameIDOK()
}

// PlayGame is called on '/api/v1/games/{game_id}:' PUT request.
// It receives a game input from user. Make a move and returns game output in response body to user.
func (g *GameInfo) PlayGame(gameID strfmt.UUID, game *models.Game) middleware.Responder {
	// Fetch game details from db
	existGame, errCode := g.db.GetGameData(gameID)
	if errCode != db.Success {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: "This game doesn't exist"})
	}

	// When we receive PUT request, it must be user turn in db.
	if !existGame.UserTurn {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: "Please wait for your turn"})
	}

	// Validate user input
	userSym, err := utils.ValidateUserMove(*game.Board, *existGame.Game.Board)
	if err != nil {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: err.Error()})
	}
	if userSym != existGame.UserSymbol {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: "Invalid Symbol used"})
	}

	// Check board for any win or draw
	gameStatus := utils.GetGameStatus(*game.Board)
	backendGame := models.Game{ID: gameID}
	if gameStatus == models.GameStatusRUNNING {
		bkSym := utils.GetBkSym(userSym)

		// make a backend move using bot.RobotMoveOptimum()
		bkBoard := bot.RobotMoveOptimum([]rune(*game.Board), bkSym)
		backendGame.Board = &bkBoard

		// Check board for any win or draw after backend move
		gameStatus = utils.GetGameStatus(bkBoard)
	} else {
		backendGame.Board = game.Board
	}

	// Update the db with new board and details if game is still running
	backendGame.Status = gameStatus
	gameD := db.GameData{UserSymbol: userSym, UserTurn: true, Game: backendGame}
	if gameStatus == models.GameStatusRUNNING {
		if errCode := g.db.UpdateGame(gameD, gameID); errCode != db.Success {
			return operations.NewPutAPIV1GamesGameIDInternalServerError()
		}
	} else {
		// Deletes the game from db if game is draw or any win
		if errCode := g.db.DeleteGame(gameID); errCode != db.Success {
			log.Printf("Failed to delete after Game over. Background cleanup must clean this game %+v", gameID)
		}
	}
	return operations.NewPutAPIV1GamesGameIDOK().WithPayload(&backendGame)
}
