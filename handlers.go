package main

import (
	"encoding/json"
	"github.com/groomly/api/trelloClient"
	"gopkg.in/mgo.v2"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

func LinkTrelloHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logger.Error("Error decoding user ", err)

	}
	ctx := r.Context()
	db := ctx.Value("db").(*mgo.Database)
	trelloMember, trelloBoards := trelloClient.LinkTrello(u.TrelloToken)
	u.TrelloMember = trelloMember
	u.TrelloBoards = trelloBoards
	u.update(db)
	return
}
