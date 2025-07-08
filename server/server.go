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

func main() {
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	// 1. Poslušamo na določenem naslovu in portu
	listener, err := net.Listen("tcp", ":9999") // zažene se TCP strežnik na portu 9999
	if err != nil {
		fmt.Println("Napaka pri poslušanju:", err)
		return
	}
	defer listener.Close() // port se zapre, da se lahko naslednjič zažene

	fmt.Println("Strežnik teče na portu 9999. Čakam igralce...")

	// 2. Sprejmemo prvega igralca
	conn1, err := listener.Accept() // čakanje da se en poveže
	if err != nil {
		fmt.Println("Napaka pri sprejemu igralca 1", err)
		return
	}
	fmt.Println("Igralec 1 je povezan!")

	// 3. Sprejmemo drugega igralca
	conn2, err := listener.Accept() // čakanje da se drug poveže
	if err != nil {
		fmt.Println("Napaka pri sprejemu igralca 2", err)
		return
	}
	fmt.Println("Igralec 2 je povezan!")

	// Obvestimo oba igralca
	conn1.Write([]byte("Povezan si kot Igralec 1\n"))
	conn2.Write([]byte("Povezan si kot Igralec 2\n"))

	currentPlayer := conn1
	otherPlayer := conn2
	currentName := "Igralec 1"
	currentSymbol := "X"

	for {
		boardStr := tictactoe.BoardToString(board)
		currentPlayer.Write([]byte("\nTrenutna plošča:\n" + boardStr))
		currentPlayer.Write([]byte("\nNa potezi si. Vnesi potezo v obliki 'vrstica,stolpec' (npr. 2,3):\n"))

		move, err := bufio.NewReader(currentPlayer).ReadString('\n')
		if err != nil {
			fmt.Println(currentName, "je zapustil igro.")
			return
		}

		move = strings.TrimSpace(move)
		coords := strings.Split(move, ",")
		if len(coords) != 2 {
			currentPlayer.Write([]byte("\nNapačen format! Uporabi obliko 'vrstica,stolpec' (npr. 2,3)"))
			continue
		}

		row, err1 := strconv.Atoi(coords[0])
		col, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil || row < 1 || row > 3 || col < 1 || col > 3 {
			currentPlayer.Write([]byte("\nNapačne koordinate! Vnesi števila od 1 do 3.\n"))
			continue
		}

		if board[row-1][col-1] != "-" {
			currentPlayer.Write([]byte("\nTo polje je že zasedeno!\n"))
			continue
		}

		board[row-1][col-1] = currentSymbol

		boardStr = tictactoe.BoardToString(board)
		currentPlayer.Write([]byte("\nPoteza sprejeta.\n"))
		otherPlayer.Write([]byte(currentName + " je naredil potezo: " + move))
		currentPlayer.Write([]byte("\nNova plošča\n" + boardStr))
		otherPlayer.Write([]byte("\nNova plošča\n" + boardStr))

		koncano, zmagovalec := tictactoe.IsGameOver(board)
		if koncano {
			if zmagovalec == "draw" {
				currentPlayer.Write([]byte("\nIgra je neodločena!\n"))
				otherPlayer.Write([]byte("\nIgra je neodločena!\n"))
			} else if zmagovalec == currentSymbol {
				currentPlayer.Write([]byte("\nZmagal si!\n"))
				otherPlayer.Write([]byte("\nIzgubil si! Zmagovalec je " + currentName + "\n"))
			} else {
				currentPlayer.Write([]byte("\nIzgubil si! Zmagovalec je nasprotnik.\n"))
				otherPlayer.Write([]byte("\nZmagal si!\n"))
			}
			break
		}

		if currentPlayer == conn1 {
			currentPlayer = conn2
			otherPlayer = conn1
			currentName = "Igralec 2"
			currentSymbol = "O"
		} else {
			currentPlayer = conn1
			otherPlayer = conn2
			currentName = "Igralec 1"
			currentSymbol = "X"
		}
	}

	conn1.Close()
	conn2.Close()
	fmt.Println("\n\nIgra je končana")
}

func handlePlayer(conn net.Conn, otherConn net.Conn, name string, wg *sync.WaitGroup, board *[][]string) {
	defer wg.Done()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(name, "je zapustil igro")
			return
		}

		message = strings.TrimSpace(message)
		coords := strings.Split(message, ",")
		if len(coords) != 2 {
			conn.Write([]byte("Vpiši koordinati v obliki: vrstica,stolpec (npr. 2,3)\n"))
			continue
		}

		row, err1 := strconv.Atoi(coords[0])
		col, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil || row < 1 || row > 3 || col < 1 || col > 3 {
			conn.Write([]byte("Koordinati morata biti števili med 1 in 3!\n"))
			continue
		}

		if (*board)[row-1][col-1] != "-" {
			conn.Write([]byte("To polje je že zasedeno!\n"))
			continue
		}

		simbol := "X"
		if name == "Igralec 2" {
			simbol = "O"
		}
		(*board)[row-1][col-1] = simbol

		boardStr := tictactoe.BoardToString(*board)
		conn.Write([]byte("Nova plošča:\n" + boardStr + "\n"))
		otherConn.Write([]byte("Nova plošča:\n" + boardStr + "\n"))

		// tu se posreduje sporočilo drugemu igralcu
		_, err = otherConn.Write([]byte(name + ": " + message))
		if err != nil {
			fmt.Println("Napaka pri poslušanju sporočila", err)
			return
		}

	}
}
