package main

import (
	"context"
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
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  20 * time.Second,
	}
	api := r.PathPrefix("/api").Subrouter()
	api.Use(jsonContent)

	/// Create (POST)
	api.Path("").HandlerFunc(createPerson).Methods(http.MethodPost)
	/// Read (GET)
	api.Path("/{userId}").HandlerFunc(readPerson).Methods(http.MethodGet)
	/// Update (PUT)
	api.Path("/{userId}").HandlerFunc(updatePerson).Methods(http.MethodPut)
	/// Delete (DELETE)
	api.Path("/{userId}").HandlerFunc(deletePerson).Methods(http.MethodDelete)
	defer func(dbCon *pgx.Conn, ctx context.Context) {
		err := dbCon.Close(ctx)
		if err != nil {

		}
	}(dbCon, context.Background())

	/// Launch server
	log.Fatal(srv.ListenAndServe())
}
