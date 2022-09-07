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

	sqlStatement := "SELECT \"productId\", \"productGender\" FROM products LIMIT $1"

	//
	rows, err := db.Query(sqlStatement, 3)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type Product struct {
		ProductId string
		Gender    int
	}

	for rows.Next() {
		product := Product{}

		err = rows.Scan(&product.ProductId, &product.Gender)
		if err != nil {
			panic(err)
		}

		fmt.Println(product)
	}

	err = rows.Err()
	if err != nil {
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
