package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func runner() {

	/*Request multiplexer*/
	//mux := http.NewServeMux()
	mux := mux.NewRouter()
	/*registering the pattern and function with the default multiplexer
	defining the routes
	*/

	mux.HandleFunc("/index", index)
	mux.HandleFunc("/customers", getAllCustomers)
	/*starting the server*/
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
