package database

import (
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

/* Create global engine variable */
var PGEngine *gorm.DB // Postgre engine

/* This function will activate connection of selected Engine variable (in this case Engine1 is Postgresql DB Connection). Therefore you will need to call this function before perform operation to database; and do not forget to Close when operation finish! */
func ConnectPGEngine() {

	log.Println("Connecting to postgre database...")

	/* Getting app configuration */
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Println(err)
	}

	/* Create Gorm database engine */
	var err error
	PGEngine, err = gorm.Open("postgres", "host="+config.DB_HOST+" port="+config.DB_PORT+" user="+config.DB_USER+" dbname="+config.DB_NAME+" sslmode="+config.DB_SSL_MODE+" password="+config.DB_PASSWORD+"")
	if err != nil {
		panic(err)
	}

	// PGEngine.LogMode(true)
	/* Returning database engine */
	//return PGEngine

}
