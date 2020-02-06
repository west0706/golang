package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(172.17.0.7:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT * FROM mysql WHERE User = root").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
}
