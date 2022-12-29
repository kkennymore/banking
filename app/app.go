package app

import (
	"log"
	"net/http"
)

func Start() {
	/*registering the pattern and function with the default multiplexer
	defining the routes
	*/
	http.HandleFunc("/index", index)
	http.HandleFunc("/customers", getAllCustomers)
	/*starting the server*/
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
