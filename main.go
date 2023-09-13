package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"time"
)

var dbCon *pgx.Conn
var pingTicker *time.Ticker
var dbRestartTicker *time.Ticker

func init() {
	var err error
	var dbUrl = os.Getenv("PDB_URL")

	dbCon, err = pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		log.Println("DB error: ", err.Error())
	}
	/// Check if db is still connected
	/// Ping render to keep the server alive
	pingTicker = time.NewTicker(time.Minute * 10)
	dbRestartTicker = time.NewTicker(time.Second * 10)

	go func() {
		/*for range pingTicker.C {
			pingRender()
		}*/

		for {
			select {
			case <-pingTicker.C:
				pingRender()
			case <-dbRestartTicker.C:
				/*err := dbCon.Ping(context.Background())
				if err != nil {
					log.Println(err)
				}

				fmt.Println("Executed", dbCon.IsClosed())*/
				if dbCon.IsClosed() {
					err = dbCon.Close(context.Background())
					if err != nil {
						log.Println(err)
					}
					dbCon, err = pgx.Connect(context.Background(), dbUrl)
					log.Println("Restarted connection")
					if err != nil {
						log.Println("DB restart error: ", err.Error())
					}
				}
			}
		}
	}()

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
	defer pingTicker.Stop()
	defer dbRestartTicker.Stop()

	/// Launch server
	log.Fatal(srv.ListenAndServe())
}
