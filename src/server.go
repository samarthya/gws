package main

import (
	"fmt"
	"log"
	"net/http"
)

// PathAPI API path to be used as base
const PathAPI string = "/api"

// main Entry point for our code.
func main() {
	fmt.Printf(" Webservice.\n")

	users.SetupService(PathAPI)

	// http.HandleFunc("/users", handleUsers)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
