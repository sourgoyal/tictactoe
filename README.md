# Game tic-tac-toe
REST backend to implement game of tic-tac-toe

## Implementation 
- Go server code is generated using swagger spec file.
- For backend moves, bot package is created. It makes a move on the basic of win possibility of self and opponent and then make a move.
- Dockerfile is provided to run it in docker container. 
- playgame_test.go <TestE2EFullGame> runs, end-to-end test. We can define NumberOfConcurrentUsers in constants. It represents number of ongoing game test has to play. Each game is full simulation till the end result of win or draw. Test also uses bot package to make a move. 

# Game and code flow
- The client (player) starts a game by making a POST request to /games. The POST request contains a representation of a game board, either empty (computer starts) or with the first move made (player starts). The player/computer can choose either noughts or crosses.
- The backend responds with the location URL of the started game.
```bash
tictactoe/src/tictactoe/playgame/playgame.go
(g *GameInfo) CreateGame(game *models.Game, reqURL *url.URL) middleware.Responder 
```
- Client GETs the board state from the URL.
```bash
tictactoe/src/tictactoe/playgame/playgame.go
(g *GameInfo) GetGameDetails(ID strfmt.UUID) middleware.Responder
```
- Client PUTs the board state with a new move to the URL.
- Backend validates the move, makes it's own move and updates the game state.
  The updated game state is returned in the PUT response.
```bash
tictactoe/src/tictactoe/playgame/playgame.go
(g *GameInfo) PlayGame(gameID strfmt.UUID, game *models.Game) middleware.Responder
```
- And so on. The game is over once the computer or the player gets 3 noughts
  or crosses, horizontally, vertically or diagonally or there are no moves to
  be made.
- Game is deleted from db once game is over, or inactive for continous 10 mins. Although user can delete it using DELETE API 
```bash
tictactoe/src/tictactoe/playgame/playgame.go
(g *GameInfo) DeleteGame(ID strfmt.UUID) middleware.Responder
```

# Install
### Run in a docker container 
```bash
docker build -t tictactoe .
docker run -it --publish 8080:8080 --name tictactoe --rm tictactoe
```
### Run in local
```bash 
. dep-ensure.sh
. install.shs
```

# North Star 
- Add AI functionality into bot package. 
- Implement throttling
- Deploy using k8s and Helm
- Use mongodb instead of own db implementation

# License
MIT License



