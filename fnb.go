package main

import "fmt"

type FnB struct {
	ID             uint
	MenuName       string
	Price          float64
	MenuCategories []MenuCategory `gorm:"many2many:m_fnb_category;"`
}

func (FnB) TableName() string {
	return "wmb.m_fnb"
}
func (fb FnB) String() string {
	return fmt.Sprintf("Id: %d, Menu: %v, Price: %f, Category:%v", fb.ID, fb.MenuName, fb.Price, fb.MenuCategories)
}
