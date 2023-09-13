package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CRUD Endpoints
func createPerson(w http.ResponseWriter, req *http.Request) {
	ret := Response{}
	name := req.PostFormValue("name")
	id, err := createUser(name)
	if err != nil {
		ret.Status = "409"
		ret.Message = err.Error()
		w.WriteHeader(http.StatusConflict)
		err = json.NewEncoder(w).Encode(ret)
		if err != nil {
			log.Println(err)
		}
		return
	}
	ret.Id = id
	ret.Name = name
	ret.Status = "200"
	ret.Message = "Successfully created name"
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		log.Println(err)
	}
	return
}
func readPerson(w http.ResponseWriter, req *http.Request) {
	ret := Response{}
	value := mux.Vars(req)["userId"]
	name, id, err := getUser(value)
	if err != nil {
		if err != nil {
			ret.Status = "404"
			ret.Message = err.Error()
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(ret)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
	ret.Id = id
	ret.Name = name
	ret.Status = "200"
	ret.Message = "Successful"
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		log.Println(err)
	}
	return
}
func updatePerson(w http.ResponseWriter, req *http.Request) {
	ret := Response{}
	userId := mux.Vars(req)["userId"]
	newName := req.PostFormValue("name")
	id, err := updateUser(userId, newName)
	if err != nil {
		ret.Status = "404"
		ret.Message = err.Error()
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(ret)
		if err != nil {
			log.Println(err)
		}
		return

	}
	ret.Id = id
	ret.Name = newName
	ret.Status = "200"
	ret.Message = "Successful"
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		log.Println(err)
	}
	return
}
func deletePerson(w http.ResponseWriter, req *http.Request) {
	ret := Response{}
	userId := mux.Vars(req)["userId"]

	err := deleteUser(userId)
	if err != nil {
		ret.Status = "404"
		ret.Message = err.Error()
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(ret)
		if err != nil {
			log.Println(err)
		}
		return
	}
	ret.Message = "Deleted Successfully"
	ret.Status = "200"
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		log.Println(err)
	}
	return
}
