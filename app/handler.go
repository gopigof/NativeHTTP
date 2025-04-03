package main

type Handlers struct {
	req *Request
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (ctx *RequestContext) RegisterRoutes(router *Router) {
	router.Add("GET", "/", HomeHandler, true)
	router.Add("GET", "/echo/", EchoHandler, false)
	router.Add("GET", "/user-agent", UserAgentHandler, true)
	router.Add("GET", "/files", FilesHandler, false)
	router.Add("POST", "/files", FilesUploadHandler, false)
}

func (ctx *RequestContext) ok(body []byte) *Response {
	return generateResponse(ctx.Request, 200, make(map[string]string), body)
}

func (ctx *RequestContext) okWithHeaders(headers map[string]string, body []byte) *Response {
	return generateResponse(ctx.Request, 200, headers, body)
}

func (ctx *RequestContext) created() *Response {
	return generateResponse(ctx.Request, 201, make(map[string]string), make([]byte, 0))
}

func (ctx *RequestContext) badRequest(message string) *Response {
	return generateResponse(ctx.Request, 400, make(map[string]string), []byte(message))
}

func (ctx *RequestContext) notFound() *Response {
	return generateResponse(ctx.Request, 404, make(map[string]string), []byte("Not Found"))
}
