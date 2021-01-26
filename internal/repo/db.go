package repo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbInstance *sqlx.DB

func InitDb() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("DB_HOST is not set")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, errors.New("DB_PORT is not set")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("DB_USER is not set")
	}
	pass := os.Getenv("DB_USER_PASS")
	if pass == "" {
		return nil, errors.New("DB_USER_PASS is not set")
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, errors.New("DB_NAME is not set")
	}
	mode := os.Getenv("DB_SSL_MODE")
	if mode == "" {
		return nil, errors.New("DB_SSL_MODE is not set")
	}
	moc := os.Getenv("DB_MAX_OPEN_CONN")
	if moc == "" {
		return nil, errors.New("DB_MAX_OPEN_CONN is not set")
	}
	conns, err := strconv.Atoi(moc)
	if err != nil {
		return nil, err
	}

	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=15",
		host, port, user, pass, name, mode)

	driverName := "postgres"

	db, err := sqlx.Connect(driverName, ds)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db.SetMaxOpenConns(conns)
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetMaxIdleConns(0)
	return db, nil
}

// Set db instance
func SetDb(db *sqlx.DB) {
	dbInstance = db
}

// Get db instance
func GetDb() (*sqlx.DB, error) {
	if dbInstance == nil {
		return nil, errors.New("db instance is empty")
	}
	return dbInstance, nil
}
