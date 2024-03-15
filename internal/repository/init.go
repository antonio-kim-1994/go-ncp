package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var db *sql.DB

func InitDB() error {
	var (
		err error

		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		passwd   = os.Getenv("DB_PASSWD")
		database = "devops"

		addr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, host, port, database)
	)

	db, err = sql.Open("mysql", addr)
	if err != nil {
		log.Fatal().Msg("Failed to login MySQL server. Check user and server information.")
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}
