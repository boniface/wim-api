package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"wim-api/api"
	"wim-api/io"
)

func main() {
	fmt.Println(" WIM API Serveice is running on Port 8080")
	go api.MergeAndInsertTraining()
	io.API_Key = api.ApiKeySetter()
	r := chi.NewRouter()
	//go services.ProcessInference()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/{id},{year}", api.GetOutput)
	r.Route("/ecudata", func(r chi.Router) {
		r.Post("/simple", api.SimpleDataHandler)
		r.Post("/bulk", api.BulkVehicleDataHandler)

	})
	r.Route("/infer", func(r chi.Router) {
		r.Get("/", api.InferenceOutputHandler)
		//r.Post("/simple", api.SimpleInferenceIOHanlder)
		//	r.Post("/bulk", api.BulkInferenceIOHanlder)

	})

	http.ListenAndServe(":8080", r)

}
