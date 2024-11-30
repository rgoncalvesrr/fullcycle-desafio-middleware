package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rgoncalvesrr/fullcycle-desafio-middleware/application"
	mw "github.com/rgoncalvesrr/fullcycle-desafio-middleware/pkg/middleware"
	"net/http"
)

type ResultError struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.Limiter)

	r.Get("/{cep}", weatherHandler)
	fmt.Println("Server running on port 8080")
	_ = http.ListenAndServe(":8080", r)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(&application.WeatherOutput{
		City:       "São Paulo",
		Celsius:    23.2,
		Kelvin:     73.75,
		Fahrenheit: 296.34,
	})

	return

	// O código abaixo está totalmente funcional e foi comentado para remover os atrasos nos testes

	//
	//cep := r.PathValue("cep")
	//coordinateRepository := adapter.NewCoordinateRepository()
	//weatherRepository := adapter.NewWeatherRepository()
	//w.Header().Set("Content-Type", "application/json")
	//s := application.NewWeatherService(coordinateRepository, weatherRepository)
	//output, e := s.GetTemperature(context.Background(), cep)
	//
	//if e != nil {
	//	switch e {
	//	case application.ErrCepInvalid, application.ErrCepMalformed:
	//		w.WriteHeader(http.StatusUnprocessableEntity)
	//		_ = json.NewEncoder(w).Encode(&ResultError{Message: "invalid zipcode"})
	//	case application.ErrCepNotFound:
	//		w.WriteHeader(http.StatusNotFound)
	//		_ = json.NewEncoder(w).Encode(&ResultError{Message: "can not find zipcode"})
	//	default:
	//		w.WriteHeader(http.StatusInternalServerError)
	//		_ = json.NewEncoder(w).Encode(&ResultError{Message: "internal server error" + e.Error()})
	//	}
	//	return
	//}
	//
	//_ = json.NewEncoder(w).Encode(output)
}
