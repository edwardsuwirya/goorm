package main

import (
	"fmt"
	"time"
)

type Customer struct {
	CustomerId       uint   `gorm:"primaryKey;autoIncrement"`
	MobilePhoneNo    string `gorm:"unique;size:20;not null"`
	Name             string `gorm:"size:100;not null"`
	ActiveMember     bool   `gorm:"default:false"`
	JoinDate         time.Time
	UserCredentialId int
	UserCredential   UserCredential `gorm:"foreignKey:UserCredentialId"`
}

func (Customer) TableName() string {
	// Apabila menggunakan schema wmb, kalau tidak didefinisikan akan masuk ke schema public
	return "wmb.m_customer"
}
func (c Customer) String() string {
	return fmt.Sprintf("Id: %d, Phone: %s, Name: %s, Active Member: %v, Join Date: %v, Credential: %v", c.CustomerId, c.MobilePhoneNo, c.Name, c.ActiveMember, c.JoinDate, c.UserCredential)
}
