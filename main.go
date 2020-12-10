package main

import (
	"net/http"

	"github.com/June17ang/go-crud-user/controllers"
)

func main() {
	controllers.routeController()
	http.ListenAndServe(":3000", nil)
}
