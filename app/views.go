package main

import (
	"os"
	"strconv"
)

func HomeHandler(ctx *RequestContext) *Response {
	return ctx.ok(make([]byte, 0))
}

func EchoHandler(ctx *RequestContext) *Response {
	return ctx.ok([]byte(ctx.Request.uri[6:]))
}

func UserAgentHandler(ctx *RequestContext) *Response {
	return ctx.ok([]byte(ctx.Request.headers["User-Agent"]))
}

func FilesHandler(ctx *RequestContext) *Response {
	directory := os.Args[2]
	filename := ctx.Request.uri[7:]
	file, err := readFile(directory, filename)

	if err != nil {
		return ctx.notFound()
	}

	headers := map[string]string{"Content-Type": "application/octet-stream"}
	return ctx.okWithHeaders(headers, file)
}

func FilesUploadHandler(ctx *RequestContext) *Response {
	if ctx.Request.headers["Content-Type"] != "application/octet-stream" {
		return ctx.badRequest("Bad Request")
	}

	directory := os.Args[2]
	filename := ctx.Request.uri[7:]
	contentLength, _ := strconv.Atoi(ctx.Request.headers["Content-Length"])

	err := writeFile(directory, filename, ctx.Request.requestBody[:contentLength])
	if err != nil {
		return ctx.badRequest("Bad Request")
	}
	return ctx.created()
}
