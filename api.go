package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// "runtime"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	// vault "github.com/mch1307/vaultlib"
	// "github.com/mitchellh/mapstructure"
)

const (
	host     = "carpentrytutor_db"
	port     = 5432
	user     = "postgres"
	password = "Fay89058"
	dbname   = "carpentrytutor "
)

var db *sql.DB
var dbError error

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)

	fmt.Printf("ADDR: %s, ROLE %s, SECRET %s, TOKEN %s\n", os.Getenv("VAULT_ADDR"), os.Getenv("VAULT_ROLEID"), os.Getenv("VAULT_SECRETID"), os.Getenv("VAULT_TOKEN"))

	// vcConf := vault.NewConfig()

	// // Create new client
	// fmt.Printf("# goroutines before new cli %v\n", runtime.NumGoroutine())
	// vaultCli, err := vault.NewClient(vcConf)
	// if err != nil {
	// 	fmt.Printf("DAMN: %v\n", vcConf)
	// 	log.Fatal(err)
	// }
	// fmt.Printf("AppRole token: %v\n", vaultCli.GetTokenInfo().ID)
	// fmt.Printf("Client status: %v\n", vaultCli.GetStatus())
	// //Get the Vault secret kv_v1/path/my-secret
	// fmt.Printf("# goroutines before getsecret %v\n", runtime.NumGoroutine())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// kv, err := vaultCli.GetSecret("database/creds")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Actual Value of secret: %v\n", kv)

	// loginCredentials := Database{}
	// mapstructure.Decode(kv.KV, &loginCredentials)

	// fmt.Printf("Values once bound to struct: %v\n", loginCredentials)
	// for k, v := range kv.KV {
	// 	fmt.Printf("Secret: %v --- %v\n", k, v)
	// }
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
	dbError = db.Ping()
	if dbError != nil {
		testDatabaseConnectionLog("api.go", "establishDatabase", "Could not ping the database!!!!", PANIC)
		fmt.Printf("%v\n", dbError)
	}
	return
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("This is just a test printout")))

	dbError := db.Ping()
	if dbError != nil {
		testDatabaseConnectionLog("api.go", "establishDatabase", "Could not ping the database!!!!", PANIC)
	} else {
		testDatabaseConnectionLog("api.go", "establishDatabase", "Good to Go!!!", INFO)
	}
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
