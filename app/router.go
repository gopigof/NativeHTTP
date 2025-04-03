package main

import (
	"strings"
)

type RequestContext struct {
	Request *Request
}

func NewRequestContext(req *Request) *RequestContext {
	return &RequestContext{
		Request: req,
	}
}

type ContextHandler func(ctx *RequestContext) *Response

type RouteEntry struct {
	Method      string
	Path        string
	Handler     ContextHandler
	StrictMatch bool
}

type Router struct {
	routes []RouteEntry
}

func (rtr *Router) Add(method, path string, handler ContextHandler, strictMatch bool) {
	re := RouteEntry{
		Path:        path,
		Method:      method,
		Handler:     handler,
		StrictMatch: strictMatch,
	}
	rtr.routes = append(rtr.routes, re)
}

func (rtr *Router) RouteRequests(req *Request) *Response {
	ctx := NewRequestContext(req)

	for _, route := range rtr.routes {
		if req.method != route.Method {
			continue
		}

		if route.StrictMatch && req.uri != route.Path {
			continue
		}

		if !route.StrictMatch && !strings.HasPrefix(req.uri, route.Path) {
			continue
		}

		return route.Handler(ctx)
	}

	// Default 404 if no route matches
	return ctx.notFound()
}
