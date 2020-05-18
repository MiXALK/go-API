package controllers

import (
	"fmt"
	"github.com/MiXALK/go-API/internal/models"
	"github.com/MiXALK/go-API/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		Data: map[string][]models.Portfolio{
			"portfolios": portfolios,
		},
	}
	server.Respond(w, r, http.StatusOK, response)
}

func (server *APIServer) GetPortfolioById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portfolioId := vars["id"]
	ID, err := strconv.ParseInt(portfolioId, 0, 0)
	if err != nil {
		fmt.Println("error_while_parsing")
	}
	portfolioDetails, getError := models.GetPortfolioById(server.DB, ID)
	if getError != nil {
		response := &Response{
			Status: "error",
			Data:   getError.Error(),
		}

		server.Respond(w, r, http.StatusNotFound, response)
	} else {
		response := &Response{
			Status: "success",
			Data: map[string]*models.Portfolio{
				"portfolio": portfolioDetails,
			},
		}

		server.Respond(w, r, http.StatusOK, response)
	}
}

func (server *APIServer) UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
	updatePortfolio := &models.Portfolio{}
	utils.ParseBody(r, updatePortfolio)
	vars := mux.Vars(r)
	portfolioId := vars["id"]
	ID, err := strconv.ParseInt(portfolioId, 0, 0)
	if err != nil {
		fmt.Println("error_while_parsing")
	}

	portfolioDetails, getError := models.GetPortfolioById(server.DB, ID)
	if getError != nil {
		response := &Response{
			Status: "error",
			Data:   getError.Error(),
		}

		server.Respond(w, r, http.StatusNotFound, response)
	} else {
		if updatePortfolio.Name == "" {
			response := &Response{
				Status: "error",
				Data:   "not_correct_name_value",
			}

			server.Respond(w, r, http.StatusBadRequest, response)
		} else {
			portfolioDetails.Name = updatePortfolio.Name
			portfolioDetails.UpdatePortfolio(server.DB)
			response := &Response{
				Status: "success",
				Data: map[string]int64{
					"portfolio": portfolioDetails.ID,
				},
			}

			server.Respond(w, r, http.StatusOK, response)
		}
	}
}

func (server *APIServer) DeletePortfolio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portfolioId := vars["id"]
	ID, err := strconv.ParseInt(portfolioId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	deletedPortfolio, getError := models.GetPortfolioById(server.DB, ID)
	if getError != nil {
		response := &Response{
			Status: "error",
			Data:   getError.Error(),
		}

		server.Respond(w, r, http.StatusNotFound, response)
	} else {
		models.DeletePortfolio(server.DB, ID)

		response := &Response{
			Status: "success",
			Data:   deletedPortfolio.ID,
		}

		server.Respond(w, r, http.StatusOK, response)
	}
}
