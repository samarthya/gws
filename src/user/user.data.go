package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// FileName filename
const FileName = "users.json"

// usersMap Stores the user information in memory.
var usersMap = struct {
	sync.RWMutex
	u map[int]User // map of users, which can be easily picked with the index
}{u: make(map[int]User)}

//init Initializes the usermap
func init() {
	fmt.Println(" loading users....")
	uM, err := loadUsersMap()

	usersMap.u = uM

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" Loaded %d users...", len(usersMap.u))
}

// loadUsersMap Utility function to load the users from a file.
func loadUsersMap() (map[int]User, error) {
	// os.Stat returns the FileInfo structure describing file.
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		return nil, fmt.Errorf("[%s] does not exist", FileName)
	}
	userList := make([]User, 0)
	file, _ := ioutil.ReadFile(FileName)

	// Unmarshal the data from the JSON file
	err := json.Unmarshal([]byte(file), &userList)

	if err != nil {
		log.Fatal(err)
	}

	userMap := make(map[int]User)

	// for i := 0; i < len(userList); i++ {
	// 	userMap[userList[i].ID] = userList[i]
	// }

	for _, u := range userList {
		userMap[u.ID] = u
	}
	return userMap, nil
}

// getUser Returns a user.
func getUser(i int) *User {
	usersMap.RLock()
	defer usersMap.RUnlock()

	for user, ok := usersMap.u[i]; ok; {
		return &user
	}

	return nil
}

//getUserList returns a user
func getUserList() []User {
	usersMap.RLock()
	defer usersMap.RUnlock()
	users := make([]User, 0)
	for _, value := range usersMap.u {
		users = append(users, value)
	}
	return users
}

//getUserIDs Sort and return the ID's
func getUserIDs() []int {
	usersMap.RLock()
	userIds := []int{}
	for key := range usersMap.u {
		userIds = append(userIds, key)
	}
	usersMap.RUnlock()
	sort.Ints(userIds)
	return userIds
}

//getNextID returns the next available ID
func getNextID() int {
	userIds := getUserIDs()
	return userIds[len(userIds)-1] + 1
}

//addOrUpdateUser handles the POST or PUT request of add or update
func addOrUpdateUser(user User) (int, error) {
	addOrUpdateID := -1
	if user.ID > 0 {
		oldUser := getUser(user.ID)
		// if it exists, replace it, otherwise return error
		if oldUser == nil {
			return 0, fmt.Errorf("User-ID [%d] doesn't exist", user.ID)
		}
		addOrUpdateID = user.ID
	} else {
		addOrUpdateID = getNextID()
		user.ID = addOrUpdateID
	}

	usersMap.Lock()
	usersMap.u[addOrUpdateID] = user
	usersMap.Unlock()
	return addOrUpdateID, nil
}
