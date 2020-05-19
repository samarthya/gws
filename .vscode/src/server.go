package main

import (
	"fmt"
	"log"
	"net/http"

	u "github.com/gws/user"
)

// PathAPI API path to be used as base
const PathAPI string = "/api"

// main Entry point for our code.
func main() {
	fmt.Printf(" Webservice.\n")
	//hell.SetupService(PathAPI)
	u.SetupService(PathAPI)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
