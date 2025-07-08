package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printBoard(board [][]string) {
	for _, value := range board {
		fmt.Println(value)
	}
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func insertCharacter(board [][]string, row int, col int, value string) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
		fmt.Println("Invalid position!")
		return false
	}
	if board[row][col] != "-" {
		fmt.Println("There is already a character at this place!")
		return false
	}
	board[row][col] = value
	return true
}

func getNumOfCharactersInBoard(board [][]string) int {
	total := 0
	for _, row := range board {
		for _, value := range row {
			if value == "X" || value == "O" {
				total++
			}
		}
	}
	return total
}

func promptOptions(board [][]string, currentChar *string, reader *bufio.Reader) {
	if *currentChar == "-" {
		char, _ := getInput("Insert character (X,O): ", reader)
		char = strings.ToUpper(char)
		if char != "X" && char != "O" {
			fmt.Println("Invalid character! Insert X or O.")
			promptOptions(board, currentChar, reader)
			return
		}
		*currentChar = char
	}

	fmt.Println("--------------------------")
	fmt.Printf("It's %s's turn!\n", *currentChar)

	row := getValidCoordinate("row", reader)
	col := getValidCoordinate("column", reader)

	if !insertCharacter(board, row-1, col-1, *currentChar) {
		promptOptions(board, currentChar, reader)
		return
	}

	// Switch player
	if *currentChar == "X" {
		*currentChar = "O"
	} else {
		*currentChar = "X"
	}
}

func getValidCoordinate(label string, reader *bufio.Reader) int {
	for {
		input, _ := getInput(fmt.Sprintf("Insert %s (1,2,3): ", label), reader)
		num, err := strconv.Atoi(input)
		if err == nil && num >= 1 && num <= 3 {
			return num
		}
		fmt.Printf("%s must be an integer between 1 and 3!\n", strings.Title(label))
	}
}

func isGameOver(board [][]string) (bool, string) {
	winnerBool, winner := isThereWinner(board)
	if winnerBool {
		return true, winner
	}
	if getNumOfCharactersInBoard(board) == 9 {
		return true, "draw"
	}
	return false, "-"
}

func isThereWinner(board [][]string) (bool, string) {
	// Rows
	for i := 0; i < 3; i++ {
		if board[i][0] != "-" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true, board[i][0]
		}
	}
	// Columns
	for i := 0; i < 3; i++ {
		if board[0][i] != "-" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true, board[0][i]
		}
	}
	// Diagonals
	if board[0][0] != "-" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true, board[0][0]
	}
	if board[0][2] != "-" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true, board[0][2]
	}
	return false, "-"
}

func playGame() {
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	currentChar := "-"
	reader := bufio.NewReader(os.Stdin)

	for {
		printBoard(board)
		over, winner := isGameOver(board)
		if over {
			if winner == "draw" {
				fmt.Println("The game is a draw!")
			} else {
				fmt.Println("The winner is:", winner)
			}
			break
		}
		promptOptions(board, &currentChar, reader)
	}

	fmt.Println("The game is over!")
}

func main() {
	playGame()
}
