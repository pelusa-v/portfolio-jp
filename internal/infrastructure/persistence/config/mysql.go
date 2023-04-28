package config

import (
	"log"

	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlDB() (*gorm.DB, error) {
	dbString := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	MigrateMysqlDB(db)

	return db, err
}

func MigrateMysqlDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Tag{},
		// &metricsmodels.Card{},
		// &userentities.User{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
