package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	c "github.com/EtoNeJa00/postpars/controllers"

)



func main() {
	  var addressHandler c.AddressHandler

	  fmt.Println("BAR:", os.Getenv("LISTEN_PORT"))

	  r := mux.NewRouter()
	  r.HandleFunc("/api/adresses", addressHandler.GetAddresses).Methods("POST")
	  http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), r)
}


