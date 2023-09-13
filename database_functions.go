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

func createUser(user string) (id int, err error) {
	exec := dbCon.QueryRow(context.Background(), "INSERT INTO persons(name) VALUES($1) RETURNING id", user)

	err = exec.Scan(&id)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			err = errors.New(fmt.Sprintf("Name: %s, already exist", user))
		}
	}

	fmt.Println("ID ", id)

	return
}
func updateUser(userId string, newName string) (id int, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if isDigit(userId) {
		id, err = strconv.Atoi(userId)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = dbCon.Exec(ctx, ""+
			"UPDATE persons SET name = $1 WHERE id = $2", newName, id)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		p := dbCon.QueryRow(ctx, ""+
			"UPDATE persons SET name = $1 WHERE  id = $2 RETURNING id", newName, userId)
		err = p.Scan(&id)
		if err != nil {
			log.Println("Error from, GetUser: ", err)
		}
	}

	return
}
func deleteUser(value string) (err error) {

	var id int
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if isDigit(value) {
		id, err = strconv.Atoi(value)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = dbCon.Exec(ctx, ""+
			"DELETE FROM persons WHERE id = $1", id)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		_, err = dbCon.Exec(ctx, ""+
			"DELETE FROM persons WHERE name = $1", value)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return

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
