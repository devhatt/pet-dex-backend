package db

import "github.com/jmoiron/sqlx"

type OngRepository struct {
	dbconnection *sqlx.DB
}
