package oauth

import (
	"context"
	"fmt"
	"net/http"

	"Hush_API/oauth/storage"
	//"Hush_API/models"

	"github.com/RangelReale/osin"
	"github.com/gorilla/mux"
)

var (
	server *osin.Server
)

func InitOauth(router *mux.Router, a auth) {
	db, err := storage.InitDB("postgresql://postgres@localhost:32791/Oauth?sslmode=disable")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}

	//Add or delete the AllowedAccessType as you want
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.REFRESH_TOKEN, osin.PASSWORD}
	sconfig.AllowGetAccessRequest = true
	sconfig.AllowClientSecretInParams = true
	sconfig.AccessExpiration = 3600
	server = osin.NewServer(sconfig, storage.NewStorage(db))
	client := storage.Client{
		ID:          "testclient",
		Secret:      "testpass",
		RedirectUri: "http://localhost:14000/appauth/code",
	}
	if err = db.First(&client).Error; err != nil {
		if err := db.Create(&client).Error; err != nil {
			panic(err)
		}
	}
	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ir := server.HandleInfoRequest(resp, r); ir != nil {
			newContext := context.WithValue(r.Context(), "Login", ir.AccessData.UserData)
			newRequest := r.WithContext(newContext)
			//fmt.Printf("%+v\n", ir.AccessData)
			fmt.Printf("%+v\n", newRequest.Context().Value("Login"))
			fmt.Printf("%+v\n", newContext.Value("Login"))
		}
		osin.OutputJSON(resp, w, r)
	})
	router.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			switch ar.Type {
			case osin.AUTHORIZATION_CODE:
				ar.Authorized = true
			case osin.REFRESH_TOKEN:
				ar.Authorized = true
			case osin.PASSWORD:
				Login, err := a.Login(ar.Username, ar.Password)
				if err == nil {
					err = Login.CheckPassword(ar.Password)
					if err == nil {
						ar.UserData = Login
						ar.Authorized = true
					}
				}
			case osin.CLIENT_CREDENTIALS:
				ar.Authorized = true
			case osin.ASSERTION:
				if ar.AssertionType == "urn:osin.example.complete" && ar.Assertion == "osin.data" {
					ar.Authorized = true
				}
			}
			server.FinishAccessRequest(resp, r, ar)
		}
		if resp.IsError && resp.InternalError != nil {
			fmt.Printf("ERROR: %s\n", resp.InternalError)
		}
		if !resp.IsError {
			resp.Output["custom_parameter"] = 19923
		}
		osin.OutputJSON(resp, w, r)
	})

}
func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()
		// Our middleware logic goes here...
		if ir := server.HandleInfoRequest(resp, r); ir != nil {
			newContext := context.WithValue(r.Context(), "Login", ir.AccessData.UserData)
			newRequest := r.WithContext(newContext)
			next.ServeHTTP(w, newRequest)
		} else {
			osin.OutputJSON(resp, w, r)
		}
	})
}
