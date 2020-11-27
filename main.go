package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	c "github.com/EtoNeJa00/postpars/controllers"

)

func main() {

	addressHandler, err := c.NewAddressHandler()
	if err!=nil{
		fmt.Println(err)
		return
	}

	port:=":"+os.Getenv("LISTEN_PORT")

	r := mux.NewRouter()
	r.HandleFunc("/api/adresses", addressHandler.GetAddresses).Methods("POST")
	http.ListenAndServe(port, r)
}


