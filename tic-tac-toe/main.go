package main

import "fmt"

type Board [3][3]string

func NewBoard() *Board {
	var board Board

	for i := range board {
		for j := range board[i] {
			board[i][j] = "_"
		}
	}

	return &board
}

type Player struct {
	Name string
	Mark string
}

func NewPlayer(name string, mark string) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}
}

type GameController struct {
	board           *Board
	players         [2]*Player
	nextPlayerIndex int
}

func NewGameController(player1, player2 *Player) *GameController {
	return &GameController{
		players:         [2]*Player{player1, player2},
		board:           NewBoard(),
		nextPlayerIndex: 0,
	}
}

func (c *GameController) PrintBoard() {
	for r := range c.board {
		fmt.Println(c.board[r])
	}
}

func (c *GameController) Start() {
	for !c.IsGameOver() {
		c.PrintBoard()
		c.NextTurn()
	}
	c.PrintBoard()

	fmt.Println("Game Over")
}

func (c *GameController) NextTurn() {
	currentPlayer := c.players[c.nextPlayerIndex]

	fmt.Printf("Player %s- %s turn\n", currentPlayer.Name, currentPlayer.Mark)

	var row, col int
	// prompt player to input
	fmt.Printf("Enter row and column to place marker, e.g. 1 1 for 1st row and 1st column\n")

	_, err := fmt.Scanln(&row, &col)
	if err != nil {
		fmt.Println("Error scanning marker for row and column ")
		return
	}

	if !c.IsValidMove(row-1, col-1) {
		fmt.Println("Please enter a valid move\n")
		return
	}

	c.board[row-1][col-1] = currentPlayer.Mark

	c.nextPlayerIndex = (c.nextPlayerIndex + 1) % 2

}

func (c *GameController) IsValidMove(row, col int) bool {
	if row < 0 || row >= len(c.board) || col < 0 || col >= len(c.board[0]) {
		return false
	}

	if c.board[row][col] != "_" {
		return false
	}

	return true

}

func (c *GameController) IsGameOver() bool {
	if c.IsMarkWonGame(c.players[0].Mark) {
		fmt.Printf("%s-%s won game\n ", c.players[0].Name, c.players[0].Mark)
		return true
	} else if c.IsMarkWonGame(c.players[1].Mark) {
		fmt.Printf("%s-%s won game\n ", c.players[1].Name, c.players[1].Mark)
		return true

	}

	if c.IsBoardFull() {
		fmt.Println("Board is full\n")
		return true
	}
	return false
}

func (c *GameController) IsMarkWonGame(mark string) bool {
	board := c.board
	// check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == mark && board[i][1] == mark && board[i][2] == mark {
			return true
		}
	}

	// check cols
	for j := 0; j < 3; j++ {
		if board[0][j] == mark && board[1][j] == mark && board[2][j] == mark {
			return true
		}
	}

	// check daigonals
	if board[0][0] == mark && board[1][1] == mark && board[2][2] == mark {
		return true
	}

	if board[0][2] == mark && board[1][1] == mark && board[2][0] == mark {
		return true
	}

	return false
}

func (c *GameController) IsBoardFull() bool {
	for row := range c.board {
		for col := range c.board[row] {
			if c.board[row][col] == "_" {
				return false
			}
		}
	}
	return true
}

func main() {
	player1 := NewPlayer("Ram", "X")
	player2 := NewPlayer("Sam", "Y")

	gc := NewGameController(player1, player2)
	gc.Start()
}
