package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
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

func generateResponse(req *Request, responseCode int, headers map[string]string, responseBody []byte) *Response {
	resp := Response{}

	if reasonPhrase, exists := reasonPhraseMap[responseCode]; exists {
		resp.statusCode = responseCode
		resp.responsePhrase = reasonPhrase
		resp.protocol = "HTTP/1.1"

		resp.headers = headers

		// TODO: Copy default headers from request to response
		// set default Content-Type
		if _, exists := headers["Content-Type"]; !exists {
			resp.headers["Content-Type"] = "text/plain"
		}
		// set Connection headers
		if connHeader, exists := req.headers["Connection"]; exists {
			resp.headers["Connection"] = connHeader
		}
		// check Encoding
		responseEncodingTypes := filterSupportedEncodingTypes(req.headers["Accept-Encoding"])
		if responseEncodingTypes != "" {
			encodedResponse, err := encode(responseBody, responseEncodingTypes)
			if err != nil {
				fmt.Println("Error encoding response: ", err)
			}
			resp.responseBody = encodedResponse
			resp.headers["Content-Encoding"] = responseEncodingTypes
		} else {
			resp.responseBody = responseBody
		}
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

func encode(responseBody []byte, supportedEncodings string) ([]byte, error) {
	var compressedBuffer bytes.Buffer

	if strings.Contains(supportedEncodings, "gzip") {
		gzipWriter, err := gzip.NewWriterLevel(&compressedBuffer, gzip.BestCompression)
		if err != nil {
			return nil, err
		}
		_, err = gzipWriter.Write(responseBody)
		if err != nil {
			return nil, err
		}
		if err := gzipWriter.Close(); err != nil {
			return nil, err
		}
		return compressedBuffer.Bytes(), nil
	}
	return make([]byte, 0), nil
}
