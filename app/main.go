package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn, router *Router) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	parsedRequest := parseRequest(reader)
	logRequest(parsedRequest)

	response := router.RouteRequests(parsedRequest)
	sendResponse(conn, response)
}

func main() {
	fmt.Println("Server started")

	// Register Routes
	router := &Router{}
	router.Add("GET", "/", HomeHandler, true)
	router.Add("GET", "/echo/", EchoHandler, false)
	router.Add("GET", "/user-agent", UserAgentHandler, true)
	router.Add("GET", "/files", FilesHandler, false)
	router.Add("POST", "/files", FilesUploadHandler, false)

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
		go handleConnection(conn, router)
	}
}
