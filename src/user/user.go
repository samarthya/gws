package user

import "encoding/json"

//User Stores the information about a user
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"firstName"`
	Middlename string `json:"middleName,omitempty"`
	Surname    string `json:"lastName"`
	Email      string `json:"userName"`
	Address    string `json:"address"`
}

//JSONUser marshal the User to a json structure
func JSONUser(user User) []byte {
	b, err := json.Marshal(user)
	if err != nil {
		// Error while unmarshaling
		return []byte("")
	}
	return b
}
