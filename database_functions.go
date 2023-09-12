package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

func createUser(user string) (id string, err error) {
	exec := dbCon.QueryRow(context.Background(), "INSERT INTO persons(name) VALUES($1) RETURNING id", user)

	err = exec.Scan(&id)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			err = errors.New(fmt.Sprintf("Name: %s, already exist", user))
		}
	}

	fmt.Println("ID ", id)
	//fmt.Printf("%v\nType: %T", exec, id)
	return
}
func updateUser() {

}
func deleteUser() {

}
func getUser(value string) (name string, idDb int, err error) {
	var id int
	var person pgx.Row
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if isDigit(value) {
		id, err = strconv.Atoi(value)
		if err != nil {
			log.Println(err)
			return "", 0, err
		}
		person = dbCon.QueryRow(ctx, ""+
			"SELECT name, id FROM persons WHERE id = $1", id)
	} else {
		person = dbCon.QueryRow(ctx, ""+
			"SELECT name, id FROM persons WHERE name = $1", value)
	}
	err = person.Scan(&name, &idDb)
	if err != nil {
		log.Println("Error from, GetUser: ", err)
	}
	fmt.Println(name)
	return
}
