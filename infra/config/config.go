package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func InitConfigs() *sql.DB {
	db, err = sql.Open("mysql", "maria:123@tcp(localhost:3306)/petdex")
	if err != nil {
		panic(err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}
