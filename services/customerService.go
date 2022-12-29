package services

import (
	"github.com/kkennymore/banking/domain"
	"github.com/kkennymore/banking/errorException"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errorException.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

/*all customers details, to return customers data*/
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

/*single customer service to rturn a customer data*/
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errorException.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
