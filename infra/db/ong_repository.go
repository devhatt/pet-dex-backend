package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type OngRepository struct {
	dbconnection *sql.DB
}

func NewOngRepository(db *sql.DB) interfaces.OngRepository {
	return &OngRepository{
		dbconnection: db,
	}
}

func (or *OngRepository) Save(entity.Ong) error {
	rows, err := or.dbconnection.Query("INSERT INTO ongs (cnpj, email, endereco, cidade, estado, imagem, facebook, instagram) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil && err != sql.ErrConnDone {
		err = fmt.Errorf("error inserting ong: %w", err)
		return err
	}
	defer rows.Close()
	return nil
}
