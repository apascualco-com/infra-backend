package node

import (
	"apascualco.com/api/server/router"
)

type node struct {
	routes    []router.Route
	masterUrl string
	config    string
}

func (s *node) Routes() []router.Route {
	return s.routes
}

func New(masterUrl string, config string) router.Router {
	r := &node{
		masterUrl: masterUrl,
		config:    config,
	}
	r.routes = []router.Route{
		router.NewGetRoute("/node/lists", r.getNodeList),
	}
	return r
}
