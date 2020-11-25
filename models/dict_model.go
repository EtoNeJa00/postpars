package models

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

	type StateDictionary struct{ 
		States map[string]string `json:"state_dictionary"`
	} 

	func GetDict ()StateDictionary{
		file, err := os.Open("dictionaries.json")
		if err != nil{
			fmt.Print(err)
		}
	
		defer file.Close()

		byteValue, _ := ioutil.ReadAll(file)
		if err != nil{
			fmt.Print(err)
		}
		
		var stateDict StateDictionary

		err = json.Unmarshal(byteValue, &stateDict)
		if err != nil{
			fmt.Print(err)
		}
		return stateDict
	}