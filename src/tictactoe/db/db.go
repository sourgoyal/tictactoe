package db

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

func (db *GamesDB) Total() int {
	db.mux.RLock()
	defer db.mux.RUnlock()

	return len(db.gamesMap)
}

func (db *GamesDB) GetGame(gameID strfmt.UUID) (models.Game, GameDBErrorCode) {
	gameD, err := db.GetGameData(gameID)
	var game models.Game
	if err == Success {
		return gameD.Game, Success
	}
	return game, err
}

func (db *GamesDB) GetGameData(gameID strfmt.UUID) (GameData, GameDBErrorCode) {
	db.mux.RLock()
	defer db.mux.RUnlock()
	game, exists := db.gamesMap[gameID]
	if !exists {
		return game, GameNotFound
	}

	return game, Success
}

func (db *GamesDB) GameExists(gameID strfmt.UUID) bool {
	db.mux.RLock()
	defer db.mux.RUnlock()
	_, exists := db.gamesMap[gameID]
	return exists
}

func (db *GamesDB) AddGame(gameData GameData, gameID strfmt.UUID) GameDBErrorCode {
	if db.GameExists(gameID) {
		return CreateFailedGameExists
	}

	db.mux.Lock()
	defer db.mux.Unlock()
	gameData.ExpireTimeStamp = time.Now().Add(10 * time.Minute).UnixNano()
	db.gamesMap[gameID] = gameData
	return Success
}

func (db *GamesDB) UpdateGame(gameData GameData, gameID strfmt.UUID) GameDBErrorCode {
	if !db.GameExists(gameID) {
		return UpdateFailedGameNotFound
	}

	db.mux.Lock()
	defer db.mux.Unlock()
	gameData.ExpireTimeStamp = time.Now().Add(10 * time.Minute).UnixNano()
	db.gamesMap[gameID] = gameData
	return Success
}

func (db *GamesDB) DeleteGame(gameID strfmt.UUID) GameDBErrorCode {
	if !db.GameExists(gameID) {
		return DeleteFailedGameNotFound
	}

	db.mux.Lock()
	defer db.mux.Unlock()
	delete(db.gamesMap, gameID)
	return Success
}

func (db *GamesDB) Range(cb func(GameData)) {
	db.mux.RLock()

	for _, game := range db.gamesMap {
		db.mux.RUnlock()
		cb(game)
		db.mux.RLock()
	}
	db.mux.RUnlock()
}

func (db *GamesDB) Cleanup() {

	for {
		time.Sleep(10 * time.Minute)

		db.Range(func(game GameData) {
			if time.Now().UnixNano() > game.ExpireTimeStamp {
				db.DeleteGame(game.Game.ID)
			}
		})
	}
}
