package main

import (
	"encoding/json"
	//"os"
	"fmt"
	model "github.com/EtoNeJa00/postpars/models"
)



func main() {
	// Open our jsonFile
	myJsonString := `{
		"req_type": "parseAddress",
		"data": [
		  {
			"item":"John Daggett, 341 King Road, Plymouth MA"
		  },
		  {
			"item":"Alice Ford, 22 East Broadway, Richmond VA"
		  },
		  {
			"item":"Orville Thomas, 11345 Oak Bridge Road, Tulsa OK"
		  },
		  {
			"item": "Terry Kalkas, 402 Lans Road, Beaver Falls PA"
		  },
		  {
			"item": " Eric Adams, 20 Post Road, Sudbury MA"
		  },
		  {
			"item": "Hubert Sims, 328A Brook Road, Roanoke VA"
		  },
		  {
			"item": "Amy Wilde, 334 Bayshore Pkwy, Mountain View CA"
		  },
		  {
			"item": ""
		  },
		  {
			"item": "Sal Carpenter, 73 6th Street, Boston MA"
		  }
		]
	  }`

	var test model.Request
	err := json.Unmarshal([]byte(myJsonString), &test)
	fmt.Println(err)
	fmt.Println(test.Data)


}
