package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json: "date_created"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Avatar      string `json: "avatar"`
	Status      string `json:"status"`
	DateCreated string `json: "date_created"`
}

//Marshaller and return a sungle user
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			Status:      user.Status,
			DateCreated: user.DateCreated,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

// Marshaller and return multi users (slide of users)
func (users Users) Marshall(isPublic bool) []interface{} {

	results := make([]interface{}, len(users))
	for index, user := range users {
		results[index] = user.Marshall(isPublic)
	}
	return results
}
