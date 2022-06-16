package main

import (
	"fmt"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Bill struct {
	ID         string `gorm:"size:36;primary_key"`
	TableNo    string
	TransDate  time.Time
	BillDetail []BillDetail
}

func (b *Bill) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = guuid.New().String()
	return nil
}
func (Bill) TableName() string {
	return "wmb.t_pos_bill"
}
func (b Bill) String() string {
	return fmt.Sprintf(`
Bill No: %s 
Table No: [%v]
Trans Date: %v
Bill Detail: [%v]
`, b.ID, b.TableNo, b.TransDate, b.BillDetail)
}
