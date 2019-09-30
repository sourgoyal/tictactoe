package game

import (
	"log"
	"net/url"
	"tictactoe/db"
	"tictactoe/gen/models"
	"tictactoe/gen/restapi/operations"
	"tictactoe/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// DNS Info data structure
type GameInfo struct {
	db   *db.GamesDB
	Info GameOper
}

// DNS Operations
type GameOper interface {
	GetAllGames() middleware.Responder
	GetGameDetails(strfmt.UUID) middleware.Responder
	DeleteGame(strfmt.UUID) middleware.Responder
	CreateGame(models.Game, *url.URL) middleware.Responder
	PlayGame(strfmt.UUID, models.Game) middleware.Responder
}

// Create a new instance of DNS
func New() *GameInfo {
	game := &GameInfo{db: db.NewGameDB()}
	go game.db.Cleanup()
	return game
}

func (g *GameInfo) PlayGame(gameID strfmt.UUID, game *models.Game) middleware.Responder {
	existGame, errCode := g.db.GetGameData(gameID)

	if errCode != db.Success {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: "This game doesn't exist"})
	}

	if !existGame.UserTurn {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: "Please wait for your turn"})
	}

	userSym, err := utils.ValidateUserMove(*game.Board, *existGame.Game.Board)
	if err != nil {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: err.Error()})
	}

	if userSym != existGame.UserSymbol {
		return operations.NewPutAPIV1GamesGameIDBadRequest().WithPayload(&operations.PutAPIV1GamesGameIDBadRequestBody{Reason: "Invalid Symbol used"})
	}

	gameStatus := utils.GetGameStatus(*game.Board)
	backendGame := models.Game{ID: gameID}
	backendGame.Board = game.Board
	if gameStatus == models.GameStatusRUNNING {
		bkSym := utils.GetBkSym(userSym)
		bkBoard := utils.BackendMove([]rune(*game.Board), bkSym)
		backendGame.Board = &bkBoard
		gameStatus = utils.GetGameStatus(bkBoard)
	}

	backendGame.Status = gameStatus
	gameD := db.GameData{UserSymbol: userSym, UserTurn: true, Game: backendGame}
	if gameStatus == models.GameStatusRUNNING {
		if errCode := g.db.UpdateGame(gameD, gameID); errCode != db.Success {
			return operations.NewPutAPIV1GamesGameIDInternalServerError()
		}
	} else {
		if errCode := g.db.DeleteGame(gameID); errCode != db.Success {
			log.Printf("Failed to delete after Game over. Background cleanup must clean this game %+v", gameID)
		}
	}
	log.Printf("PUT GameID %+v Received Board %s, userSym %#U, CSBoard %s", gameID, *game.Board, userSym, *backendGame.Board)
	return operations.NewPutAPIV1GamesGameIDOK().WithPayload(&backendGame)
}

func (g *GameInfo) CreateGame(game *models.Game, reqURL *url.URL) middleware.Responder {
	id, err := uuid.NewUUID()
	if err != nil {
		return operations.NewPostAPIV1GamesInternalServerError()
	}
	gameID := strfmt.UUID(id.String())

	if g.db.GameExists(gameID) {
		return operations.NewPostAPIV1GamesInternalServerError()
	}

	board := game.Board

	userSym, err := utils.ValidateUserMove(*board, utils.Blank)
	if err != nil {
		return operations.NewPostAPIV1GamesBadRequest().WithPayload(&operations.PostAPIV1GamesBadRequestBody{Reason: err.Error()})
		// return operations.NewPostAPIV1GamesBadRequest()
	}

	bkSym := utils.GetBkSym(userSym)
	bkBoard := utils.BackendMove([]rune(*board), bkSym)

	backendGame := models.Game{Board: &bkBoard, ID: gameID, Status: models.GameStatusRUNNING}
	gameD := db.GameData{UserSymbol: userSym, UserTurn: true, Game: backendGame}
	if errCode := g.db.AddGame(gameD, gameID); errCode != db.Success {
		return operations.NewPostAPIV1GamesInternalServerError()
	}

	log.Printf("Created a new gameID %+v", gameID)
	log.Printf("Received Board %s, userSym %#U, URL %s", *board, userSym, reqURL.String())
	respURL := reqURL
	respURL.Path = respURL.Path + "/" + string(gameID)
	respPost := operations.NewPostAPIV1GamesCreated()
	respPost.SetLocation(reqURL.String())
	respPost.SetPayload(&operations.PostAPIV1GamesCreatedBody{Location: respURL.String()})
	return respPost
}

func (g *GameInfo) DeleteGame(ID strfmt.UUID) middleware.Responder {
	if errCode := g.db.DeleteGame(ID); errCode != db.Success {
		return operations.NewDeleteAPIV1GamesGameIDNotFound()
	}

	return operations.NewDeleteAPIV1GamesGameIDOK()
}
func (g *GameInfo) GetGameDetails(ID strfmt.UUID) middleware.Responder {
	var errCode db.GameDBErrorCode
	var game models.Game
	if game, errCode = g.db.GetGame(ID); errCode != db.Success {
		return operations.NewGetAPIV1GamesGameIDNotFound()
	}

	return operations.NewGetAPIV1GamesGameIDOK().WithPayload(&game)
}

func (g *GameInfo) GetAllGames() middleware.Responder {
	var payload []*models.Game

	g.db.Range(func(game db.GameData) {
		payload = append(payload, &game.Game)
	})
	return operations.NewGetAPIV1GamesOK().WithPayload(payload)
}
