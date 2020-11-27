package models

import (
	"fmt"
	"errors"
	"io/ioutil"
	"encoding/json"
	"os"
)

	type Dictionary struct{ 
		States map[string]string `json:"state_dictionary"`
	} 

	func GetDict ()( Dictionary, error){
		var stateDict  Dictionary
		
		file, err := os.Open("dictionaries.json")
		if err != nil{
			fmt.Println(err)
			return stateDict, errors.New("internal server error")
		}
	
		defer file.Close()

		byteValue, _ := ioutil.ReadAll(file)
		if err != nil{
			fmt.Println(err)
			return stateDict, errors.New("internal server error")
		}

		err = json.Unmarshal(byteValue, &stateDict)
		if err != nil{
			fmt.Println(err)
			return stateDict, errors.New("internal server error")
		}
		
		return stateDict, nil
	}