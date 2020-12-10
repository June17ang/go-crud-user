package models

import "os/user"

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

func GetUsers() []*User {
	return users
}

func GetUserById(id int) (User, error) {
	for _, user range users {
		if user.ID == id {
			return user, nil
		}
	}
	return [], nil
}

func InsertNewUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
