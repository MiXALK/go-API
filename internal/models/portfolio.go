package models

import (
	"database/sql"
	"errors"
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
		e = rows.Scan(&portfolio.ID, &portfolio.Name)
		utils.ErrorCheck(e)
		Portfolios = append(Portfolios, portfolio)
	}

	return Portfolios
}

func GetPortfolioById(db *sql.DB, Id int64) (*Portfolio, error) {
	portfolio := &Portfolio{}
	err := db.QueryRow(
		"SELECT * FROM `portfolio` WHERE id = ?;", Id,
	).Scan(
		&portfolio.ID,
		&portfolio.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("portfolio_not_found")
		}

		return nil, err
	}

	return portfolio, nil
}

func (portfolio *Portfolio) UpdatePortfolio(db *sql.DB) {
	sqlStatement := `UPDATE portfolio SET name = ? WHERE id = ?;`
	_, err := db.Exec(sqlStatement, &portfolio.Name, &portfolio.ID)
	utils.ErrorCheck(err)
}

func DeletePortfolio(db *sql.DB, ID int64) {
	sqlStatement := `DELETE FROM portfolio WHERE id = ?;`
	_, err := db.Exec(sqlStatement, ID)
	utils.ErrorCheck(err)
}
