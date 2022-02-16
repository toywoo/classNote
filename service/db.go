package service

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	FALSE = iota
	TRUE
	ERROR
)

type DB struct {
	Connection *sql.DB
}

func NewDBConnection() *DB {
	jsonBytes, err := ioutil.ReadFile("./key/dbInfo.json")

	if err != nil {
		log.Println("Error: ReadFile dbInfo.json")
		panic(err.Error())
	}

	dbInfo := make(map[string]string)

	errJson := json.Unmarshal(jsonBytes, &dbInfo)

	if errJson != nil {
		log.Println("Error: fail unmarshaling jsonfile")
		panic(err.Error())
	}

	connInfoSlice := []string{"user=", dbInfo["user"], " ", "password=", dbInfo["password"], " ", "dbname=", dbInfo["dbname"], " ", "sslmode=", dbInfo["sslmode"]}
	dbConnectionInfo := strings.Join(connInfoSlice, "")

	connection, err := sql.Open("postgres", dbConnectionInfo)

	if err != nil {
		log.Println("Error: fail Connecting DB")
		panic(err.Error())
	}

	return &DB{connection} // don't worry dangling pointer
} // steady connection 이 performance 관점에서 좋음

func (db *DB) IsExistContent(title string, username string) int {
	var isExistContent bool
	query := "SELECT EXISTS(SELECT * FROM public.note WHERE title=$1 AND username=$2)"

	err := db.Connection.QueryRow(query, title, username).Scan(&isExistContent)
	if err != nil {
		return ERROR
	} else if isExistContent {
		return TRUE
	} else {
		return FALSE
	}
}
