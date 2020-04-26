package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	passedInID := vars["userid"]
	currentUser := User{}

	sqlStatment := `SELECT user_id, first_name, last_name, username, role_id FROM users WHERE user_id = $1;`
	fmt.Println(passedInID)
	row := db.QueryRow(sqlStatment, passedInID)

	switch err := row.Scan(&currentUser.Userid, &currentUser.Firstname, &currentUser.Lastname,
		&currentUser.Username, &currentUser.Roleid); err {
	case sql.ErrNoRows:
		getUserLog("api", "getUser", "No rows were returned!", WARN)
	case nil:
		fmt.Printf("This is the current User: %+v\n", currentUser)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(currentUser)
		getUserLog("api", "getUser", "JSON Was parsed A-ok!", INFO)

	default:
		getUserLog("api", "getUser", "Something went very wrong, panicking now!", PANIC)
		fmt.Printf("%v\n", err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

	user := User{}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	/* Need to make sure we are incrementing our ID, could probably use postgres to do this on it's own but we'll see*/
	// user.Userid = getHighestUserID()
	// log.Printf("We got this back: %+v\n", user)

	sqlStatement := `INSERT INTO users (user_id, first_name, last_name, username, role_id) VALUES (DEFAULT, $1, $2, $3, $4);`
	_, err = db.Exec(sqlStatement, user.Firstname, user.Lastname, user.Username, user.Roleid)

	if err != nil {
		panic(err)
	} else {
		log.Printf("And it was inserts!\n")
	}
}

/* Deprecated as we just use postgres to auto increment the value of the new users ID.
Leaving this here for now but will probably get rid of it over time */
func getHighestUserID() int {

	var highestID int

	sqlStatement := `SELECT user_id FROM users ORDER BY user_id DESC LIMIT 1;`
	row := db.QueryRow(sqlStatement)

	switch err := row.Scan(&highestID); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(highestID)
	default:
		panic(err)
	}

	return (highestID + 1)
}
