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
		var info interface{}
		rows.Scan(&info)
		switch info.(type) {
		case []interface{}:
			for _, v := range info.([]interface{}) {
				log.Println("info:", v)
			}
		default:
			log.Println("unkonw")
		}

	}
}
