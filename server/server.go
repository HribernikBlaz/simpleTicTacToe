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

	boardStr := tictactoe.BoardToString(board)
	conn1.Write([]byte("Trenutna plošča:\n" + boardStr))
	conn2.Write([]byte("Trenutna plošča:\n" + boardStr))

	var wg sync.WaitGroup
	wg.Add(2)

	// Poženemo komunikacijo za oba igralca
	go handlePlayer(conn1, conn2, "Igralec 1", &wg, &board)
	go handlePlayer(conn2, conn1, "Igralec 2", &wg, &board)

	wg.Wait() // Počakamo, da oba igralca zaključita
	fmt.Println("Oba igralca sta zapustila igro. Strežnik se zapira.")

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

		// simbol določiš glede na igralca (lahko dodaš param)
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
