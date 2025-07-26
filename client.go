package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	StartClient()
}

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Failed connection to `localhost:8080` %v", err)
		return
	}
	defer conn.Close()

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	if strings.TrimSpace(response) == "OK" {
		fmt.Println("Server response `OK` correctly got!\nClosing connection...\nConnection close!")
	} else {
		fmt.Printf("Unexpected response from server: %q", response)
	}
}
