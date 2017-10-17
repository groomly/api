package main

import (
	"context"
	"net/http"
)

func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// do some stuff before
	next(rw, r)
	// do some stuff after
}

func DatabaseMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	db := db()
	defer db.Session.Close()
	ctx := context.WithValue(r.Context(), "db", db)
	next(rw, r.WithContext(ctx))
}
