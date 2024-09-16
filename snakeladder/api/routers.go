package api

import "github.com/gorilla/mux"
import 	"github.com/prometheus/client_golang/prometheus/promhttp"

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/play", PlayHandler).Methods("GET")
	r.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())


	return r
}
