package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samarthya/counter"
	"github.com/samarthya/user"
)

// PathAPI API path to be used as base
const PathAPI string = "/api"

// main Entry point for our code.
func main() {
	fmt.Printf(" Webservice.\n")

	counter.SetupCountRoute(PathAPI)
	user.SetupService(PathAPI)

	// http.HandleFunc("/users", handleUsers)
	log.Printf(" Listening on 8090....")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
