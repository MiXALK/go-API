package routes

import (
	"github.com/MiXALK/go-API/internal/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/portfolio", controllers.CreatePortfolio).Methods("POST")
}
