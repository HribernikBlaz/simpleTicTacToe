package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func printBoard(board [][]string) {
	for _, value := range board {
		fmt.Println(value)
	}
}

func insertCharacter(board [][]string, row int, col int, value string) [][]string {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
		fmt.Println("Neveljavna pozicija!")
		return board
	}

	if board[row][col] != "-" {
		fmt.Println("Na tem mestu je že znak!")
		return board
	}

	board[row][col] = value
	return board
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

func promptOptions(board [][]string, currentChar *string) {
	reader := bufio.NewReader(os.Stdin)

	if *currentChar == "-" {
		char, _ := getInput("Vnesite znak (X,O): ", reader)
		char = strings.ToUpper(char)
		if char != "X" && char != "O" {
			fmt.Println("Napačen znak! Vnesite X ali O.")
			promptOptions(board, currentChar)
			return
		}
		*currentChar = char
	} else {
		fmt.Println("--------------------------")
		if *currentChar == "X" {
			fmt.Println("Na potezi je igralec O")
			*currentChar = "O"
		} else {
			*currentChar = "X"
			fmt.Println("Na potezi je igralec X")
		}
	}

	opt, _ := getInput("Vnesite vrstico (1,2,3): ", reader)
	row, err := strconv.Atoi(opt)
	if err != nil || row < 1 || row > 3 {
		fmt.Println("Vrstica mora biti celo število med 1 in 3!")
		promptOptions(board, currentChar)
		return
	}

	opt, _ = getInput("Vnesite stolpec (1,2,3): ", reader)
	col, err := strconv.Atoi(opt)
	if err != nil || col < 1 || col > 3 {
		fmt.Println("Stolpec mora biti celo število med 1 in 3!")
		promptOptions(board, currentChar)
		return
	}

	insertCharacter(board, row-1, col-1, *currentChar)

}

func isGameOver(board [][]string) (bool, string) {
	winnerBool, winner := isThereWinner(board)
	if getNumOfCharactersInBoard(board) == 9 && !winnerBool {
		return false, "izenačeno"
	}
	if winnerBool {
		return true, winner
	}
	return false, winner
}

func isThereWinner(board [][]string) (bool, string) {
	if board[0][0] != "-" && board[0][0] == board[0][1] && board[0][1] == board[0][2] { // vrstica 1
		return true, board[0][0]
	} else if board[1][0] != "-" && board[1][0] == board[1][1] && board[1][1] == board[1][2] { // vrstica 2
		return true, board[1][0]
	} else if board[2][0] != "-" && board[2][0] == board[2][1] && board[2][1] == board[2][2] { // vrstica 3
		return true, board[2][0]
	} else if board[0][0] != "-" && board[0][0] == board[1][0] && board[1][0] == board[2][0] { // stolpec 1
		return true, board[0][0]
	} else if board[0][1] != "-" && board[0][1] == board[1][1] && board[1][1] == board[1][2] { // stolpec 2
		return true, board[0][1]
	} else if board[0][2] != "-" && board[0][2] == board[1][2] && board[1][2] == board[2][2] { // stolpec 3
		return true, board[0][2]
	} else if board[0][0] != "-" && board[0][0] == board[1][1] && board[1][1] == board[2][2] { // diagonala 1
		return true, board[0][0]
	} else if board[1][2] != "-" && board[1][2] == board[1][1] && board[1][1] == board[2][1] { // diagonala 2
		return true, board[1][2]
	} else {
		return false, "-"
	}
}

func main() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	currentChar := "-"
	for getNumOfCharactersInBoard(board) < 9 {
		printBoard(board)
		over, winner := isGameOver(board)
		if !over {
			promptOptions(board, &currentChar)
		} else {
			fmt.Println("Zmagovalec je: ", winner)
			break
		}
	}

	printBoard(board)
	fmt.Println("Igra je končana!")

}
