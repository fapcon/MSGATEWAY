package router

import (
	"github.com/go-chi/chi"
	"user/internal/controller"
)

func Route(cnt *controller.HandleUser) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/user/create", cnt.Create)
	r.Get("/user/get", cnt.Get)
	return r
}
