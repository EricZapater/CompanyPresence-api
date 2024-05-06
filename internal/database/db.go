package database

import (
	"companypresence-api/internal/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func NewDatabase()(*Database, error){
	env := config.LoadEnvironment()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", env.DbHost, env.DbUser, env.DbPass, env.DbName)
	db, err := sql.Open("postgres", dsn)
	if err !=nil{
		return nil, err
	}
	db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)
	
	return &Database{DB: db}, nil
}