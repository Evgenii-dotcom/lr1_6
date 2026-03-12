package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		message = strings.TrimSpace(message)
		fmt.Println("Получено:", message)

		response := "Echo: " + message + "\n"
		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP сервер запущен на порту 8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}