package main

import (
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"net"
	"os"
)

func db() *mgo.Database {
	url := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInfo, err := mgo.ParseURL(url)

	if err != nil {
		log.Error("Error with db dial ", err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	defer session.Close()

	return session.Clone().DB(dbName)
}
