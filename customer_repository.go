package main

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CustomerRepository interface {
	Create(customer *Customer) error
	FindById(id uint) (Customer, error)
	Retrieve() ([]Customer, error)
	FindFirstBy(by map[string]interface{}) (Customer, error)
	FindAllBy(by map[string]interface{}) ([]Customer, error)
	FindBy(by string, vals ...interface{}) ([]Customer, error)
	Update(customer *Customer, by map[string]interface{}) error
	UpdateBy(existingCustomer *Customer) error
	Delete(id uint) error
	BaseRepositoryAggregation
	BaseRepositoryPaging
	BaseRepositoryRaw
	BaseRepositoryAdvQuery
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) preload(model string) {
	c.db.Preload(model)
}

func (c *customerRepository) Create(customer *Customer) error {
	result := c.db.Create(customer)
	return result.Error
}

func (c *customerRepository) FindBy(by string, vals ...interface{}) ([]Customer, error) {
	var customers []Customer
	result := c.db.Where(by, vals...).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}

// When querying with struct, GORM will only query with non-zero fields
func (c *customerRepository) FindFirstBy(by map[string]interface{}) (Customer, error) {
	var customer Customer
	result := c.db.Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error) {
	var customer Customer
	result := c.db.Preload(preload).Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}
func (c *customerRepository) FindFirstAllPreload(by map[string]interface{}) (interface{}, error) {
	var customer Customer
	result := c.db.Preload(clause.Associations).Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindAllBy(by map[string]interface{}) ([]Customer, error) {
	var customers []Customer
	result := c.db.Where(by).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}
func (c *customerRepository) FindById(id uint) (Customer, error) {
	var customer Customer
	result := c.db.First(&customer, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) Retrieve() ([]Customer, error) {
	var customers []Customer
	result := c.db.First(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}

func (c *customerRepository) UpdateBy(existingCustomer *Customer) error {
	result := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(existingCustomer)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Update(customer *Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Delete(id uint) error {
	result := c.db.Delete(&Customer{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Count(groupBy string) (int64, error) {
	var total int64
	result := c.db.Model(&Customer{}).Select("count(*)").Group(groupBy).First(&total)
	if err := result.Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (c *customerRepository) GroupBy(result interface{}, selectBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&Customer{}).Select(selectBy).Where(whereBy).Group(groupBy).Find(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (c *customerRepository) Paging(itemPerPage int, page int) (interface{}, error) {
	var customers []Customer
	offset := (page - 1) * itemPerPage
	res := c.db.Order("customer_id").Limit(itemPerPage).Offset(offset).Find(&customers)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}

func (c *customerRepository) Query(result interface{}, sql string, vals ...interface{}) error {
	res := c.db.Raw(sql, vals...).Scan(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
