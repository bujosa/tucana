package app

import (
	"encoding/json"
	"net/http"
	"saturn-golang/service"
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