package main

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	c "github.com/EtoNeJa00/postpars/controllers"

)

func main() {
	  var addressHandler c.AddressHandler

	  port:=":"+os.Getenv("LISTEN_PORT")

	  r := mux.NewRouter()
	  r.HandleFunc("/api/adresses", addressHandler.GetAddresses).Methods("POST")
	  http.ListenAndServe(port, r)
}


