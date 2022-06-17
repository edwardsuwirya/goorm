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
	//err = db.AutoMigrate(&Customer{}, &UserCredential{}, &CustomerPrivileges{}, &Bill{}, &BillDetail{}, &FnB{}, &MenuCategory{})
	//if err != nil {
	//	panic(err)
	//}

	//repo := NewCustomerRepository(db)

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
	//customer01, _ := repo.FindFirstWithPreload(map[string]interface{}{"customer_id": 3}, "UserCredential")
	//c := customer01.(Customer)
	//c.UserCredential.UserPassword = "AbCdEf"
	//err = repo.UpdateBy(&c)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//customer01, _ = repo.FindFirstWithPreload(map[string]interface{}{"customer_id": 3}, "UserCredential")
	//fmt.Println(customer01)

	// Has One
	//01. Create Customer With Credential
	//customer01 := Customer{
	//	MobilePhoneNo: "0856000113",
	//	Name:          "Mamang",
	//	UserCredential: UserCredential{
	//		UserName:     "M4M4n6",
	//		UserPassword: "racing",
	//	},
	//	CustomerPrivileges: CustomerPrivileges{
	//		DiscountPct: new(float64),
	//	},
	//}
	//
	//repo.Create(&customer01)

	//02. Get Customer Info & Poin
	//customer01, _ := repo.FindFirstWithPreload(map[string]interface{}{"customer_id": 31}, "CustomerPrivileges")
	//fmt.Println(customer01)

	//03. Get All Customer Info (User Cred & Priv)
	//customer01, _ := repo.FindFirstAllPreload(map[string]interface{}{"customer_id": 31})
	//fmt.Println(customer01)

	//03. Update discount
	//c := customer01.(Customer)
	//disc := 10.0
	//c.CustomerPrivileges.DiscountPct = &disc
	//err = repo.UpdateBy(&c)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// Has Many
	//repo := NewBillRepository(db)
	//01.Create Bill
	//bill := Bill{
	//	TableNo:   "1",
	//	TransDate: time.Now(),
	//	BillDetail: []BillDetail{
	//		{
	//			MenuName: "Nasi Goreng",
	//			Qty:      1,
	//		},
	//	},
	//}
	//repo.Create(&bill)

	//02. Find Bill
	//bill, _ := repo.FindFirstWithPreload(map[string]interface{}{"id": "4e5f726a-fda3-4845-bd2e-46f4e7c2fc68"}, "BillDetail")
	//fmt.Println(bill)

	//Many To Many
	//01.Register New Menu
	repo := NewFnBRepository(db)
	//fnb01 := FnB{
	//	MenuName: "Nasi Goreng",
	//	Price:    0,
	//	MenuCategories: []MenuCategory{
	//		{
	//			ID: 1,
	//		},
	//		{
	//			ID: 16,
	//		},
	//		{
	//			ID: 13,
	//		},
	//	},
	//}
	//repo.Create(&fnb01)
	fnb01, err := repo.FindFirstWithPreload(map[string]interface{}{"id": 1}, "MenuCategories")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(fnb01)

	//02.Delete Category from a menu
	fnb := fnb01.(FnB)
	err = repo.UpdateAssociation(&fnb, "MenuCategories", []MenuCategory{
		{
			ID: 31111,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fnb)

	//03.Clear Category from a menu
	//err = repo.ClearAssociation(&fnb, "MenuCategories")
	//fmt.Println(fnb)
}
