package controllers

import (
	"net/http"
	"regexp"
)

// userController object
type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServerHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connected!"))
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/user/(\d+)/?`),
	}
}
