package main

import (
	"fmt"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type BillDetail struct {
	ID       string `gorm:"size:36;primary_key"`
	BillID   string
	MenuName string
	Qty      int
}

func (bd *BillDetail) BeforeCreate(tx *gorm.DB) (err error) {
	bd.ID = guuid.New().String()
	return nil
}
func (BillDetail) TableName() string {
	return "wmb.t_pos_bill_detail"
}

func (bd BillDetail) String() string {
	return fmt.Sprintf("Id: %s,Bill ID: %s, Menu: %v, Qty: %d", bd.ID, bd.BillID, bd.MenuName, bd.Qty)
}
