package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/MiXALK/go-API/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	DB     *sql.DB
	Router *mux.Router
}

type Response struct {
	Status string
	Data   interface{}
}

func (server *APIServer) Initialize(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	database, openErr := sql.Open(dbDriver, connectionString)
	utils.ErrorCheck(openErr)

	server.DB = database
	pingErr := server.DB.Ping()
	utils.ErrorCheck(pingErr)
	fmt.Println("Ping status: OK")
	server.createRequiredDbTables()

	server.Router = mux.NewRouter()
	server.registerRoutes()
}

func (server *APIServer) Run(addr string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", addr), server.Router))
}

func (server *APIServer) registerRoutes() {
	server.Router.HandleFunc("/portfolio", server.CreatePortfolio).Methods("POST")
	server.Router.HandleFunc("/portfolio", server.GetPortfolios).Methods("GET")
	server.Router.HandleFunc("/portfolio/{id}", server.GetPortfolioById).Methods("GET")
	server.Router.HandleFunc("/portfolio/{id}", server.UpdatePortfolio).Methods("PUT")
	server.Router.HandleFunc("/portfolio/{id}", server.DeletePortfolio).Methods("DELETE")
}

func (server *APIServer) createRequiredDbTables() {
	stmt, prepareErr := server.DB.Prepare(
		"CREATE TABLE IF NOT EXISTS `portfolio` (id int NOT NULL AUTO_INCREMENT, name varchar(50), PRIMARY KEY (id));")
	utils.ErrorCheck(prepareErr)

	_, execErr := stmt.Exec()
	utils.ErrorCheck(execErr)
	fmt.Println("Tables created successfully.")
}

func (server *APIServer) Respond(w http.ResponseWriter, r *http.Request, code int, data *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		res, _ := json.Marshal(*data)
		_, err := w.Write(res)
		utils.ErrorCheck(err)
	}
}
