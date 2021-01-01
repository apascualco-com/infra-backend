package server

import (
	"apascualco.com/api"
	"apascualco.com/api/server/router/namespace"
	"apascualco.com/api/server/router/node"
	"apascualco.com/api/server/router/pod"
	"fmt"
	"log"
	"net/http"
	"os"

	"apascualco.com/api/server/router"
	"apascualco.com/api/server/router/health"
	"github.com/gorilla/mux"
)

func initRoutes() []router.Router {
	masterUrl := getEnvOrDefault("K8S_MASTER_URL", "https://localhost:6443")
	kubeConfig := getEnvOrDefault("K8S_KUBE_CONFIG", "/etc/infra/config")
	fmt.Println("Master url: " + masterUrl)
	fmt.Println("Kube Config path: " + kubeConfig)
	routes := []router.Router{
		health.New(),
		node.New(masterUrl, kubeConfig),
		namespace.NewNamespace(masterUrl, kubeConfig),
		pod.NewPod(masterUrl, kubeConfig),
	}
	return routes
}

func Start() {
	m := mux.NewRouter()
	routers := initRoutes()
	for _, r := range routers {
		for _, ro := range r.Routes() {
			handler := api.BuildHandler(ro.Handler())
			m.Path(ro.Path()).Methods(ro.Method()).Handler(handler)
		}
	}
	log.Fatal(http.ListenAndServe(":80", m))
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
