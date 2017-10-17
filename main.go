package main

import (
	"github.com/urfave/negroni"
)

func main() {
	router := NewRouter()
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(AuthMiddleware))
	n.UseHandler(router)
	n.Run(":8080")

}
