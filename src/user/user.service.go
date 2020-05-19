package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const userPath string = "users"

// SetupService Service end point
func SetupService(basePath string) {
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	usersHandler := http.HandlerFunc(HandleUsers)
	userHandler := http.HandlerFunc(HandleUser)
	http.Handle(fmt.Sprintf("%s/%s", basePath, userPath), usersHandler)
	http.Handle(fmt.Sprintf("%s/%s/", basePath, userPath), userHandler)
}

// HandleUsers Will expose this API to handle user commands
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	// Handlers can handle request with multiple request methods.
	// Every request has a method a simple string
	switch r.Method {
	case http.MethodGet:
		log.Println(" GET: called")

		// Get the list of users.
		users := getUserList()

		//Marshal the userlist
		b, err := json.Marshal(users)
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
		var newUser User
		e, er := ioutil.ReadAll(r.Body)
		if er != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		er = json.Unmarshal(e, &newUser)
		if er != nil {
			// Error unmarsheling the data
			w.WriteHeader(http.StatusBadRequest)
		}

		fmt.Printf(" User : %s", newUser.Email)

		_, err := addOrUpdateUser(newUser)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		w.Write(JSONUser(newUser))

	case http.MethodOptions:
		return

	default:
		log.Printf(" Method: %s\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{msg: method not supported}"))
	}

}

// HandleUser For a single user request
func HandleUser(w http.ResponseWriter, r *http.Request) {
	// URL specifies either the URI being requested (for server
	// requests) or the URL to access (for client requests).
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", userPath))

	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Convert the last element of the slice (should be an int)
	userID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//userID will be used in the switch below based on the HTTP method
	switch r.Method {
	case http.MethodGet:
		user := getUser(userID)
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		j, err := json.Marshal(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if user.ID != userID {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ msg: id supplied in request and in body mismatch}"))
			return
		}
		_, err = addOrUpdateUser(user)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		//deleteUser(userID)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
