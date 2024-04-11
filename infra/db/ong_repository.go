package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"

	"github.com/jmoiron/sqlx"
)

type OngRepository struct {
	dbconnection *sqlx.DB
}

func NewOngRepository(db *sqlx.DB) interfaces.OngRepository {
	return &OngRepository{
		dbconnection: db,
	}
}

func (or *OngRepository) Save(ong *entity.Ong) error {
	ongId := uniqueEntityId.NewID()

	_, err := or.dbconnection.Query("INSERT INTO legal_persons (id, userId, phone, links, openingHours, adoptionPolicy) VALUES (?,?,?,?,?,?)", ongId, user.Id, user.phone, ong.links, ong.openingHours, ong.adoptionPolicy)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngRepository.Save error: %w", err))
		err = fmt.Errorf("error on saving ong")
		return err
	}

	return nil
}
