package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	name string
	method string
	pattern string
	handleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes{
		router. 
				Name(route.name).		
				Methods(route.method).
				Path(route.pattern).
				Handler(route.handleFunc)
	} 

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Planets",
		"GET",
		"/planets",
		PlanetList,
	},
}