package routes

import (
	"net/http"

	c "gostore/controllers"
)

func GetRoutes() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.NewRegister)
	http.HandleFunc("/insert", c.Insert)
	http.HandleFunc("/delete", c.Delete)
	http.HandleFunc("/edit", c.Edit)
}
