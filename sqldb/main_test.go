package sqldb

import(
  "testing"
  //"github.com/jinzhu/gorm"
)

func TestCanConnectToDB(t *testing.T){
  t.Parallel()
  if testing.Short() {
    t.Skip("Skiiping Long Test")
  }

  db, err := Connect()

  defer db.Close()
  if err != nil{
      t.Error("Failed to connect to db")
  }
}
