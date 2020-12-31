package health

import (
	"apascualco.com/api"
	"net/http"

	"apascualco.com/api/type/status"
)

func (s *health) getHealth(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	status := &status.Health{
		Status:      "ok",
		Version:     "0.0.1",
		Description: "Health status",
	}

	return api.WriteJSON(w, http.StatusOK, status)
}
