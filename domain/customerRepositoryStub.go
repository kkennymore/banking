package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "001", Name: "Kenneth", City: "Lagos", ZipCode: "234", DateOfBirth: "1982-7-24", Status: "1"},
		{Id: "002", Name: "Joseph", City: "Benin", ZipCode: "234", DateOfBirth: "2002-8-1", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
