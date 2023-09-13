package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	PersonNotFoundDB = "no rows in result set"
	// PersonNotDeleted This is an error message
	PersonNotDeleted = "person not deleted. incorrect input parameter"
	// PersonNotUpdated This is an error message
	PersonNotUpdated = "person not updated. incorrect input parameter"
)

func createUser(user string) (id int, err error) {
	exec := dbCon.QueryRow(context.Background(), "INSERT INTO persons(name) VALUES($1) RETURNING id", user)

	err = exec.Scan(&id)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			err = errors.New(fmt.Sprintf("name: %s, already exist", user))
		}
	}

	fmt.Println("ID ", id)

	return
}
func updateUser(userId string, newName string) (id int, err error) {
	var deletedRow pgconn.CommandTag
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if isDigit(userId) {
		id, err = strconv.Atoi(userId)
		if err != nil {
			log.Println(err)
			return
		}
		deletedRow, err = dbCon.Exec(ctx, ""+
			"UPDATE persons SET name = $1 WHERE id = $2", newName, id)
		if err != nil {
			log.Println(err)
			return
		}
		if deletedRow.RowsAffected() == 0 {
			err = errors.New(PersonNotUpdated)
			return
		}
	} else {
		p := dbCon.QueryRow(ctx, ""+
			"UPDATE persons SET name = $1 WHERE  name = $2 RETURNING id", newName, userId)
		err = p.Scan(&id)
		if err != nil {
			if strings.Contains(err.Error(), PersonNotFoundDB) {
				err = errors.New(PersonNotUpdated)
				return
			}
			log.Println("error: ", err)
		}

	}

	return
}
func deleteUser(value string) (err error) {

	var id int
	var deletedRow pgconn.CommandTag
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if isDigit(value) {
		id, err = strconv.Atoi(value)
		if err != nil {
			log.Println(err)
			return
		}
		deletedRow, err = dbCon.Exec(ctx, ""+
			"DELETE FROM persons WHERE id = $1", id)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		deletedRow, err = dbCon.Exec(ctx, ""+
			"DELETE FROM persons WHERE name = $1", value)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if deletedRow.RowsAffected() == 0 {
		err = errors.New(PersonNotDeleted)
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
		if strings.Contains(err.Error(), PersonNotFoundDB) {
			err = errors.New("person not found")
			return
		}
		log.Println("error: ", err)
	}
	return
}
