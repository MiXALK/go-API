package models

import (
	"database/sql"
	"github.com/MiXALK/go-API/internal/config"
	"github.com/MiXALK/go-API/internal/utils"
)

var db *sql.DB

type Portfolio struct {
	ID   int64
	Name string
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (portfolio *Portfolio) CreatePortfolio() int64 {
	stmt, e := db.Prepare("INSERT INTO portfolio (name) VALUES (?)")
	utils.ErrorCheck(e)

	res, e := stmt.Exec(portfolio.Name)
	utils.ErrorCheck(e)

	id, e := res.LastInsertId()
	utils.ErrorCheck(e)

	return id
}
