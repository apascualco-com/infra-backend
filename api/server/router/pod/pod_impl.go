package pod

import (
	"apascualco.com/api"
	"apascualco.com/pkg/k8s"
	"net/http"
)

func (s *pod) getPodList(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	client := k8s.New(s.masterUrl, s.config)
	namespace := ""
	if val, ok := vars["namespace"]; ok {
		namespace = val
	}
	nodes := client.Pods(namespace)
	return api.WriteJSON(w, http.StatusOK, nodes)
}
