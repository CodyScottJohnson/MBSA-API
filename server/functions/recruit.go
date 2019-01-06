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

func EditRecruit(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recruit := models.Recruit{}
		_ = json.NewDecoder(r.Body).Decode(&recruit)
		db.Save(&recruit)
		payload, err := json.Marshal(recruit)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
func GetRecruit(db *gorm.DB) http.HandlerFunc {
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
