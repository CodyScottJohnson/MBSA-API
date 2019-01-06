package functions

import (
	"encoding/json"
	"net/http"
  "log"
	"fmt"

  "github.com/jinzhu/gorm"
  "github.com/gorilla/mux"

	"Hush_API/models"
)

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		User := models.Shouter{}
		_ = json.NewDecoder(r.Body).Decode(&User)

    var count int
    if  db.Model(&models.Login{}).Where("username = ?", User.Login.Username).Count(&count); count > 0{
      w.Header().Set("Content-Type", "application/json")
  		w.Write([]byte(`{"Message":"User Already Exists"}`))
      return
    }
		User.Login.SetPassword(User.Login.Password)
		if err := db.Create(&User).Error; err != nil {
			panic(err)
		}
    User.Login.Password = ""
		payload, err := json.Marshal(User)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
func GetUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%+v\n", r.Context().Value("Login"))
    vars := mux.Vars(r)
    recruit := models.Recruit{}
		db.Preload("Person").Preload("Person.Emails").First(&recruit,vars["id"])
		payload, err := json.Marshal(recruit)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
