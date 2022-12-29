package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

/*create a struct to get customer details*/
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

/*the index function*/
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

/*create a function to get all customers in xml*/
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
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
