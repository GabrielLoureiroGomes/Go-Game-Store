package routes

import (
	"net/http"

	c "gostore/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.NewRegister)
	http.HandleFunc("/insert", c.Insert)
}
