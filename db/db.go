package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" // postgres schema

	_ "github.com/joho/godotenv/autoload" // autoload env
)

var db *gorp.DbMap

func connectDB(credentials string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", credentials)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap, err
}

// GetDB - db getter
func GetDB() *gorp.DbMap {
	return db
}

// Init - initialize connection to db
func Init() {
	credentials := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = connectDB(credentials)
	if err != nil {
		log.Fatal(err)
	}
}
