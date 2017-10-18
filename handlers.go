package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		logger.Error("Error decoding user ", err)

	}
	ctx := r.Context()
	db := ctx.Value("db").(*mgo.Database)
	u.create(db)
	return
}
