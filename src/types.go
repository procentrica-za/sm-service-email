package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses

type ForgotPasswordEmail struct {
	ToEmail  string `json:"toemail"`
	Subject  string `json:"subject"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

type EmailResult struct {
	Message string `json:"message"`
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
