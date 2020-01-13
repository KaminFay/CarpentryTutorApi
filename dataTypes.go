package main

// User data from within the user table
type User struct {
	Userid    int    `json:"userid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Roleid    int    `json:"roleid"`
}
