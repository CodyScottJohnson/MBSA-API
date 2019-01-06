package sqldb

import(
  "fmt"

  "github.com/gosuri/uilive"
	"github.com/jinzhu/gorm"
	"github.com/ttacon/chalk"

  	"Hush_API/models"
)

func PopulateInitialData(db *gorm.DB) {
	writer := uilive.New()
	writer.Start()
	fmt.Fprintln(writer, chalk.Yellow.Color(chalk.Bold.TextStyle("() Populating Database")))
	Login := models.Login{
		Username: "Admin",
	}
	Login.SetPassword("Pass")
	//if err = db.First(&Login).Error; err != nil {
	if err := db.Create(&Login).Error; err != nil {
		panic(err)
	}
	fmt.Fprintln(writer, chalk.Blue.Color(chalk.Bold.TextStyle("(*) Database Populated")))

	writer.Stop()
}
