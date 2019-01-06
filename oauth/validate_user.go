package oauth

import (
	//"Hush_API/oauth/storage"
	//"github.com/RangelReale/osin"

  "Hush_API/models"
)

type auth interface {
  Login(Username string,Password string)(i models.Login,err error)
}
