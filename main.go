package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/groomly/api/trelloClient"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"os"
)

var logger *log.Logger
var JWTMiddleware *jwtmiddleware.JWTMiddleware

type Response struct {
	Message string `json:"message"`
}

func init() {
	JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("AUTH0_CLIENfT_SECRET")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			log.Error("Authentication err: ", err)
			http.Error(w, err, http.StatusUnauthorized)
		},
	})
}

func main() {
	trelloMember, trelloBoards := trelloClient.LinkTrello("79f9002817946dcc922e9c6a629b1763f45d62dc029f340b7dffd6ab93545336")
	log.Info(trelloMember)
	for _, board := range trelloBoards {
		log.Info(board.Lists)
	}
	router := NewRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	n.Use(negroni.HandlerFunc(JWTMiddleware.HandlerWithNext))
	n.Use(negroni.HandlerFunc(DatabaseMiddleware))
	n.Run(":8080")

}
