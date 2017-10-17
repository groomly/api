package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Handlerfunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"CreateCard",
		"POST",
		"/create-card",
		CreateCard,
	},
}
