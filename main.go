package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"time"
)

const (
	dbUrl = ""
)

var dbCon *pgx.Conn

func init() {
	var err error
	dbCon, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Println("DB error: ", err.Error())
	}

}
func main() {
	log.SetFlags(log.Lshortfile)
	r := mux.NewRouter().StrictSlash(true)
	//r.Headers("Content-Type", "application/json")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  20 * time.Second,
	}
	api := r.PathPrefix("/api").Subrouter()
	api.Use(jsonContent)

	//api.NewRoute().HandlerFunc(createPerson).Methods(http.MethodPost)
	//api.Path("/{userId}").HandlerFunc(createPerson).Methods(http.MethodPost)

	/// Create (POST)
	api.Path("").HandlerFunc(createPerson).Methods(http.MethodPost)
	/// Read (Get)
	api.Path("/{userId}").HandlerFunc(readPerson).Methods(http.MethodGet)
	/// Update (PUT)
	api.Path("/{userId}").HandlerFunc(updatePerson).Methods(http.MethodPut)
	/// Delete
	api.Path("/{userId}").HandlerFunc(deletePerson).Methods(http.MethodDelete)
	defer func(dbCon *pgx.Conn, ctx context.Context) {
		err := dbCon.Close(ctx)
		if err != nil {

		}
	}(dbCon, context.Background())

	/// Launch server
	log.Fatal(srv.ListenAndServe())
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	ret := Response{}
	name := req.PostFormValue("name")
	id, err := createUser(name)
	if err != nil {
		ret.Status = "404"
		ret.Message = err.Error()
		w.WriteHeader(http.StatusConflict)
		err = json.NewEncoder(w).Encode(ret)
		if err != nil {
			log.Println(err)
		}
		return
	}
	ret.Status = "200"
	ret.Message = fmt.Sprintf("Created: %s with id: %s", name, id)
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
			w.WriteHeader(http.StatusConflict)
			err = json.NewEncoder(w).Encode(ret)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
	ret.Status = "200"
	ret.Message = fmt.Sprintf("Name: %s with id: %d", name, id)
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		log.Println(err)
	}
	return
}
func updatePerson(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Thanks from update")
}
func deletePerson(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Thanks from delete")
}
