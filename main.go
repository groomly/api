package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	n := negroni.Classic()
	n.Use(AuthMiddleware)
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":8000", n))

}
