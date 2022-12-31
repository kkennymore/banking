package domain

import (
	"github.com/kkennymore/banking/dto"
	"github.com/kkennymore/banking/errorException"
)

type Customer struct {
	Id          string `db:"id"`
	Name        string `db:"name" `
	City        string `db:"city"`
	ZipCode     string `db:"zip_code"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}

}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, *errorException.AppError)
}
