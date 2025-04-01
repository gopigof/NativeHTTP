package main

import (
	"fmt"
	"net"
	"os"
)

var responseStatusCodes = map[int]string{
	200: "OK",
	404: "Not Found",
}

func sendHttpResponse(conn net.Conn, content string, responseCode int) {
	statusCode, exists := responseStatusCodes[responseCode]
	if !exists {
		fmt.Println("Unknown response code")
		os.Exit(1)
	}

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", responseCode, statusCode)
	headers := fmt.Sprintf("Content-Type: text/plain\r\nContent-Length: %d\r\n\r\n", len(content))
	responseBody := content

	_, err := conn.Write([]byte(statusLine + headers + responseBody))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		return
	}
}
