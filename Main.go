package main

import (
	"Caro_Game/config"
	"Caro_Game/routers"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		message = strings.TrimSpace(message)

		if message == "quit" {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}

		fmt.Println("Received message:", message)

		// Handle game logic here...
		db := config.ConnectionDatabase()
		routers.InitializeRouters(db)

		response := "OK\n"
		conn.Write([]byte(response))
	}
}

func main() {
	server, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}
	//db := config.ConnectionDatabase()
	//routers.InitializeRouters(db)
}
