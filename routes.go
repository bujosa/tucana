package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	name string
	method string
	pattern string
	handledFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StaticSlash(true)

	for _, route := range routes{
		route. 
		Name(route.name).		
		Methods(route.method).
		Path(route.pattern).
		Handler(route.handledFunc)
	} 
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index
	},
	Route{
		"Planets",
		"GET",
		"/planets",
		Planets
	},
}