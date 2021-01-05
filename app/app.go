package app

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
)

func Start() {
	// mux
	router:= mux.NewRouter()
	ch:= CustomerHandlers{ service: service.NewPlanetService(domain.NewPlanetRepositoryStub())}
	// define routes
	router.HandleFunc(path: "/planets", ch.getAllPlanets).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe( addr: "localhost:8000", router))
}