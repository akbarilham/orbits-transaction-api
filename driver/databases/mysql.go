package database

import (
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"log"
)

var MYSQLEngine *gorm.DB // Mysql engine

func ConnectMYSQLEngine() {

	log.Println("Connecting to mysql database...")

	/* Getting app configuration */
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Println(err)
	}

	/* Create Gorm database engine */
	var err error
	MYSQLEngine, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	/* Returning database engine */
	MYSQLEngine.LogMode(true)

}
