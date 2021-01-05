package app

import (
	"log"
	"net/http"
	"saturn-golang/domain"
	"saturn-golang/service"

	"github.com/gorilla/mux"
)

func Start() {
	// mux
	router:= mux.NewRouter()

	// wiring
	ch:= PlanetHandlers{ service: service.NewPlanetService(domain.NewPlanetRepositoryStub())}

	// define routes
	router.HandleFunc("/planets", ch.getAllPlanets).Methods(http.MethodGet)

	//  stating server
	log.Fatal(http.ListenAndServe( "localhost:8000", router))
}