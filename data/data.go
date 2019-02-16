package data

import (
	"encoding/json"

	"github.com/parkmenow/PMN-api/models"
)

var (
	users models.UserList
)

func Users() *models.UserList  {
	return &users
}

func UserRegistration(reqBody string) string {
	user := models.User{}
	json.Unmarshal([]byte(reqBody), &user)
	users.List = append(users.List, user)
	// To Do: Add the user into the database
	return "New User Added!"
}