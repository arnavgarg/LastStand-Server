package main

import (
	"fmt"
	"os"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("[  ERROR  ] Failed to Listen on Port 8080")
		os.Exit(3)
	}
	defer ln.Close()
	go startGame()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[ ERROR ] Failed to Accept Incoming Connections")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}

func startGame() {

}
