package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"saturn-golang/service"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	_id  bson.ObjectId    `json:"id,omitempty"`
	name string           `json:"name,omitempty"`
}

type PlanetHandlers struct {
	service service.IPlanetService
}

func (ch *PlanetHandlers) getAllPlanets(w http.ResponseWriter, r *http.Request){
	planets, _ := ch.service.GetAllPlanets() 
	fmt.Println("resultados: ",planets)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planets)
}

func (ch *PlanetHandlers) getPlanet(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	id := vars["id"]
	fmt.Println("id: ",id)
	oid := bson.ObjectIdHex(id)
	planet, err := ch.service.GetPlanet(oid) 
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(planet)
	}	
}