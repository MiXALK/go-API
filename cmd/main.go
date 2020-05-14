package main

import (
	"github.com/MiXALK/go-API/internal/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
