package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CarRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(TimerTrace)

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", carsHandle)
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", brandHandle)
			r.Route("/{model}", func(r chi.Router) {
				r.Get("/", modelHandle)

			})
		})
	})
	return r
}
