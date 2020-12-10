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

// UpdateUserInfo func(u User) (User, error)
func UpdateUserInfo(u User) (User, error) {
	for key, user := range users {
		if user.ID == u.ID {
			users[key] = &u
			return u, nil
		}
	}
	return User{}, errors.New("Failed to update user information")
}

// RemoveUser func(id int) (string, error)
func RemoveUser(id int) (string, error) {
	for key, user := range users {
		if user.ID == id {
			users = append(users[:key], users[key+1:]...)
			return "Successful delete user from list", nil
		}
	}
	return "", errors.New("Failed to delete user")
}
