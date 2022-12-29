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

	router.HandleFunc("/index", index)
	router.HandleFunc("/customers", getAllCustomers)
	/*starting the server*/
	log.Fatal(http.ListenAndServe("localhost:4000", router))
}
