package database

/* Database configuration (postgre)*/
type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	DB_SSL_MODE string
}
