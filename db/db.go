package db

import (
	"database/sql"
	"echo-rest-postgres/config"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses terkoneksi")

}

func CreateCon() *sql.DB {
	return db
}
