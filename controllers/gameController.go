package controlles

import (
	"html/template"
	"net/http"
	"strconv"

	g "gostore/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	games, err := g.GetAllGames()

	if err != nil {
		panic(err.Error())
	}

	temp.ExecuteTemplate(w, "Index", games)

}

func NewRegister(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "NewRegister", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		producer := r.FormValue("producer")
		platform := r.FormValue("platform")
		parentalRating := r.FormValue("parentalRating")
		cooperative := r.FormValue("cooperative")
		rating := r.FormValue("rating")

		parentalRatingConv, err := strconv.Atoi(parentalRating)

		if err != nil {
			panic(err.Error())
		}

		coopConv, err := strconv.ParseBool(cooperative)

		if err != nil {
			panic(err.Error())
		}

		ratingConv, err := strconv.Atoi(rating)

		if err != nil {
			panic(err.Error())
		}

		g.CreateGame(g.Game{Id: primitive.NewObjectID(), Name: name, Producer: producer, Platform: platform, ParentalRating: parentalRatingConv, Cooperative: coopConv, Rating: ratingConv})
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
