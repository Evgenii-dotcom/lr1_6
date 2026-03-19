package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type TCPServer struct {
	address string
}

func NewTCPServer(address string) *TCPServer {
	return &TCPServer{address: address}
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("TCP сервер запущен на %s\n", s.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка подключения:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	ProcessSingleMessage(conn)
}


func ProcessSingleMessage(rw io.ReadWriter) error {
	reader := bufio.NewReader(rw)

	message, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	message = strings.TrimSpace(message)

	response := buildResponse(message)
	_, err = rw.Write([]byte(response))

	return err
}


func buildResponse(message string) string {
	return "Echo: " + message + "\n"
}

func main() {
	server := NewTCPServer(":8081")

	if err := server.Start(); err != nil {
		log.Fatalf("Ошибка запуска: %v", err)
	}
}