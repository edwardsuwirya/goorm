package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type FnBRepository interface {
	Create(fnbRepository *FnB) error
	FindById(id uint) (FnB, error)
	BaseRepositoryAdvQuery
	BaseRepositoryAssociationUpdate
}

type fnbRepository struct {
	db *gorm.DB
}

func (m *fnbRepository) Create(fnb *FnB) error {
	result := m.db.Omit("MenuCategories").Create(fnb)
	return result.Error
}

func (m *fnbRepository) FindById(id uint) (FnB, error) {
	var fnb FnB
	result := m.db.First(&fnb, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fnb, nil
		} else {
			return fnb, err
		}
	}
	return fnb, nil
}

func (m *fnbRepository) FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error) {
	var fnb FnB
	result := m.db.Preload(preload).Where(by).First(&fnb)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fnb, nil
		} else {
			return fnb, err
		}
	}
	return fnb, nil
}

func (m *fnbRepository) FindFirstAllPreload(by map[string]interface{}) (interface{}, error) {
	panic("implement me")
}

func (m *fnbRepository) UpdateAssociation(assocModel interface{}, assocName string, assocNewValue interface{}) error {
	err := m.db.Omit(assocName).Model(assocModel).Association(assocName).Replace(assocNewValue)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (m *fnbRepository) ClearAssociation(assocModel interface{}, assocName string) error {
	return m.db.Model(assocModel).Association(assocName).Clear()
}

func NewFnBRepository(db *gorm.DB) FnBRepository {
	repo := new(fnbRepository)
	repo.db = db
	return repo
}
