package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type TableRow struct {
	id         int
	name       string
	age        int
	profession string
}

type Table struct {
	Columns []string
	Data    []TableRow
}

func main() {
	cfg := mysql.Config{
		User:                 "test",
		Passwd:               "example",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "crudder",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error while processing database config", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Could not connect to database", pingErr)
	}

	rows, err := db.Query("select * from test_db")
	defer rows.Close()

	tableData := &Table{}
	columns, err := rows.Columns()

	if err != nil {
		log.Fatal(err)
	}

	tableData.Columns = columns

	for rows.Next() {
		currRow := TableRow{}
		rows.Scan(&currRow.id, &currRow.name, &currRow.age, &currRow.profession)
		tableData.Data = append(tableData.Data, currRow)
	}

	fmt.Println(tableData.Columns)
	for _, test := range tableData.Data {
		fmt.Println(test)
	}
}
