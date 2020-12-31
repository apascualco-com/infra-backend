package health

import "apascualco.com/api/server/router"

type health struct {
	routes []router.Route
}

func (s *health) Routes() []router.Route {
	return s.routes
}

func New() router.Router {
	r := &health{}
	r.routes = []router.Route{
		router.NewGetRoute("/health", r.getHealth),
	}
	return r
}
