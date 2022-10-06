package config

import (
	"aqua-developer-assesment/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("dsn")), &gorm.Config{})

	if err != nil {
		log.Println("Database Connection Failed!")
		return
	}

	log.Println("Database Connection Succeed!")
}

func Migrate() {
	DB.Debug().AutoMigrate(
		&entity.Product{},
		&entity.Cart{},
		&entity.CartDetail{},
	)
}
