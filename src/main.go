package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var config Config

//create initialisation functions
func init() {
	config = CreateConfig()
	fmt.Println("Config file has loaded")
	fmt.Printf("UsermanagerHost: %v\n", config.USERMANAGERHost)
	fmt.Printf("UsermanagerPort: %v\n", config.USERMANAGERPort)
	fmt.Printf("EmailHost: %v\n", config.EMAILHost)
	fmt.Printf("EmailPort: %v\n", config.EMAILPort)
}

//create config functions
func CreateConfig() Config {
	conf := Config{
		USERMANAGERHost: os.Getenv("USERMANAGER_Host"),
		USERMANAGERPort: os.Getenv("USERMANAGER_Port"),
		EMAILHost:       os.Getenv("EMAIL_Host"),
		EMAILPort:       os.Getenv("EMAIL_PORT"),
	}
	return conf
}
func main() {

	server := Server{
		router: mux.NewRouter(),
	}
	//Set up routes for server
	server.routes()
	handler := removeTrailingSlash(server.router)
	fmt.Print("starting server on port " + config.EMAILPort + "\n")
	log.Fatal(http.ListenAndServe(":"+config.EMAILPort, handler))

}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

func sendMail(toEmail string, subject string, body string) {
	sender := NewSender("studymoneygradproject@gmail.com", "Studymoney123")

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{toEmail}

	Subject := subject
	message := fmt.Fprintf(`
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>%v<br>
	<div class="moz-signature"><i><br>
	<br>
	Regards<br>
	Study Money<br>
	<i></div>
	</body>
	</html>
	`, body)
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)

	sender.SendMail(Receiver, Subject, bodyMessage)
}
