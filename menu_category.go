package main

import "fmt"

type MenuCategory struct {
	ID           uint
	CategoryName string
}

func (MenuCategory) TableName() string {
	return "wmb.m_menu_category"
}
func (fb MenuCategory) String() string {
	return fmt.Sprintf("Id: %d, Category: %s", fb.ID, fb.CategoryName)
}
