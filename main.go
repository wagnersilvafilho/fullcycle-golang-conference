package main

import (
	"errors"
	"fmt"
)

type DB struct {
	opened bool
}

func open() *DB {
	db := new(DB)
	db.opened = true
	fmt.Println("connection opened")

	return db
}

func save(db *DB) error {
	if !db.opened {
		return errors.New("could not save due to connection closed")
	}
	fmt.Println("data is saved")

	return nil
}

func close(db *DB) {
	if !db.opened {
		fmt.Println("connection already closed")
		fmt.Println("db status: ", db.opened)
		return
	}
	db.opened = false
	fmt.Println("connection is closed")
	fmt.Println("db status: ", db.opened)
}

func main() {
	db := open()
	close(db)
	defer close(db)

	err := save(db)
	fmt.Println(err)
}
