package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn, router *Router) bool {
	err := conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	if err != nil {
		fmt.Println("Connection expired after 2s: ", err)
		return false
	}
	reader := bufio.NewReader(conn)
	parsedRequest, err := parseRequest(reader)
	if err != nil {
		if err == io.EOF {
			return false
		}
	}

	response := router.RouteRequests(parsedRequest)
	sendResponse(conn, response)
	logRequest(parsedRequest)
	return parsedRequest.headers["Connection"] == "keep-alive"
}

func handlePersistentConnection(conn net.Conn, router *Router) {
	defer conn.Close()

	keepAlive := true
	for keepAlive {
		keepAlive = handleConnection(conn, router)
	}
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
		go handlePersistentConnection(conn, router)
	}
}
