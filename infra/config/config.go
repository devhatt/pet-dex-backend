package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db                 *sql.DB
	err                error
	StandardDateLayout = "2006-01-02"
)

func InitConfigs() *sql.DB {
	env := GetEnvConfig()

	db, err = sql.Open("mysql", env.DBUrl)
	if err != nil {
		panic(err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}
