package app

import (
	"net/http"
)

type Planet struct {
	id string
	name string
}
type PlanetHandlers struct {
	service service.IPlanetService
}

func (ch *PlanetHandlers) getAllPlanets(w http.ResponseWriter, r *http.Request){
    planets, _ := ch.service.GetAllPlanets() 
	w.Header().Add(key: "Content-Type", value:"application/json")
	json.NewEncoder(w).Encode(planets)
	return planets
}