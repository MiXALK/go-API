package models

import (
	"database/sql"
	"fmt"
	"github.com/MiXALK/go-API/internal/utils"
)

type Portfolio struct {
	ID   int64
	Name string
}

func (portfolio *Portfolio) CreatePortfolio(db *sql.DB) int64 {
	stmt, e := db.Prepare("INSERT INTO portfolio (name) VALUES (?);")
	utils.ErrorCheck(e)

	res, e := stmt.Exec(portfolio.Name)
	utils.ErrorCheck(e)

	id, e := res.LastInsertId()
	utils.ErrorCheck(e)

	return id
}

func GetAllPortfolios(db *sql.DB) []Portfolio {
	var Portfolios []Portfolio
	rows, e := db.Query("SELECT * FROM `portfolio`;")
	utils.ErrorCheck(e)

	var portfolio = Portfolio{}
	for rows.Next() {
		e = rows.Scan(portfolio.ID, portfolio.Name)
		utils.ErrorCheck(e)
		fmt.Println(portfolio)
	}

	return Portfolios
}

//func GetPortfolioById(Id int64) (*Portfolio , *gorm.DB){
//	var getPortfolio Portfolio
//	db:=db.Where("ID = ?", Id).Find(&getPortfolio)
//	return &getPortfolio, db
//}
//
//func DeletePortfolio(ID int64) Portfolio {
//	var portfolio Portfolio
//	db.Where("ID = ?", ID).Delete(portfolio)
//	return portfolio
//}
