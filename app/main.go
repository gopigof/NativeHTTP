package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Server started")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 32*1024) // usual limit of 32KB per request
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return
	}

	request := strings.Split(string(buffer[:n]), "\r\n")
	requestLine := strings.Split(request[0], " ")

	if requestLine[1] == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
			return
		}
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
			return
		}
	}
	fmt.Println("Response sent and server closed!")
}
