package app

import (
	"encoding/json"
	"net/http"
	"saturn-golang/service"

	"github.com/gorilla/mux"
)

type Planet struct {
	id   string  
	name string   
}
type PlanetHandlers struct {
	service service.IPlanetService
}

func (ch *PlanetHandlers) getAllPlanets(w http.ResponseWriter, r *http.Request){
    planets, _ := ch.service.GetAllPlanets() 
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planets)
}

func (ch *PlanetHandlers) getPlanet(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	id := vars["planet_id"]
	planet, err := ch.service.GetPlanet(id) 
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(planet)
	}	
}