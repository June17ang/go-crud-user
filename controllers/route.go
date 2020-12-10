package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func routeController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/user/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
