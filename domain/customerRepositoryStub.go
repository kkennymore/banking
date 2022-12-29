package domain

/*This is the domain object*/
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

/*create an adapter for this port*/
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "001", Name: "Kenneth", City: "Lagos", ZipCode: "332133", DateOfBirth: "1982-7-24", Status: "1"},
		{Id: "002", Name: "Joseph", City: "Portharcurt", ZipCode: "75829", DateOfBirth: "2002-8-1", Status: "1"},
		{Id: "003", Name: "Pepsi Luke", City: "Lugali, Kenya", ZipCode: "298161", DateOfBirth: "1982-7-24", Status: "1"},
		{Id: "004", Name: "Lolo Jame", City: "London, UK", ZipCode: "87662", DateOfBirth: "2002-8-1", Status: "1"},
		{Id: "005", Name: "Jessica Liu", City: "Hangzhou, China", ZipCode: "245141", DateOfBirth: "1982-7-24", Status: "1"},
		{Id: "006", Name: "Michel Muhamed", City: "Benin", ZipCode: "234", DateOfBirth: "2002-8-1", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
