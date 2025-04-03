package main

import (
	"strings"
)

type RequestHandler func(req *Request) *Response

type RouteEntry struct {
	Method      string
	Path        string
	Handler     RequestHandler
	StrictMatch bool
}

type Router struct {
	routes []RouteEntry
}

func (rtr *Router) Add(method, path string, handler RequestHandler, strictMatch bool) {
	re := RouteEntry{
		Path:        path,
		Method:      method,
		Handler:     handler,
		StrictMatch: strictMatch,
	}
	rtr.routes = append(rtr.routes, re)
}

func (re *RouteEntry) Match(req Request) bool {
	if req.method != re.Method {
		return false
	}
	if re.StrictMatch && req.uri != re.Path {
		return false
	}
	if !re.StrictMatch && !strings.HasPrefix(req.uri, re.Path) {
		return false
	}
	return true
}

func (rtr *Router) RouteRequests(req *Request) *Response {
	for _, route := range rtr.routes {
		match := route.Match(*req)
		if !match {
			continue
		}
		return route.Handler(req)
	}
	return createNotFoundResponse()
}
