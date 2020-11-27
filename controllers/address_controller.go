package controllers

import (
	"strings"
	"sort"
	"net/http"
	"encoding/json"
	"github.com/EtoNeJa00/postpars/models"
)

type AddressHandler struct {
	Dictionary models.Dictionary
}

func NewAddressHandler()(AddressHandler, error){
	var adressHandl AddressHandler

	dict, err := models.GetDict()
	if err != nil{
		return adressHandl, err
	}
	
	adressHandl.Dictionary=dict

	return adressHandl, nil
}

func (a *AddressHandler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var request models.AddressRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if (err!= nil){
		if strings.Contains(err.Error(), "internal"){
			w.WriteHeader(500)
		}else{
			w.WriteHeader(http.StatusBadRequest)
		}
		response := models.AddressResponse{ 
			Res_type : request.ReqType,
			Result : "fail",
			Data : err.Error(),
		}
		json.NewEncoder(w).Encode(response)
	}else{
		response := models.AddressResponse{ 
			Res_type : request.ReqType,
			Result : "successful",
			Data : getResultString(request.Data,a.Dictionary.States),
		}
		json.NewEncoder(w).Encode(response)
	}
}

func getResultString(data []models.Person, dict map[string]string) string{
	for i:=0; i<len(data);i++{
		data[i].State=dict[data[i].State]
	}

	sortRuls := func(i,j int) bool{
		if data[i].State == data[j].State{
			return data[i].Name < data[j].Name 
		} else {
			return data[i].State < data[j].State
		}
	}
	sort.Slice(data, sortRuls)

	var resultstr string
	for i:=0; i < len(data); i++{
		if (data[i].Name=="")||(data[i].Address=="")||(data[i].City=="")||(data[i].State==""){
			continue
		}
		if resultstr==""{
			resultstr = data[i].State+data[i].ToString()

		} else if data[i].State==data[i-1].State{
			resultstr += data[i].ToString()

		} else {
			resultstr += "\n" +" "+data[i].State + data[i].ToString()
		}
	}
	return resultstr
}