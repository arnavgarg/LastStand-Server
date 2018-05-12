package main

import (
	"fmt"
	"os"
	"net"
	"io"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("[  ERROR  ] Failed to Listen on Port 8080\n" + err.Error())
		os.Exit(3)
	}
	defer ln.Close()
	go startGame()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[  ERROR  ] Failed to Accept Incoming Connections\n" + err.Error())
		}
		fmt.Println("[ SUCCESS ] Connection from " + conn.RemoteAddr().String())
		io.WriteString(conn, fmt.Sprint("[ SUCCESS ] Connection established at ", time.Now() , "\n"))

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}

func startGame() {

}
