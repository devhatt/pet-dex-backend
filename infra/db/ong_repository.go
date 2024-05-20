package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
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

func (or *OngRepository) FindByID(ID uniqueEntityId.ID) (*entity.Ong, error) {
	var ong entity.Ong

	err := or.dbconnection.Get(&ong, `SELECT
	l.id,
	l.userId,
	l.links,
	l.openingHours,
	l.adoptionPolicy
FROM
	legal_persons l
WHERE
	l.id = ?`, ID)

	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error retrieving ong %d: %w", ID, err)
		return nil, err
	}

	return &ong, nil
}

func (or *OngRepository) Update(id uniqueEntityId.ID, ongToUpdate entity.Ong) error {

	query := "UPDATE legal_persons SET"
	var values []interface{}

	if ongToUpdate.Phone != "" {
		query = query + " phone =?"
		values = append(values, ongToUpdate.Phone)
	}

	if ongToUpdate.OpeningHours != "" {
		query = query + " openingHours =?"
		values = append(values, ongToUpdate.OpeningHours)
	}

	if ongToUpdate.AdoptionPolicy != "" {
		query = query + " adoptionPolicy =?"
		values = append(values, ongToUpdate.AdoptionPolicy)
	}

	if string(*ongToUpdate.Links) != "" {
		query = query + " links =?"
		values = append(values, ongToUpdate.Links)
	}

	query = query + " updated_at =?,"
	values = append(values, time.Now())

	n := len(query)
	query = query[:n-1] + " WHERE id =?"
	values = append(values, id)

	fmt.Printf("Query to update: %s", query)

	_, err := or.dbconnection.Exec(query, values...)

	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error on updating ong")
		return err
	}

	return nil
}
