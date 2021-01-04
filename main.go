package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/planets",Planets)
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}


