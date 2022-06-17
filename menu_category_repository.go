package main

import (
	"errors"
	"gorm.io/gorm"
)

type MenuCategoryRepository interface {
	Create(customer *MenuCategory) error
	FindById(id uint) (MenuCategory, error)
}

type menuCategoryRepository struct {
	db *gorm.DB
}

func (m *menuCategoryRepository) Create(menuCategory *MenuCategory) error {
	result := m.db.Create(menuCategory)
	return result.Error
}

func (m *menuCategoryRepository) FindById(id uint) (MenuCategory, error) {
	var menuCategory MenuCategory
	result := m.db.First(&menuCategory, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menuCategory, nil
		} else {
			return menuCategory, err
		}
	}
	return menuCategory, nil
}

func NewMenuCategoryRepository(db *gorm.DB) MenuCategoryRepository {
	repo := new(menuCategoryRepository)
	repo.db = db
	return repo
}
