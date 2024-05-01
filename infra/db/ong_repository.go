package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"

	"github.com/jmoiron/sqlx"
)

type OngRepository struct {
	dbconnection *sqlx.DB
	logger       config.Logger
}

func NewOngRepository(db *sqlx.DB) interfaces.OngRepository {
	return &OngRepository{
		dbconnection: db,
		logger:       *config.GetLogger("ong-repository"),
	}
}

func (or *OngRepository) Save(ong *entity.Ong) error {

	_, err := or.dbconnection.NamedExec("INSERT INTO legal_persons (id, userId, phone, links, openingHours, adoptionPolicy) VALUES (:id, :userId, :phone, :links, :openingHours, :adoptionPolicy)", &ong)

	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error on saving ong")
		return err
	}

	return nil
}

func (or *OngRepository) Update(ongId uniqueEntityId.ID, ongDto *dto.OngUpdateDto) error {

	query := "UPDATE legal_persons SET"
	values := []interface{}{}

	/*
		Add queries

	*/

	query = query + " updated_at =?,"
	values = append(values, time.Now())

	n := len(query)
	query = query[:n-1] + " WHERE id =?"
	values = append(values, ongId)

	fmt.Printf("Query to update: %s", query)

	_, err := or.dbconnection.Exec(query, values...)

	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error on updating ong")
		return err
	}

	return nil
}
