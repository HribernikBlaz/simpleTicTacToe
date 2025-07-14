package main

import (
	"bufio"
	"fmt"
	"net"
	"simpletictactoe/tictactoe"
	"strconv"
	"strings"
	"sync"
)

type Game struct {
	Board         [][]string
	Player1       net.Conn
	Player2       net.Conn
	CurrentPlayer net.Conn
	CurrentSymbol string
	Mutex         sync.Mutex
}

func main() {
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	// 1. Listening on specific port
	listener, err := net.Listen("tcp", ":9999") // TCP server starts at port 9999
	if err != nil {
		fmt.Println("Error while listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("The server is running on port 9999. Waiting for players...")

	// 2. Let's accept the first player.
	conn1, err := listener.Accept() // waiting for one to connect
	if err != nil {
		fmt.Println("Error accepting Player 1", err)
		return
	}
	fmt.Println("Player 1 connected!")

	// 3. We accept another player
	conn2, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting Player 2", err)
		return
	}
	fmt.Println("Player 2 is connected!")

	// We alert both players
	conn1.Write([]byte("You are connected as Player 1\n"))
	conn2.Write([]byte("You are connected as Player 2\n"))

	currentPlayer := conn1
	otherPlayer := conn2
	currentName := "Player 1"
	currentSymbol := "X"

	for {
		boardStr := tictactoe.BoardToString(board)
		currentPlayer.Write([]byte("\nCurrent board:\n" + boardStr))
		currentPlayer.Write([]byte("\nIt's your turn. Enter your move in the form 'row,column' (e.g. 2,3):\n"))

		move, err := bufio.NewReader(currentPlayer).ReadString('\n')
		if err != nil {
			fmt.Println(currentName, "has left the game!")
			return
		}

		move = strings.TrimSpace(move)
		coords := strings.Split(move, ",")
		if len(coords) != 2 {
			currentPlayer.Write([]byte("\nWrong format! Use 'row,column' format (e.g. 2,3)"))
			continue
		}

		row, err1 := strconv.Atoi(coords[0])
		col, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil || row < 1 || row > 3 || col < 1 || col > 3 {
			currentPlayer.Write([]byte("\nWrong coordinates! Enter numbers from 1 to 3.\n"))
			continue
		}

		if board[row-1][col-1] != "-" {
			currentPlayer.Write([]byte("\nThis field is already taken!\n"))
			continue
		}

		board[row-1][col-1] = currentSymbol

		boardStr = tictactoe.BoardToString(board)
		currentPlayer.Write([]byte("\nMove accepted\n"))
		otherPlayer.Write([]byte(currentName + " has made the move: " + move))
		currentPlayer.Write([]byte("\nNew board:\n" + boardStr))
		otherPlayer.Write([]byte("\nNew board:\n" + boardStr))

		koncano, zmagovalec := tictactoe.IsGameOver(board)
		if koncano {
			if zmagovalec == "draw" {
				currentPlayer.Write([]byte("\nThe game is a draw!\n"))
				otherPlayer.Write([]byte("\nThe game is a draw\n"))
			} else if zmagovalec == currentSymbol {
				currentPlayer.Write([]byte("\nYou won!\n"))
				otherPlayer.Write([]byte("\nYou lost! The winner is " + currentName + "\n"))
			} else {
				currentPlayer.Write([]byte("\nYou lost! The winner is opponent!\n"))
				otherPlayer.Write([]byte("\nYou won!\n"))
			}
			break
		}

		if currentPlayer == conn1 {
			currentPlayer = conn2
			otherPlayer = conn1
			currentName = "Player 2"
			currentSymbol = "O"
		} else {
			currentPlayer = conn1
			otherPlayer = conn2
			currentName = "Player 1"
			currentSymbol = "X"
		}
	}

	conn1.Close()
	conn2.Close()
	fmt.Println("\n\nThe game is over!!!")
}
