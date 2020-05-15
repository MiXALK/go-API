package controllers

import (
	"github.com/MiXALK/go-API/internal/models"
	"github.com/MiXALK/go-API/internal/utils"
	"net/http"
)

func (server *APIServer) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	portfolio := &models.Portfolio{}
	utils.ParseBody(r, portfolio)
	portfolioId := portfolio.CreatePortfolio(server.DB)

	response := &Response{
		Status: "success",
		Data: map[string]int64{
			"id": portfolioId,
		},
	}

	server.Respond(w, r, http.StatusCreated, response)
}

func (server *APIServer) GetPortfolios(w http.ResponseWriter, r *http.Request) {
	portfolios := models.GetAllPortfolios(server.DB)

	response := &Response{
		Status: "success",
		Data:   portfolios,
	}
	server.Respond(w, r, http.StatusOK, response)
}

//func GetPortfolioById(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	PortfolioId := vars["PortfolioId"]
//	ID, err:= strconv.ParseInt(PortfolioId, 0, 0)
//	if err != nil {
//		fmt.Println("Error while parsing")
//	}
//	portfolioDetails, _:= models.GetPortfolioById(ID)
//	res, _ := json.Marshal(portfolioDetails)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}
//
//func UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
//	var updatePortfolio = &models.Portfolio{}
//	utils.ParseBody(r, updatePortfolio)
//	vars := mux.Vars(r)
//	portfolioId := vars["portfolioId"]
//	ID, err:= strconv.ParseInt(portfolioId, 0, 0)
//	if err != nil {
//		fmt.Println("Error while parsing")
//	}
//	portfolioDetails, db:= models.GetPortfolioById(ID)
//	if updatePortfolio.Name != "" {
//		portfolioDetails.Name = updatePortfolio.Name
//	}
//
//	db.Save(&portfolioDetails)
//	res, _ := json.Marshal(portfolioDetails)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}
//
//func DeletePortfolio(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	portfolioId := vars["portfolioId"]
//	ID, err:= strconv.ParseInt(portfolioId, 0, 0)
//	if err != nil {
//		fmt.Println("Error while parsing")
//	}
//	portfolio:= models.DeletePortfolio(ID)
//	res, _ := json.Marshal(portfolio)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//}
