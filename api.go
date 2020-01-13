package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Fay89058"
	dbname   = "carpentrytutor"
)

var db *sql.DB
var dbError error

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func establishDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, dbError = sql.Open("postgres", psqlInfo)
	if dbError != nil {
		testDatabaseConnectionLog("api.go", "establishDatabase", "Failure to connect to the database", PANIC)
	}

	dbError = db.Ping()
	if dbError != nil {
		testDatabaseConnectionLog("api.go", "establishDatabase", "Could not ping the database!!!!", PANIC)
	}

	testDatabaseConnectionLog("api.go", "establishDatabase", "Conection to the database was successfuly established", INFO)
	return
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is just a test printout")))

	var firstName string
	var lastName string
	var id int

	sqlStatment := `SELECT first_name, last_name, id FROM users WHERE id = '3';`
	row := db.QueryRow(sqlStatment)
	switch err := row.Scan(&firstName, &lastName, &id); err {
	case sql.ErrNoRows:
		w.Write([]byte(fmt.Sprintf("No rows were returned")))
	case nil:
		w.Write([]byte(fmt.Sprintf("FirstName = %s, LastName= %s, id = %d", firstName, lastName, id)))
	default:
		panic(err)
	}
}

type Profile struct {
	Email    string
	Username string
	ID       int
}

func testPOST(w http.ResponseWriter, r *http.Request) {
	// profile := Profile{Email: "abhirockzz@gmail.com", Username: "abhirockzz", ID: 3}
	user := User{Userid: 3, Firstname: "Kamin", Lastname: "Fay", Username: "kfay", Roleid: 4}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&user)
	if err != nil {
		panic(err)
	}
}

func getClass() {

}

func createClass() {

}
