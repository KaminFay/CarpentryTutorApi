package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "thecarpentrytutor"
)

func testFunction() {
	fmt.Println("This is just a test print from the api.go function")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is just a test printout")))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var first_name string
	var last_name string
	var id int

	sqlStatment := `SELECT first_name, last_name, id FROM users WHERE id = '3';`
	row := db.QueryRow(sqlStatment)
	switch err := row.Scan(&first_name, &last_name, &id); err {
	case sql.ErrNoRows:
		w.Write([]byte(fmt.Sprintf("No rows were returned")))
	case nil:
		w.Write([]byte(fmt.Sprintf("FirstName = %s, LastName= %s, id = %d", first_name, last_name, id)))
	default:
		panic(err)
	}
}
