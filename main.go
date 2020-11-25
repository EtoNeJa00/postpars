package main

import (
	"sort"
	"encoding/json"
	"fmt"
	models "github.com/EtoNeJa00/postpars/models"
)



func main() {
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

	var request models.Request
	err := json.Unmarshal([]byte(myJsonString), &request)


	fmt.Println(err)

	sortRuls := func(i,j int) bool{
		if request.Data[i].State == request.Data[j].State{
			return request.Data[i].Name < request.Data[j].Name 
		} else {
			return request.Data[i].State < request.Data[j].State
		}
	}
	sort.Slice(request.Data, sortRuls)

	var resultstr string
	for i:=0; i < len(request.Data); i++{
		if (request.Data[i].Name=="")||(request.Data[i].Address=="")||(request.Data[i].City=="")||(request.Data[i].State==""){
			continue
		}
		if resultstr==""{
			resultstr = request.Data[i].State+request.Data[i].ToString()

		} else if request.Data[i].State==request.Data[i-1].State{
			resultstr += request.Data[i].ToString()

		} else {
			resultstr += "\n" +" "+request.Data[i].State +  request.Data[i].ToString()
		}
	}

	fmt.Println(resultstr)
}


