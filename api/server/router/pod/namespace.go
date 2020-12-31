package pod

import (
	"apascualco.com/api/server/router"
)

type namespace struct {
	routes    []router.Route
	masterUrl string
	config    string
}

func (s *namespace) Routes() []router.Route {
	return s.routes
}

func NewNamespace(masterUrl string, config string) router.Router {
	r := &namespace{
		masterUrl: masterUrl,
		config:    config,
	}
	r.routes = []router.Route{
		router.NewGetRoute("/pod/namespace/lists", r.getNamespaceList),
	}
	return r
}
