package main

// User data from within the user table
type User struct {
	Userid    int    `json:"userid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Roleid    int    `json:"roleid"`
}

// Database data that get's pulled from vault.
type Database struct {
	DB   string `json:"db" mapstructure:"dbname"`
	Host string `json:"host" mapstructure:"host"`
	Port string    `json:"port" mapstructure:"port"`
	Pass string `json:"pass" mapstructure:"pass"`
	User string `json:"user" mapstructure:"user"`
}
