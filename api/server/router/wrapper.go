package router

import (
	"apascualco.com/api"
	"net/http"
)

type RouteWrapper func(r Route) Route

type localRoute struct {
	method  string
	path    string
	handler api.APIFunc
}

func (l localRoute) Handler() api.APIFunc {
	return l.handler
}

func (l localRoute) Method() string {
	return l.method
}

func (l localRoute) Path() string {
	return l.path
}

func NewRoute(method, path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	var r Route = localRoute{method, path, handler}
	for _, o := range opts {
		r = o(r)
	}
	return r
}

func NewGetRoute(path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodGet, path, handler, opts...)
}

func NewPostRoute(path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodPost, path, handler, opts...)
}

func NewPutRoute(path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodPut, path, handler, opts...)
}

func NewDeleteRoute(path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodDelete, path, handler, opts...)
}

func NewOptionsRoute(path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodOptions, path, handler, opts...)
}

func NewHeadRoute(path string, handler api.APIFunc, opts ...RouteWrapper) Route {
	return NewRoute(http.MethodHead, path, handler, opts...)
}
