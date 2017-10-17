package main

import ()

func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// do some stuff before
	next(rw, r)
	// do some stuff after
}
