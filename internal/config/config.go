package config

import (
	"database/sql"
	"fmt"
	"github.com/MiXALK/go-API/internal/utils"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	serverName := "database:3306"
	user := "root"
	password := "root"
	dbName := "go_api"

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		user,
		password,
		serverName,
		dbName)
	database, openErr := sql.Open("mysql", connectionString)
	utils.ErrorCheck(openErr)

	db = database
	pingErr := db.Ping()
	utils.ErrorCheck(pingErr)
	fmt.Println("Ping status: OK")

	//defer db.Close()

	stmt, prepareErr := db.Prepare(
		"CREATE TABLE portfolio(id int NOT NULL AUTO_INCREMENT, name varchar(50), PRIMARY KEY (id));")
	utils.ErrorCheck(prepareErr)

	_, execErr := stmt.Exec()
	if execErr != nil {
		fmt.Println(execErr.Error())
	} else {
		fmt.Printf("Table created successfully..")
	}
}

func GetDB() *sql.DB {
	return db
}
