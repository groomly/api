package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/create-user",
		CreateUserHandler,
	},
	Route{
		"link-trello",
		"POST",
		"/link-trello",
		LinkTrelloHandler,
	},
}
