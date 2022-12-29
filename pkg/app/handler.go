package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkennymore/banking/service"
)

/*create a struct to get customer details*/
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

/*create a function to get all customers in xml*/
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*create a list of customers
	Data mashaling
	*/
	customers := []Customer{ //get a slice of customers with dommy values initialization
		{Name: "Kenneth", City: "Lagos", ZipCode: "234"},
		{Name: "Jude", City: "Lagos", ZipCode: "234"},
		{Name: "Faith", City: "Lagos", ZipCode: "234"},
		{Name: "Micheal", City: "Abuja", ZipCode: "267"},
		{Name: "Anthony", City: "Akure", ZipCode: "654"},
	}
	/*check request header for data content type and return according to request*/
	if r.Header.Get("Content-Type") == "application/xml" {
		/*tell the header to set the right content type*/
		w.Header().Add("Content-Type", "application/xml")
		/*call the xml encoder which take an IO writer*/
		xml.NewEncoder(w).Encode(customers)
	} else {
		/*tell the header to set the right content type*/
		w.Header().Add("Content-Type", "application/json")
		/*call the json encoder which take an IO writer*/
		json.NewEncoder(w).Encode(customers)
	}
}

/*get a single customer*/
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, vars["customer_id"])
}

/*Create customers*/
func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create customers")
}
