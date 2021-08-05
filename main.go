package main

import (
	"net/http"

	r "gostore/routes"
)

func main() {
	r.GetRoutes()
	http.ListenAndServe(":8080", nil)
}
