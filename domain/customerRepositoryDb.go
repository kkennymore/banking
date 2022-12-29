package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kkennymore/banking/errorException"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT id, name, city, zipcode, date_of_birth, status FROM customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error quering table" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers" + err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errorException.AppError) {
	customerSql := "SELECT id, name, city, zipcode, date_of_birth, status FROM customers WHERE id = ?"
	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorException.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer" + err.Error())
			return nil, errorException.NewUexpectedError("Unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
