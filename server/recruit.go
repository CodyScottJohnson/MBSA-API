package server

import (
	//"net/http"

	"Hush_API/server/functions"
	"Hush_API/oauth"
)

func IntitializeRecruits() {
	RecruitRouter := router.PathPrefix("/Recruit").Subrouter()
  //RecruitsRouter := router.PathPrefix("/Recruits").Subrouter()

	RecruitRouter.Handle("", functions.EditRecruit(DB)).Methods("POST")
	RecruitRouter.Handle("/{id}", oauth.AuthHandler(functions.GetRecruit(DB))).Methods("GET")
}
