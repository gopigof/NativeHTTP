package main

import (
	"fmt"
	"os"
	"strconv"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) RegisterRoutes(router *Router) {
	router.Add("GET", "/", h.HomeHandler, true)
	router.Add("GET", "/echo/", h.EchoHandler, false)
	router.Add("GET", "/user-agent", h.UserAgentHandler, true)
	router.Add("GET", "/files", h.FilesHandler, false)
	router.Add("POST", "/files", h.FilesUploadHandler, false)
}

func (h *Handlers) HomeHandler(req *Request) *Response {
	return generateResponse(200, make(map[string]string), make([]byte, 0))
}

func (h *Handlers) EchoHandler(req *Request) *Response {
	response := generateResponse(200, make(map[string]string), make([]byte, 0))
	response.responseBody = []byte(req.uri[6:])
	return response
}

func (h *Handlers) UserAgentHandler(req *Request) *Response {
	response := generateResponse(200, make(map[string]string), make([]byte, 0))
	response.responseBody = []byte(req.headers["User-Agent"])
	return response
}

func (h *Handlers) FilesHandler(req *Request) *Response {
	directory := os.Args[2]
	filename := req.uri[7:]
	file, err := readFile(directory, filename)
	var response *Response

	if err != nil {
		response = generateResponse(404, make(map[string]string), []byte("Not Found"))
	} else {
		response = generateResponse(200, map[string]string{"Content-Type": "application/octet-stream"}, file)
		fmt.Println(response.headers)
	}
	return response
}

func (h *Handlers) FilesUploadHandler(req *Request) *Response {
	if req.headers["Content-Type"] != "application/octet-stream" {
		return generateResponse(400, make(map[string]string), []byte("Bad Request"))
	}

	directory := os.Args[2]
	filename := req.uri[7:]
	contentLength, _ := strconv.Atoi(req.headers["Content-Length"])

	err := writeFile(directory, filename, req.requestBody[:contentLength])
	if err != nil {
		return generateResponse(400, make(map[string]string), []byte("Bad Request"))
	}
	return generateResponse(201, make(map[string]string), make([]byte, 0))
}
