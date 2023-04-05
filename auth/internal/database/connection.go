package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connection() (conn *gorm.DB, err error) {
	conn, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=streaming password=strserver dbname=streaming_server port=5430 sslmode=disable TimeZone=America/Sao_Paulo",
	}), &gorm.Config{})

	return conn, err
}




