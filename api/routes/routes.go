package routes

import (
	"covid-19/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)
	api := routes.PathPrefix("/covid-19/").Subrouter()

	api.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	api.HandleFunc("/nation-series", controllers.GetNationTimeSeries).Methods("GET")
	return routes
}
