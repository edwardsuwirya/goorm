package main

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BillRepository interface {
	Create(bill *Bill) error
	FindById(id string) (Bill, error)
	BaseRepositoryAdvQuery
}

type billRepository struct {
	db *gorm.DB
}

func (b *billRepository) Create(bill *Bill) error {
	result := b.db.Create(bill)
	return result.Error
}

func (b *billRepository) FindById(id string) (Bill, error) {
	var bill Bill
	result := b.db.First(&bill, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billRepository) FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error) {
	var bill Bill
	result := b.db.Preload(preload).Where(by).First(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billRepository) FindFirstAllPreload(by map[string]interface{}) (interface{}, error) {
	var bill Bill
	result := b.db.Preload(clause.Associations).Where(by).First(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func NewBillRepository(db *gorm.DB) BillRepository {
	repo := new(billRepository)
	repo.db = db
	return repo
}
