package initializer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to DB")
	}
}

func Migrate(models ...interface{}) {
	for _, model := range models {
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Fatalf("failed to migrate model %v: %v", model, err)
		}
	}
}
