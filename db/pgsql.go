package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var GDB *gorm.DB

func OpenDB() {
	_config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("pgsql.host"),
		viper.GetString("pgsql.port"),
		viper.GetString("pgsql.username"),
		viper.GetString("pgsql.password"),
		viper.GetString("pgsql.database"),
	)

	db, err := gorm.Open("postgres", _config)
	if err != nil {
		panic(err.Error())
	}

	GDB = db
}
