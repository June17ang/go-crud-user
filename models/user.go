package models

import (
	"errors"
)

// User model
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers func() []*User
func GetUsers() []*User {
	return users
}

// GetUserByID func(id int) (*User, error)
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, errors.New("User not found")
}

// InsertNewUser func(u User) (*User, error)
func InsertNewUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("Unable to create new user")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
