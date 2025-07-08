package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. Povezava na strežnik
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Napaka pri povezavi:", err)
		return
	}
	defer conn.Close()

	// Bere sporočila od strežnika v ločeni gorutini
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Strežnik je prekinil povezavo.")
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	// Pošilja uporabnikov vnos
	input := bufio.NewReader(os.Stdin)
	for {
		text, _ := input.ReadString('\n')
		conn.Write([]byte(text))
	}
}
