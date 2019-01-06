package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type Login struct {
	gorm.Model
	Username       string    `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Password      string    `gorm:"size:255" json:",omitempty"` // Default size for string is 255, reset it with this tag
  OwnerId   int
  OwnerType string
}
func (login *Login) CheckPassword(p string)(error){
	err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(p))
	if(err==nil){
		return nil
	}
	return errors.New("Invalid User or Password")
}
func (login *Login) SetPassword(p string){
	password, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err == nil {
		login.Password = string(password)
		}
}
