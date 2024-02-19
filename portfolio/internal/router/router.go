package router

import (
	"github.com/go-chi/chi"
	"portfolio/internal/controller"
)

func Route(cnt *controller.HandleUser) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/portfolio/create", cnt.Create)
	r.Get("/portfolio/get", cnt.Get)
	return r
}
