package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Header map[string]string

type Request struct {
	// request line
	method   string
	uri      string
	protocol string

	// headers
	headers Header

	// request body
	requestBody []byte
}

func parseRequest(buff *bufio.Reader) (*Request, error) {
	req := new(Request)

	requestLine, err := buff.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request - Request Line: ", err)
		return nil, err
	}
	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		fmt.Println("Invalid Request Line: ", parts)
	}

	req.method = parts[0]
	req.uri = parts[1]
	req.protocol = parts[2]
	req.headers = make(Header)

	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading request - Header: ", err)
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		colonIndex := strings.Index(line, ":")
		if colonIndex > 0 {
			header := strings.TrimSpace(line[:colonIndex])
			value := strings.TrimSpace(line[colonIndex+1:])
			req.headers[header] = value
		}
	}
	// Populate default headers
	if _, exists := req.headers["Connection"]; !exists {
		req.headers["Connection"] = "keep-alive"
	}

	var body []byte
	if contentLength, ok := req.headers["Content-Length"]; ok {
		length, err := strconv.Atoi(contentLength)
		if err == nil && length > 0 {
			body = make([]byte, length)
			_, err = io.ReadFull(buff, body)
			if err != nil {
				fmt.Println("Error reading request - Body: ", err)
				return nil, err
			}
		}
	}
	req.requestBody = body
	return req, nil
}
