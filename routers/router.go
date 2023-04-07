package routers

import (
	"github.com/DipayanP007/botkube-teams-proxy/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.DisplayPayload).Methods("POST")
	r.HandleFunc("/", controllers.Home).Methods("GET", "HEAD")
	r.HandleFunc("/healthz", controllers.Health).Methods("GET", "HEAD")
	return r
}
