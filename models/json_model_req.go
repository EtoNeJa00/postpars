package models

import (
	"strings"
	"errors"
	"encoding/json"
)

type AddressRequest struct {
	ReqType string 	`json:"req_type"`
	Data []Person	`json:"data"`
}

type Person struct{
	Name string 
	Address string
	City string
	State string
}

func (p *Person) UnmarshalJSON(data []byte) error {
	type temp struct{
		TempStr string `json:"item"`
	}
	
	var item temp
	json.Unmarshal([]byte(data), &item) 

	if strings.TrimSpace(item.TempStr) == ""{
		return nil
	}
	
	personArr := strings.Split(item.TempStr, ",")
	if len(personArr)!=3{
		return errors.New("invalid format item: "+item.TempStr)
	}
	for i := range personArr {
		personArr[i] = strings.TrimSpace(personArr[i])
	}

	cityState := strings.Split(personArr[2]," ")
	if len(cityState)<2{
		return errors.New("invalid format in item: "+item.TempStr+" in "+personArr[2])
	}

	p.Name = personArr[0]
	p.Address = personArr[1]
	for i := 0;i<len(cityState)-1; i++{
		if i==0 {
			p.City=cityState[i]
			continue
		}
		p.City += " "+cityState[i]
	}

	state := cityState[len(cityState)-1]
	if state==""{
		return errors.New("invalid format in item: "+item.TempStr+" in "+personArr[2])
	}

	p.State = state

	return nil
}

func (p *Person)ToString () string{
	return "\n..... "+p.Name+" "+p.Address+" "+p.City+" "+p.State
}
