package main

import (
	"log"
	"net"
)

func main() {
	StartServer()
}

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("The server started working on port `:8080`")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Сonnection setup error: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}

	log.Println(conn.RemoteAddr())
}
