package main

import (
	"fmt"
	"strconv"
)

const (
	BoardSize = 9
)

type ticTacBox struct {
	occupier string //TODO: Should be typed to only X or O's
}

type board struct {
	squares []ticTacBox
}

func (b board) printRow(one, two, three string) {
	fmt.Println("", one, "|", two, "|", three)
}

func (b board) printBoard(args ...string) {
	b.printRow(args[0], args[1], args[2])
	fmt.Println("-----------")
	b.printRow(args[3], args[4], args[5])
	fmt.Println("-----------")
	b.printRow(args[6], args[7], args[8])
}

func (b board) printGameBoard() {
	var results []string
	for index := range b.squares {
		if b.squares[index].occupier != "" {
			results = append(results, b.squares[index].occupier)
		} else {

			results = append(results, strconv.Itoa(index))
		}
	}
	b.printBoard(results...)
}

func (b board) AllEqual(one, two, three string) bool {
	if one != "" {
		if one == two && one == three {
			return true
		}
	}
	return false
}

func (b board) hasWinner() bool {
	// Check horizontals
	if b.AllEqual(b.squares[0].occupier, b.squares[1].occupier, b.squares[2].occupier) {
		return true
	}
	if b.AllEqual(b.squares[3].occupier, b.squares[4].occupier, b.squares[5].occupier) {
		return true
	}
	if b.AllEqual(b.squares[6].occupier, b.squares[7].occupier, b.squares[8].occupier) {
		return true
	}
	//Check veritcals
	if b.AllEqual(b.squares[0].occupier, b.squares[3].occupier, b.squares[6].occupier) {
		return true
	}
	if b.AllEqual(b.squares[1].occupier, b.squares[4].occupier, b.squares[7].occupier) {
		return true
	}
	if b.AllEqual(b.squares[2].occupier, b.squares[5].occupier, b.squares[8].occupier) {
		return true
	}
	// Check Diagnols
	if b.AllEqual(b.squares[0].occupier, b.squares[4].occupier, b.squares[8].occupier) {
		return true
	}
	if b.AllEqual(b.squares[2].occupier, b.squares[4].occupier, b.squares[6].occupier) {
		return true
	}
	return false
}

func (b board) isFilled() bool {
	allFilled := true
	for _, square := range b.squares {
		allFilled = allFilled && square.occupier != ""
	}
	return allFilled
}

func (b board) checkWinner() bool {
	winner := b.hasWinner()
	if winner {
		fmt.Println("Congratulations! You've won!")
	}
	return winner
}
func (b board) checkFilled() bool {
	filled := b.isFilled()
	if filled {
		fmt.Println("Womp, stalemate!")
	}
	return filled
}

func initializeBoard(boardSize int) (gameBoard board) {
	test := make([]ticTacBox, boardSize)
	return board{test}
}

func main() {
	fmt.Println("Welcome to TicTacGo!")
	fmt.Println()
	fmt.Println("Initializing new game board...")

	gameBoard := initializeBoard(BoardSize)
	gameBoard.printGameBoard()

	fmt.Println("Please enter the board square number to play.")

	gameStillRunning := true
	playerXsTurn := true
	for gameStillRunning {
		if gameBoard.checkWinner() || gameBoard.checkFilled() {
			fmt.Println("Game over.")
			gameStillRunning = false
		} else {
			playGame(gameBoard, playerXsTurn)
			playerXsTurn = !playerXsTurn
		}

	}
}

func getUserInput(promptText string) (userInput int) {
	fmt.Print(promptText)
	fmt.Scanln(&userInput)
	// Are we supposed to catch err here though?
	return userInput
}

func playGame(gameBoard board, playerXsTurn bool) {
	var player string
	if playerXsTurn {
		player = "X"
	} else {
		player = "O"
	}

	index := getUserInput(fmt.Sprintf("Your turn, Player %s: ", player))
	for index >= BoardSize || index < 0 {
		fmt.Println("Please enter a position equal to 0 and less than", BoardSize)
		index = getUserInput(fmt.Sprintf("Your turn, Player %s: ", player))
	}
	for gameBoard.squares[index].occupier != "" {
		fmt.Println("Position", index, "is already occupied, try another!")
		index = getUserInput(fmt.Sprintf("Your turn, Player %s: ", player))
	}
	gameBoard.squares[index].occupier = player
	gameBoard.printGameBoard()
}
