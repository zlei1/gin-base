package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Setup() {
	_config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		viper.GetString("pgsql.host"),
		viper.GetString("pgsql.port"),
		viper.GetString("pgsql.username"),
		viper.GetString("pgsql.database"),
		viper.GetString("pgsql.password"),
	)

	db, err := gorm.Open("postgres", _config)
	if err != nil {
		panic(err.Error())
	}

	dbConfig(db)

	DB = db
}

func dbConfig(db *gorm.DB) {
}
