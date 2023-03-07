package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=tiny.db.elephantsql.com user=qappdlgd password=tngkllDXy1xERKOKezhgDbM_MJ5EGGHJ dbname=qappdlgd port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to db")
	}
}
