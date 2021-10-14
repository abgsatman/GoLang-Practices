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
	Protokol = "tcp"
)

func main() {
	fmt.Println("Starting " + Protokol + " server on " + Host)
	l, err := net.Listen(Protokol, Host)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("New client connected with " + c.RemoteAddr().String() + " IP.")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client has left.")
		conn.Close()
		return
	}

	log.Println("Client sent a message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}
