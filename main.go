package main

import (
	"fmt"
	"groupie_tracker/docs/datagather"
	"groupie_tracker/docs/serverhandler"
	"log"
	"net/http"
	"time"
)

/*
Main script to start Groupie-Tracker server on localhost, port 8080
*/
func main() {
	// Initialisation message
	fmt.Println("\n	     Starting server on localhost:8080")

	/*Retrieve data from RESTful-API
	As of 18-11-2022, the API "https://groupietrackers.herokuapp.com/api" consisted of:
	{"artists":"https://groupietrackers.herokuapp.com/api/artists",
	"locations":"https://groupietrackers.herokuapp.com/api/locations",
	"dates":"https://groupietrackers.herokuapp.com/api/dates",
	"relation":"https://groupietrackers.herokuapp.com/api/relation"}
	*/
	errData := datagather.SaveData("https://groupietrackers.herokuapp.com/api/artists")

	// In the event of error in backend
	// Write status codes etc. to be processed by HomeHandler
	if errData != nil {
		serverhandler.D.StatusCode = 500
		serverhandler.D.StatusMsg = "Internal Server Error"
	}

	// Register handlers
	fileServer := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fileServer))
	http.HandleFunc("/", serverhandler.HomeHandler)

	// Incorporate server timeout
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Initialise server, listen / serve on specified port
	fmt.Printf("Thank you for waiting, (ಥ﹏ಥ) Server is now listening on port %v   ᕦ(ò_óˇ)ᕤ...\n", server.Addr)
	errServer := server.ListenAndServe()
	if errServer != nil {
		serverhandler.D.StatusCode = 500
		serverhandler.D.StatusMsg = "Bad Server Request"
		fmt.Println(serverhandler.D.StatusCode, serverhandler.D.StatusMsg)
		log.Fatalln(errServer)

	}

}
