package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Runner() {

	/*Request multiplexer*/
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	/*registering the pattern and function with the default multiplexer
	defining the routes
	*/
	router.HandleFunc("/index", index).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	/*Post*/
	router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	/*starting the server*/
	log.Fatal(http.ListenAndServe("localhost:4000", router))
}
