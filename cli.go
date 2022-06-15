package main

import (
	"enigmacamp.com/goorm/config"
	"enigmacamp.com/goorm/manager"
	"enigmacamp.com/goorm/model"
	"fmt"
)

type Cli struct {
	UseCaseManager manager.UseCaseManager
}

//func (c *Cli) Migration(){
//	err = db.AutoMigrate(&Customer{})
//	if err != nil {
//		panic(err)
//	}
//}

func (c *Cli) Run() {
	// Register Customer
	customer01 := model.Customer{
		MobilePhoneNo: "08129080113",
		Name:          "Samuel",
	}
	err := c.UseCaseManager.CustomerRegistrationUseCase().Register(&customer01)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Success register %v\n", customer01)
	}

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

	//count, err := repo.Count("customer_id")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(count)

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

func Console() *Cli {
	c := config.NewConfig()
	infraManager := manager.NewInfra(c)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	return &Cli{
		UseCaseManager: useCaseManager,
	}
}
