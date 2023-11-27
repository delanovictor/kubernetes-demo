package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Message struct {
	Id    int64
	Value string
}

var Db *sql.DB

func init() {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		// "localhost", 5432, "user", "pass", "postgres",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	Db = db

	fmt.Println("Successfully connected!")

	initData, err := os.ReadFile("./database/init.sql")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(initData))

	if err != nil {
		panic(err)
	}

	Db = db
}

func ListMessages() ([]*Message, error) {

	rows, err := Db.Query(`SELECT id, value FROM messages ORDER BY id DESC`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []*Message

	for rows.Next() {

		var message Message

		err = rows.Scan(
			&message.Id,
			&message.Value,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		messages = append(messages, &message)
	}

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func AddMessage(message string) error {

	stmt, err := Db.Prepare("INSERT INTO messages(value) VALUES ($1)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(message)

	if err != nil {
		return err
	}
	// id, err := res.LastInsertId()

	// if err != nil {
	// 	return nil, err
	// }

	return nil
}
