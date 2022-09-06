package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func main() {
	db := initDatabase()
	defer db.Close()
	//

	sqlStatement := "SELECT \"productId\", \"productGender\" FROM products LIMIT 1"
	var productId string
	var gender uint8

	row := db.QueryRow(sqlStatement)
	err := row.Scan(&productId, &gender)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned!")
	case nil:
		fmt.Println(productId, gender)
	default:
		panic(err)
	}
}

func initDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
