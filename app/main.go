package main

import (
	"bufio"
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

	reader := bufio.NewReader(conn)
	parsedRequest := parseRequest(reader)

	switch {
	case parsedRequest.uri == "/":
		{
			response := generateResponse(200, make(map[string]string), "")
			sendResponse(conn, response)
		}
	case strings.HasPrefix(parsedRequest.uri, "/echo"):
		{
			response := generateResponse(200, make(map[string]string), "")
			response.responseBody = parsedRequest.uri[6:]
			sendResponse(conn, response)
		}
	case strings.HasPrefix(parsedRequest.uri, "/user-agent"):
		{
			response := generateResponse(200, make(map[string]string), "")
			response.responseBody = parsedRequest.headers["User-Agent"]
			sendResponse(conn, response)
		}
	default:
		{
			response := generateResponse(404, make(map[string]string), "NOT FOUND!")
			sendResponse(conn, response)
		}
	}

	//n, err := conn.Read(buffer)
	//if err != nil {
	//	fmt.Println("Error reading request: ", err.Error())
	//	return
	//}
	//
	//request := strings.Split(string(buffer[:n]), "\r\n")
	//statusLine := strings.Split(request[0], " ")
	//
	//switch {
	//case strings.HasPrefix(statusLine[1], "/echo"):
	//	sendHttpResponse(conn, statusLine[1][6:], 200)
	//case statusLine[1] == "/user-agent":
	//case statusLine[1] == "/":
	//	sendHttpResponse(conn, "", 200)
	//default:
	//	sendHttpResponse(conn, "", 404)
	//}

	fmt.Println("Response sent and server closed!")
}
