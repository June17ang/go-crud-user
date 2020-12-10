package controllers

import "net/http"

func routeController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/user/", *uc)
}
