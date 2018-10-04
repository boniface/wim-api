package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"wim-api/api"
)

func main() {
	fmt.Println(" This is a Server")
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Route("/ecudata", func(r chi.Router) {
		r.Post("/simple", api.VehicleDataHandler)
		//	r.Post("/bulk",api.BulkVehicleDataHandler)

	})
	/*	r.Route("/trainngdata", func(r chi.Router) {
			r.Post("/simple",api.TrainingDataHandler)
			r.Post("/bulk",api.BulkTrainingDataHandler)
			r.Get("/{id}",api.TrainingDataHandler)
		})
	*/
	http.ListenAndServe(":3000", r)

}
