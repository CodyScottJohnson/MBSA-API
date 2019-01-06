package sqldb

import (
	"fmt"
	"time"
	//"log"
	//"database/sql"
	"github.com/gosuri/uilive"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ttacon/chalk"

	"Hush_API/models"
)

// Connect Initializes a New Database Connection returns a Gorm Database pointer
func Connect(sqlconfig config) (*gorm.DB, error) {
	time.Sleep(time.Millisecond * 1)
	writer := uilive.New()
	writer.Start()
	fmt.Fprintln(writer, chalk.Yellow.Color(chalk.Bold.TextStyle("() Connecting To Database")))
	time.Sleep(time.Millisecond * 1)
	db, err := gorm.Open("postgres", sqlconfig.GetString("url"))

	if err != nil {
		fmt.Fprintln(writer, chalk.Red.Color("(-)Unable To Connect to Database:"), err)
		panic(chalk.Red.Color(err.Error()))
	}
	fmt.Fprintln(writer, chalk.Blue.Color(chalk.Bold.TextStyle("(*) Connected to Database")))

	writer.Stop()
	return db, nil

}

// SyncDatabases make sure tables are upto date
func SyncDatabases(db *gorm.DB) {
	db.AutoMigrate(&models.Login{})
}

// PopulateInitialData populates some initial databases
