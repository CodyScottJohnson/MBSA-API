package storage

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB(connection string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&Access{},
		&Authorize{},
		&Client{},
	)
	return db, nil
}
