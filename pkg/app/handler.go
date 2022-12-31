package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkennymore/banking/services"
)

type CustomerHandlers struct {
	service services.CustomerService
}

/*create a function to get all customers in xml*/
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*return data from stub*/
	customers, _ := ch.service.GetAllCustomer()

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
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)

	} else {
		/*check request header for data content type and return according to request*/
		if r.Header.Get("Content-Type") == "application/xml" {
			/*tell the header to set the right content type*/
			w.Header().Add("Content-Type", "application/xml")
			/*call the xml encoder which take an IO writer*/
			xml.NewEncoder(w).Encode(customer)
		} else {
			/*tell the header to set the right content type*/
			w.Header().Add("Content-Type", "application/json")
			/*call the json encoder which take an IO writer*/
			json.NewEncoder(w).Encode(customer)
		}
	}
}

/*Create customers*/
func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create customers")
}
