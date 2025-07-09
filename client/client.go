package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. Connecting to server
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Error while connnecting!:", err)
		return
	}
	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("The server has disconected.")
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	// Sends user's input
	input := bufio.NewReader(os.Stdin)
	for {
		text, _ := input.ReadString('\n')
		conn.Write([]byte(text))
	}
}
