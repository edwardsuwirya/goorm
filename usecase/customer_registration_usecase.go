package usecase

import (
	"enigmacamp.com/goorm/model"
	"enigmacamp.com/goorm/repository"
)

type CustomerRegistrationUseCase interface {
	Register(customer *model.Customer) error
}

type customerRegistrationUseCase struct {
	repo repository.CustomerRepository
}

func (uc *customerRegistrationUseCase) Register(customer *model.Customer) error {
	return uc.repo.Create(customer)
}
func NewCustomerRegistrationUseCase(repo repository.CustomerRepository) CustomerRegistrationUseCase {
	return &customerRegistrationUseCase{
		repo: repo,
	}
}
