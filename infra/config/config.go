package config

import (
	"database/sql"
	"time"

	"github.com/go-chi/jwtauth/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var (
	db                 *sql.DB
	logger             *Logger
	StandardDateLayout = "2006-01-02"
)

func InitConfigs() error {
	var err error
	env := GetEnvConfig()

	db, err = sql.Open("mysql", env.DBUrl)
	if err != nil {
		panic(err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
