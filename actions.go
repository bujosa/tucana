package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var planets = Planets{

}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde el servidor web con GO")
}

func PlanetList(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Pagina de los planetas")
}

func AddPlanet(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var planet_data Planet
	err := decoder.Decode(&planet_data)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	log.Println(planet_data)
	planets = append(planets, planet_data)
}
