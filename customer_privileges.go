package main

import "fmt"

/*
When creating some data with associations,
if its associations value is not zero-value, those associations will be upserted
*/
type CustomerPrivileges struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	CustomerId  uint
	DiscountPct *float64 `gorm:"default:0"`
}

func (CustomerPrivileges) TableName() string {
	return "wmb.m_customer_privileges"
}
func (c CustomerPrivileges) String() string {
	return fmt.Sprintf("Id: %d, Customer ID: %d, Disc: %v", c.ID, c.CustomerId, c.DiscountPct)
}
