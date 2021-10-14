package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	Host     = "127.0.0.1:8080"
	Protocol = "tcp"
)

func main() {
	fmt.Println("Connecting to", Protocol, "server", Host)
	conn, err := net.Dial(Protocol, Host)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Type something...\n[]")

		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Print("Sent message: " + message)
	}
}
