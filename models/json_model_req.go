package models

import (
	"strings"
	"errors"
	"encoding/json"
	"fmt"
)

type Request struct {
	Req_type string 	`json:"req_type"`
	Data []person		`json:"data"`
}

type person struct{
	Name string 
	Address string
	City string
	State string
}

func (p *person) UnmarshalJSON(data []byte) error {
	type temp struct{
		Temp_str string `json:"item"`
	}
	var item temp
	json.Unmarshal([]byte(data), &item) 

	if strings.TrimSpace(item.Temp_str) == ""{
		return nil
	}
	
	person_arr := strings.Split(item.Temp_str, ",")
	if len(person_arr)!=3{
		return errors.New("Invalid format JSON:item")
	}
	for i := range person_arr {
		person_arr[i] = strings.TrimSpace(person_arr[i])
	}



	city_state := strings.Split(person_arr[2]," ")
	if len(city_state)<2{
		return errors.New("Invalid JSON: item: city and state")
	}
//to do
	states_map := map[string]string{
		"AZ": "Arizona",
		"CA": "California",
		"ID": "Idaho",
		"IN": "Indiana",
		"MA": "Massachusetts",
		"OK": "Oklahoma",
		"PA": "Pennsylvania",
		"VA": "Virginia",
	}

	p.Name = person_arr[0]
	p.Address = person_arr[1]
	for i := 0;i<len(city_state)-1; i++{
		if i==0 {
			p.City=city_state[i]
			continue
		}
		p.City += " "+city_state[i]
	}
	p.State = states_map[city_state[len(city_state)-1]]

	fmt.Println(p)
	return nil
}
