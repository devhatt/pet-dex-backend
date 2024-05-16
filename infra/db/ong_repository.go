package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"

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
	row, err := or.dbconnection.Query(`
	SELECT
		l.id,
		l.userId,
		l.links,
		l.openingHours,
		l.adoptionPolicy
	FROM
		legal_persons l
	WHERE
		l.id = ?`,
		ID,
	)
	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error retrieving ong %d: %w", ID, err)
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, sql.ErrNoRows
	}

	var ong entity.Ong

	err = row.Scan(
		&ong.ID,
		&ong.UserID,
		&ong.OpeningHours,
		&ong.AdoptionPolicy,
		&ong.Links,
	)
	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error retrieving ong %d: %w", ID, err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		logger.Error("error on ong repository: ", err)
		return nil, fmt.Errorf("error iterating over ong rows: %w", err)
	}

	return &ong, err
}
