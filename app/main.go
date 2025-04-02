package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	parsedRequest := parseRequest(reader)

	switch {
	case parsedRequest.uri == "/":
		{
			response := generateResponse(200, make(map[string]string), make([]byte, 0))
			sendResponse(conn, response)
		}
	case strings.HasPrefix(parsedRequest.uri, "/echo"):
		{
			response := generateResponse(200, make(map[string]string), make([]byte, 0))
			response.responseBody = []byte(parsedRequest.uri[6:])
			sendResponse(conn, response)
		}
	case strings.HasPrefix(parsedRequest.uri, "/user-agent"):
		{
			response := generateResponse(200, make(map[string]string), make([]byte, 0))
			response.responseBody = []byte(parsedRequest.headers["User-Agent"])
			sendResponse(conn, response)
		}
	case strings.HasPrefix(parsedRequest.uri, "/files"):
		{
			directory := os.Args[2]
			filename := parsedRequest.uri[7:]
			file, err := readFile(directory, filename)
			var response *Response

			if err != nil {
				response = generateResponse(404, make(map[string]string), []byte("Not Found"))
			} else {
				response = generateResponse(200, map[string]string{"Content-Type": "application/octet-stream"}, file)
				fmt.Println(response.headers)
			}
			sendResponse(conn, response)
		}
	default:
		{
			response := generateResponse(404, make(map[string]string), []byte("NOT FOUND!"))
			sendResponse(conn, response)
		}
	}
	fmt.Println("Responded to request: ", parsedRequest.uri)
}

func main() {
	fmt.Println("Server started")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}
		go handleConnection(conn)
	}
}
