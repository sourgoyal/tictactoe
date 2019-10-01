# Game tic-tac-toe
REST backend to implement game of tic-tac-toe

# Implementation 
- Go server code is generated using swagger spec file.
- For backend moves, bot package is created. Currently it randomises the moves on the board. But further it can be inhanced for optimum moves.
- Dockerfile is provided to run it in docker container. 
- playgame_test.go <TestE2EFullGame> runs, end-to-end test. We can define NumberOfConcurrentUsers in constants. It represents number of ongoing game test has to play. Each game is full simulation till the end result of win or draw. Test also uses bot package to make a move. 
- Use godoc to generate the documentation. 

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
- Add AI functionality into bot package, which currently randomise the backend moves on the board
- Implement throttling
- Deploy using k8s and Helm
- Use mongodb instead of own db implementation

# License
MIT License



