package pod

import (
	"apascualco.com/api/server/router"
)

type pod struct {
	routes    []router.Route
	masterUrl string
	config    string
}

func (s *pod) Routes() []router.Route {
	return s.routes
}

func NewPod(masterUrl string, config string) router.Router {
	r := &pod{
		masterUrl: masterUrl,
		config:    config,
	}
	r.routes = []router.Route{
		router.NewGetRoute("/pod/lists/{namespace}", r.getPodList),
		router.NewGetRoute("/pod/lists", r.getPodList),
	}
	return r
}
