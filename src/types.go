package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses
type ForgotPasswordResult struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Message  string `json:"message"`
}
type ForgotPasswordEmail struct {
	ToEmail string `json:"toemail"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

//touter service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	USERMANAGERHost string
	USERMANAGERPort string
	EMAILHost       string
	EMAILPort       string
}
