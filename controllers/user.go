package controllers

import (
	"net/http"
	"regexp"
)

// userController object
type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connected!"))
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/user/(\d+)/?`),
	}
}
