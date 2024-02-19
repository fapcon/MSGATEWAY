package router

import (
	"cart/internal/controller"
	"github.com/go-chi/chi"
)

func Route(cnt *controller.HandleUser) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/cart/create", cnt.Create)
	r.Get("/cart/get", cnt.Get)
	return r
}
