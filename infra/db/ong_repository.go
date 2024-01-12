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

func (or *OngRepository) Save(ong entity.Ong) error {
	rows, err := or.dbconnection.Query("INSERT INTO ongs (cnpj, email, address, city, state, image, facebook, instagram) VALUES (?,?,?,?,?,?,?,?)", ong.CNPJ, ong.Email, ong.Location.Address, ong.Location.City, ong.Location.State, ong.Imagem, ong.SocialMedia.Facebook, ong.SocialMedia.Instagram)
	if err != nil && err != sql.ErrConnDone {
		err = fmt.Errorf("error inserting ong: %w", err)
		return err
	}
	defer rows.Close()
	return nil
}
