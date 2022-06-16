package main

import "fmt"

type UserCredential struct {
	ID           int
	UserName     string `gorm:"unique;size:50;not null"`
	UserPassword string `gorm:"size:10;not null"`
	IsBlocked    bool   `gorm:"default:true"`
}

func (c UserCredential) String() string {
	return fmt.Sprintf("Id: %d, UserName: %s, Password: %s, IsBlocked: %v", c.ID, c.UserName, c.UserPassword, c.IsBlocked)
}
