package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Setup() *gorm.DB {
	_config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		viper.GetString("pgsql.host"),
		viper.GetString("pgsql.port"),
		viper.GetString("pgsql.username"),
		viper.GetString("pgsql.database"),
		viper.GetString("pgsql.password"),
	)

	db, err := gorm.Open("postgres", _config)
	if err != nil {
		log.Fatalf("%s: %v", "Postgresql Open Failed", err)
	}

	log.Println("Postgresql Connect Succeed")

	dbConfig(db)

	DB = db

	return db
}

func dbConfig(db *gorm.DB) {
}
