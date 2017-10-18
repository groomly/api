package main

import (
	"github.com/VojtechVitek/go-trello"
	"github.com/groomly/api/trelloClient"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Auth0Id      string                     `json:"auth_0_id"`
	TrelloToken  string                     `json:"trello_id"`
	TrelloMember trello.Member              `json:"trello_member"`
	TrelloBoards []trelloClient.TrelloBoard `json:"trello_boards"`
}

func (u *User) collection() string {
	return "users"
}

func (u *User) create(db *mgo.Database) {
	db.C(u.collection()).Insert(*u)

}

func (u *User) read(db *mgo.Database) {
	db.C(u.collection()).Find(bson.M{"auth_0_id": u.Auth0Id}).One(&u)

}

func (u *User) update(db *mgo.Database) {
	db.C(u.collection()).Insert(*u)
}
