package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

//User Stores the information about a user
type User struct {
	Name       string `json:"firstName"`
	Middlename string `json:"middleName,omitempty"`
	Surname    string `json:"lastName"`
	Email      string `json:"userName"`
	Address    string `json:"address"`
}

var userList []User

// init Initialized
func init() {
	userList = []User{
		{
			Name:       "Saurabh",
			Middlename: "",
			Surname:    "Sharma",
			Email:      "saurabh@samarthya.me",
			Address:    "Auzzieland",
		},
		{
			Name:       "Gaurav",
			Middlename: "M",
			Surname:    "Sharma",
			Email:      "iam@gaurav.me",
			Address:    "Swaziland",
		},
		{
			Name:       "Bhanuni",
			Middlename: "",
			Surname:    "Sharma",
			Email:      "bhanuni@bhanuni.in",
			Address:    "Papaland",
		},
	}
}

// CountHandler Handles the count
type CountHandler struct {
	sync.Mutex // guards n
	n          int
}

func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

// testMarshalling To showcase the marshalling capability.
func testMarshalling() {
	var user User = User{
		Name:       "Saurabh",
		Middlename: "",
		Surname:    "Sharma",
		Email:      "saurabh@samarthya.me",
		Address:    "Auzzieland",
	}

	var newUser User

	u, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(" User: ", string(u))
	}

	e := json.Unmarshal(u, &newUser)

	if e == nil {
		fmt.Println(" Name : ", newUser.Name, " ", newUser.Surname)
	}

}

// handleUsers Will expose this API to handle user commands
func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Handlers can handle request with multiple request methods.

	// Every request has a method a simple string
	switch r.Method {
	case http.MethodGet:
		log.Println(" GET: called")
		b, err := json.Marshal(userList)
		if err != nil {
			// Error while unmarshaling
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(b)

	case http.MethodPost:
		log.Println(" POST: called")
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("User added"))

	default:
		log.Printf(" Method: %s\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("method not supported"))
		log.Println(" Not supported")
	}

}

// main Entry point for our code.
func main() {
	fmt.Printf(" Webservice.\n")
	http.Handle("/count", new(CountHandler))

	http.HandleFunc("/users", handleUsers)

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
