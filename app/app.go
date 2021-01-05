package app

import(

)

func Start() {
	router:= mux.NewRouter()
	log.Fatal(http.ListenAndServe( addr: "localhost:8000", router))
}