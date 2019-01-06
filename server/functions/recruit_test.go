package functions

import(
  "testing"
  //"github.com/jinzhu/gorm"
)

func TestGetRecruitDB(t *testing.T){
  t.Parallel()
  if testing.Short() {
    t.Skip("Skiping Long Test")
  }

}
