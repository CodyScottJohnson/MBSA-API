package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gosuri/uilive"
	"github.com/jinzhu/gorm"
	"github.com/mnmtanish/go-graphiql"
	"github.com/ttacon/chalk"

	"Hush_API/graphql"
	"Hush_API/models"
	"Hush_API/oauth"
)

var (
	router = mux.NewRouter().StrictSlash(true)
	DB     *gorm.DB
)

func Start(db *gorm.DB, port string) error {
	writer := uilive.New()
	writer.Start()
	fmt.Fprintln(writer, chalk.Yellow.Color(chalk.Bold.TextStyle("() Starting Server\n")))
	time.Sleep(time.Millisecond * 1)
	go graphql.InititalizeDB(db)
	a := auth{}
	OauthRouter := router.PathPrefix("/Oauth").Subrouter()
	oauth.InitOauth(OauthRouter, a)
	//db.DropTable(&models.Address{})

	//newEmail := models.Address{Address: "1305 south 400 east", City: "Bountiful", State: "Utah", Primary: true, Type: "Work"}
	//p := models.Person{}
	//db.Where("id = ?", 1).First(&p)
	//db.Model(&p).Association("Addresses").Append(newEmail)
	db.AutoMigrate(&models.Login{},&models.Shouter{})
	////////////////
	//Router Set Up
	///////////////
	DB = db
	IntitializeRecruits()
	IntitializeUsers()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Cody")
	})
	router.HandleFunc("/graphiql", graphiql.ServeGraphiQL)
	router.HandleFunc("/graphql", graphql.ServeGraphQL())
	fmt.Fprintln(writer, chalk.Blue.Color(chalk.Bold.TextStyle("(*) Server Running\n")))

	//Finish Function
	writer.Stop()
	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, router)))
	return nil
}

type auth struct {
}

func (l auth) Login(Username string, Password string) (LogedIn models.Login, err error) {
	newLogin := models.Login{}
	fmt.Println(Username)
	dberr := DB.Where("username = ?", Username).First(&newLogin).Error
	if dberr != nil {
		fmt.Println(dberr)
		err = errors.New("Invalid Login")
		return newLogin, err
	}
	return newLogin, nil

}
