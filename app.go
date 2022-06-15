package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
*/
func main() {
	dbHost := "167.172.69.254"
	dbPort := "5432"
	dbUser := "smm2"
	dbPassword := "batchTwo"
	dbName := "smm2"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	enigmaDb, err := db.DB()
	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}
	}(enigmaDb)
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	}

}
