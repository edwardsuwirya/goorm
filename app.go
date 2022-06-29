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
	dbPassword := "5mmB2"
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
	//err = db.AutoMigrate(&Customer{})
	//if err != nil {
	//	panic(err)
	//}

	repo := NewCustomerRepository(db.Debug())

	// Insert customer
	//customer01 := Customer{
	//	MobilePhoneNo: "08788123123",
	//	Name:          "Jution",
	//}
	//err = repo.Create(&customer01)
	//if err != nil {
	//	panic(err)
	//}

	//customer01, err := repo.FindById(11)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(customer01)
	//
	//customers, err := repo.Retrieve()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(customers)

	//customers, err := repo.FindAllBy(map[string]interface{}{"mobile_phone_no": "08788123123"})
	//customers, err := repo.FindFirstBy(map[string]interface{}{"mobile_phone_no": "08788123123"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(customers)

	//customers, err := repo.FindBy("name LIKE ? AND active_member = ?", "%Jut%", false)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//err = repo.Update(&customer, map[string]interface{}{"active_member": true, "name": "Jution Chandra Kirana"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = repo.Delete(1)
	//customer, err := repo.FindFirstBy(map[string]interface{}{"mobile_phone_no": "0878812312123"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(customer)

	var total int64
	err = repo.Count(&total, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)

	var TotalActiveMember []struct {
		Name         string
		ActiveMember bool
		Total        int64
	}
	err = repo.Count(&TotalActiveMember, "active_member")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(TotalActiveMember)

	//var Result []struct {
	//	ActiveMember bool
	//	Total        int64
	//}
	//err = repo.GroupBy(&Result, "active_member,count(active_member) as total", nil, "active_member")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(Result)

	//customerPaging, err := repo.Paging(2,2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(customerPaging)

	//var Result []struct {
	//	Name          string
	//	MobilePhoneNo string
	//}
	//err = repo.Query(&Result, "select name,mobile_phone_no from wmb.m_customer where active_member = ?", true)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(Result)
}
