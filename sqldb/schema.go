package sqldb

import(
  "fmt"

  "github.com/gosuri/uilive"
	"github.com/jinzhu/gorm"
	"github.com/ttacon/chalk"

  "Hush_API/models"
)

func ResetDatabase(db *gorm.DB) {
	writer := uilive.New()
	writer.Start()
	fmt.Fprintln(writer, chalk.Yellow.Color(chalk.Bold.TextStyle("() Reseting Database")))
	db.DropTableIfExists(&models.Shouter{},
                       &models.Login{},
                     )
  db.CreateTable(&models.Shouter{},
                       &models.Login{},
                      )
	fmt.Fprintln(writer, chalk.Blue.Color(chalk.Bold.TextStyle("(*) Database Reset")))

	writer.Stop()
}
