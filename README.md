# Tic Tac Toe - Console & Network Multiplayer Game 🎮

Welcome to **Tic Tac Toe**, the classic two-player game implemented in **Go**, playable both **locally** in the console or **remotely over a TCP network**!

---

## 🌐 Overview

This project offers two ways to play:
- **Local Console Mode**: Both players use the same terminal.
- **Networked Multiplayer (TCP Server/Client)**: Two players connect remotely to a server and play turn-by-turn.

---

## 🚀 Features

- 👥 Two-player turn-based gameplay  
- ✅ Input validation (format and cell availability)  
- 🎯 Win/draw detection logic  
- 📡 Networked game mode using TCP sockets  
- 🖥️ Live board rendering after each move

---

## 🧱 Project Structure

- `server.go`: TCP server that manages the game between two clients  
- `client.go`: TCP client that connects to the server to play  
- `tictactoe/`: Common package for game logic (board rendering, win checking, etc.)

---

## 🕸️ Network Mode (TCP Server & Client)

### 🔌 Server

Start the server by running: `go run server.go`
- Listens on localhost:9999
- Waits for two players to connect
- Manages turns, validates moves, and sends board updates

### 💻 Client

Start each player in a separate terminal: `go run client.go`
- Connects to the server on port 9999
- Receives game state and instructions
- Sends moves in the format: row,column (e.g., 2,3)

---

## 🧠 Technologies Used
- Go (Golang)
- TCP socket communication via `net`
- Concurrent connection handling using goroutines
- Console-based user input/output
- Modular game logic (`tictactoe` package)

---

## 📝 How to Play

1. Open a terminal and run the server: go run server.go
2. Open two more terminals, one for each player: go run client.go
3. Players take turns entering moves (e.g., 1,2)
4. The game ends with a win or draw — all players are notified

---

## 💡 Future Improvements
- 🤖 AI opponent (single-player mode)
- 🖼️ GUI interface (e.g., using Ebiten or Fyne)
- 🌐 Web version with WebSockets & frontend (React, Svelte, etc.)
- 📊 Score tracking across games



## Author 🧑‍💻

[Blaž Hribernik](https://github.com/HribernikBlaz)

Feel free to explore, contribute, and enhance Tic Tac Toe with new features like AI opponent, score tracking, or graphical interface!

