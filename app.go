package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
*/
func main() {
	dbHost := "167.172.69.254"
	dbPort := "5432"
	dbUser := "smm2"
	dbPassword := "SmmBatchTwo"
	dbName := "smm2"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wmb.",
		}})
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
	err = db.AutoMigrate(&Customer{}, &UserCredential{})
	if err != nil {
		panic(err)
	}

	repo := NewCustomerRepository(db)

	//01. Create Customer With Credential
	//customer01 := Customer{
	//	MobilePhoneNo: "0856000111",
	//	Name:          "Dika",
	//	UserCredential: UserCredential{
	//		UserName:     "dika123",
	//		UserPassword: "111222",
	//	},
	//}
	//
	//repo.Create(&customer01)

	//02. Update Existing Customer with credential
	//newCredential := UserCredential{
	//	UserName:     "tikayes",
	//	UserPassword: "888Abc",
	//}
	//customer01, _ := repo.FindById(4)
	//customer01.UserCredential = newCredential
	//err = repo.UpdateBy(&customer01)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(customer01)

	//03. Update Password
	customer01, _ := repo.FindFirstWithPreload(map[string]interface{}{"customer_id": 3}, "UserCredential")
	c := customer01.(Customer)
	c.UserCredential.UserPassword = "AbCdEf"
	err = repo.UpdateBy(&c)
	if err != nil {
		fmt.Println(err)
	}
	customer01, _ = repo.FindFirstWithPreload(map[string]interface{}{"customer_id": 3}, "UserCredential")
	fmt.Println(customer01)

}
