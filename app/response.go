package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Response struct {
	// request line
	protocol       string
	statusCode     int
	responsePhrase string

	// headers
	headers Header

	responseBody []byte
}

var reasonPhraseMap = map[int]string{
	200: "OK",
	201: "Created",
	404: "Not Found",
}

func generateResponse(req *Request, responseCode int, headers map[string]string, requestBody []byte) *Response {
	resp := Response{}

	if reasonPhrase, exists := reasonPhraseMap[responseCode]; exists {
		resp.statusCode = responseCode
		resp.responsePhrase = reasonPhrase
		resp.protocol = "HTTP/1.1"

		resp.headers = headers

		// set default Content-Type
		if _, exists := headers["Content-Type"]; !exists {
			resp.headers["Content-Type"] = "text/plain"
		}
		// check Encoding
		if encoding, exists := req.headers["Accept-Encoding"]; exists && strings.TrimSpace(encoding) == "gzip" {
			resp.headers["Content-Encoding"] = encoding
		}
		resp.responseBody = requestBody
		return &resp
	}
	fmt.Println("Unsupported response code")
	return nil
}

func sendResponse(conn net.Conn, resp *Response) {
	writer := bufio.NewWriter(conn)

	fmt.Fprintf(writer, "%s %d %s\r\n", resp.protocol, resp.statusCode, resp.responsePhrase)

	if len(resp.responseBody) > 0 {
		resp.headers["Content-Length"] = strconv.Itoa(len(resp.responseBody))
	}
	for key, value := range resp.headers {
		fmt.Fprintf(writer, "%s: %s\r\n", key, value)
	}
	fmt.Fprintf(writer, "\r\n")

	if len(resp.responseBody) > 0 {
		writer.Write(resp.responseBody)
	}
	writer.Flush()
}
