package gamesDB

import (
	"sync"
	"tictactoe/gen/models"
	"time"

	"github.com/go-openapi/strfmt"
)

type GameDBErrorCode uint8

const (
	Success                  GameDBErrorCode = 0
	CreateFailedGameExists   GameDBErrorCode = 1
	UpdateFailedGameNotFound GameDBErrorCode = 2
	DeleteFailedGameNotFound GameDBErrorCode = 3
	GameNotFound             GameDBErrorCode = 4
)

type GameData struct {
	Game            models.Game
	UserSymbol      rune // 'X' or 'O'
	UserTurn        bool // User: true or Backend: false
	ExpireTimeStamp int64
}

type GamesDB struct {
	gamesMap map[strfmt.UUID]GameData
	mux      sync.RWMutex
}

func NewGameDB() *GamesDB {
	db := &GamesDB{gamesMap: make(map[strfmt.UUID]GameData)}
	return db
}

func (gamesDB *GamesDB) Total() int {
	gamesDB.mux.RLock()
	defer gamesDB.mux.RUnlock()

	return len(gamesDB.gamesMap)
}

func (gamesDB *GamesDB) GetGame(gameID strfmt.UUID) (models.Game, GameDBErrorCode) {
	gameD, err := gamesDB.GetGameData(gameID)
	var game models.Game
	if err == Success {
		return gameD.Game, Success
	}
	return game, err
}

func (gamesDB *GamesDB) GetGameData(gameID strfmt.UUID) (GameData, GameDBErrorCode) {
	gamesDB.mux.RLock()
	defer gamesDB.mux.RUnlock()
	game, exists := gamesDB.gamesMap[gameID]
	if !exists {
		return game, GameNotFound
	}

	return game, Success
}

func (gamesDB *GamesDB) GameExists(gameID strfmt.UUID) bool {
	gamesDB.mux.RLock()
	defer gamesDB.mux.RUnlock()
	_, exists := gamesDB.gamesMap[gameID]
	return exists
}

func (gamesDB *GamesDB) AddGame(gameData GameData, gameID strfmt.UUID) GameDBErrorCode {
	if gamesDB.GameExists(gameID) {
		return CreateFailedGameExists
	}

	gamesDB.mux.Lock()
	defer gamesDB.mux.Unlock()
	gameData.ExpireTimeStamp = time.Now().Add(10 * time.Minute).UnixNano()
	gamesDB.gamesMap[gameID] = gameData
	return Success
}

func (gamesDB *GamesDB) UpdateGame(gameData GameData, gameID strfmt.UUID) GameDBErrorCode {
	if !gamesDB.GameExists(gameID) {
		return UpdateFailedGameNotFound
	}

	gamesDB.mux.Lock()
	defer gamesDB.mux.Unlock()
	gameData.ExpireTimeStamp = time.Now().Add(10 * time.Minute).UnixNano()
	gamesDB.gamesMap[gameID] = gameData
	return Success
}

func (gamesDB *GamesDB) DeleteGame(gameID strfmt.UUID) GameDBErrorCode {
	if !gamesDB.GameExists(gameID) {
		return DeleteFailedGameNotFound
	}

	gamesDB.mux.Lock()
	defer gamesDB.mux.Unlock()
	delete(gamesDB.gamesMap, gameID)
	return Success
}

func (gamesDB *GamesDB) Range(cb func(GameData)) {
	gamesDB.mux.RLock()

	for _, game := range gamesDB.gamesMap {
		gamesDB.mux.RUnlock()
		cb(game)
		gamesDB.mux.RLock()
	}
	gamesDB.mux.RUnlock()
}

func (gamesDB *GamesDB) Cleanup() {

	for {
		time.Sleep(10 * time.Minute)

		gamesDB.Range(func(game GameData) {
			if time.Now().UnixNano() > game.ExpireTimeStamp {
				gamesDB.DeleteGame(game.Game.ID)
			}
		})
	}
}
