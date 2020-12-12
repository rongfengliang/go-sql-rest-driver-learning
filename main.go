package main

import (
	"database/sql"
	"log"

	_ "github.com/adaptant-labs/go-sql-rest-driver"
)

func main() {
	db, err := sql.Open("restsql", "http://localhost:8090/query/v1")
	if err != nil {
		log.Fatalln("some wrong:", err.Error())
	}
	rows, err := db.Query("")
	for rows.Next() {
		var info map[string]interface{}
		rows.Scan(&info)
		log.Println(info)
	}
}
