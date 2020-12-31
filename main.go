package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/ziyaaktas/sa-backend/gosrc/github"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/popular-repos", github.GetPopularRepos)
	http.ListenAndServe(":3000", r)
}
