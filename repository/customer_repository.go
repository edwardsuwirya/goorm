package repository

import (
	"enigmacamp.com/goorm/model"
	"errors"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	FindById(id uint) (model.Customer, error)
	Retrieve() ([]model.Customer, error)
	FindFirstBy(by map[string]interface{}) (model.Customer, error)
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)
	FindBy(by string, vals ...interface{}) ([]model.Customer, error)
	Update(customer *model.Customer, by map[string]interface{}) error
	Delete(id uint) error
	BaseRepositoryAggregation
	BaseRepositoryPaging
	BaseRepositoryRaw
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer)
	return result.Error
}

func (c *customerRepository) FindBy(by string, vals ...interface{}) ([]model.Customer, error) {
	var customers []model.Customer
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
func (c *customerRepository) FindFirstBy(by map[string]interface{}) (model.Customer, error) {
	var customer model.Customer
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
func (c *customerRepository) FindAllBy(by map[string]interface{}) ([]model.Customer, error) {
	var customers []model.Customer
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
func (c *customerRepository) FindById(id uint) (model.Customer, error) {
	var customer model.Customer
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

func (c *customerRepository) Retrieve() ([]model.Customer, error) {
	var customers []model.Customer
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

func (c *customerRepository) Update(customer *model.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Delete(id uint) error {
	result := c.db.Delete(&model.Customer{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Count(groupBy string) (int64, error) {
	var total int64
	result := c.db.Model(&model.Customer{}).Select("count(*)").Group(groupBy).First(&total)
	if err := result.Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (c *customerRepository) GroupBy(result interface{}, selectBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&model.Customer{}).Select(selectBy).Where(whereBy).Group(groupBy).Find(result)
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
	var customers []model.Customer
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
