package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Auth0Id string `json:"auth_0_id"`
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
	return
}
