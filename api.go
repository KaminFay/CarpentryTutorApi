package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "superSecretPassword"
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

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passedInID := vars["username"]
	var currentUser User

	sqlStatment := `SELECT user_id, first_name, last_name, username, role_id FROM users WHERE username = $1;`
	fmt.Println(passedInID)
	row := db.QueryRow(sqlStatment, passedInID)

	switch err := row.Scan(&currentUser.UserID, &currentUser.FirstName, &currentUser.LastName,
		&currentUser.Username, &currentUser.RoleID); err {
	case sql.ErrNoRows:
		getUserLog("api", "getUser", "No rows were returned!", WARN)
	case nil:
		err := json.NewEncoder(w).Encode(currentUser)

		if err != nil {
			getUserLog("api", "getUser", "We were unable to parse the JSON...", WARN)
		} else {
			getUserLog("api", "getUser", "JSON Was parsed A-ok!", INFO)
		}
	default:
		getUserLog("api", "getUser", "Something went very wrong, panicking now!", PANIC)
	}
}

func createUser() {

}

func getClass() {

}

func createClass() {

}
