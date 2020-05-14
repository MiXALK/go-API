package controllers

import (
	"encoding/json"
	"github.com/MiXALK/go-API/internal/models"
	"github.com/MiXALK/go-API/internal/utils"
	"net/http"
)

func CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	portfolio := &models.Portfolio{}
	utils.ParseBody(r, portfolio)
	portfolioId := portfolio.CreatePortfolio()

	res, _ := json.Marshal(portfolioId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(res)
	utils.ErrorCheck(err)
}
