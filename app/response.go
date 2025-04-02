package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

type Response struct {
	// request line
	protocol       string
	statusCode     int
	responsePhrase string

	// headers
	headers Header

	responseBody string
}

var reasonPhraseMap = map[int]string{
	200: "OK",
	404: "Not Found",
}

func generateResponse(responseCode int, headers map[string]string, requestBody string) *Response {
	resp := Response{}

	if reasonPhrase, exists := reasonPhraseMap[responseCode]; exists {
		resp.statusCode = responseCode
		resp.responsePhrase = reasonPhrase
		resp.protocol = "HTTP/1.1"

		if headers != nil {
			headers = make(Header)
		}
		resp.headers = headers
		resp.headers["Content-Type"] = "text/plain"
		resp.responseBody = requestBody
		return &resp
	}
	fmt.Println("Unsupported response code")
	return nil
}

func sendResponse(conn net.Conn, resp *Response) {
	writer := bufio.NewWriter(conn)

	fmt.Fprintf(writer, "%s %d %s\r\n", resp.protocol, resp.statusCode, resp.responsePhrase)

	if resp.responseBody != "" {
		resp.headers["Content-Length"] = strconv.Itoa(len(resp.responseBody))
	}
	for key, value := range resp.headers {
		fmt.Fprintf(writer, "%s: %s\r\n", key, value)
	}
	fmt.Fprintf(writer, "\r\n")

	if resp.responseBody != "" {
		writer.WriteString(resp.responseBody)
	}
	writer.Flush()
}
