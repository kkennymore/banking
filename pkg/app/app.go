package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kkennymore/banking/domain"
	"github.com/kkennymore/banking/services"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined")
	}
}

func Runner() {

	/*check if all the environment variables are properly loaded*/
	sanityCheck()
	/*Request multiplexer*/
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	/*stub injection into our services during wiring and coupling*/
	//ch := CustomerHandlers{services.NewCustomerService(domain.NewCustomerRepositoryStub())}
	/*database repository injection into service dury wiring and coupling*/
	ch := CustomerHandlers{services.NewCustomerService(domain.NewCustomerRepositoryDb())}
	/*registering the pattern and function with the default multiplexer
	defining the routes
	*/
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	/*Post*/
	router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	/*starting the server*/
	fmt.Println("Server started")
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
