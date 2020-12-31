package router

import (
	"apascualco.com/api"
)

type Router interface {
	Routes() []Route
}

type Route interface {
	Handler() api.APIFunc
	Method() string
	Path() string
}
