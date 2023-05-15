package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	INTERNAL_SERVER_ERROR =[]byte("500:Internal Server Error")
)

func New() http.Handler{
	r := chi.NewRouter()

	pr := &PlayerRouter{}

	r.Mount("/player", pr.Routes())

	return r
}