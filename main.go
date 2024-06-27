package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

type TableRow struct {
	Id         int
	Name       string
	Age        int
	Profession string
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
		rows.Scan(&currRow.Id, &currRow.Name, &currRow.Age, &currRow.Profession)
		tableData.Data = append(tableData.Data, currRow)
	}

	fmt.Println(tableData.Columns)
	for _, test := range tableData.Data {
		fmt.Println(test)
	}

	fsys := os.DirFS("views/static")
	fs := http.FileServerFS(fsys)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/index.html")

		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.Execute(w, tableData)

		if err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("POST /createEntry", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		age, _ := strconv.Atoi(r.PostForm.Get("age"))
		newRow := TableRow{
			Age:        age,
			Name:       r.PostForm.Get("name"),
			Profession: r.PostForm.Get("profession"),
		}

		_, err := db.Exec("INSERT INTO test_db (name, age, profession) VALUES (?, ?, ?)", newRow.Name, newRow.Age, newRow.Profession)
		fmt.Println(err)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
