package server

import (
	//"net/http"

	"Hush_API/server/functions"
	"Hush_API/oauth"
)

func IntitializeUsers() {
	UserRouter := router.PathPrefix("/User").Subrouter()
  //RecruitsRouter := router.PathPrefix("/Recruits").Subrouter()

	UserRouter.Handle("/Register", functions.RegisterUser(DB)).Methods("POST")
	UserRouter.Handle("/{id}", oauth.AuthHandler(functions.GetRecruit(DB))).Methods("GET")
}
