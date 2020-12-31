package pod

import (
	"apascualco.com/api"
	"apascualco.com/pkg/k8s"
	"net/http"
)

func (s *namespace) getNamespaceList(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	client := k8s.New(s.masterUrl, s.config)
	nodes := client.Namespaces()
	return api.WriteJSON(w, http.StatusOK, nodes)
}
