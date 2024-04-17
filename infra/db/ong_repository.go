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

func NewOngRepository(dbconn *sqlx.DB) interfaces.OngRepository {
	return &OngRepository{
		dbconnection: dbconn,
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

func (pr *OngRepository) FindByID(ID uniqueEntityId.ID) (*entity.Ong, error) {
	row, err := pr.dbconnection.Query(`
        SELECT
            p.id,
            p.name,
        WHERE
            id = ?`,
		ID,
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving ONG %d: %w", ID, err)
	}
	defer row.Close()

	if !row.Next() {
		return nil, sql.ErrNoRows
	}

	var ong entity.Ong

	if err := row.Scan(&pr.dbconnection, "SELECT * FROM table where id = ?", 1); err != nil {
		return nil, fmt.Errorf("error scanning ONG: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over ONG rows: %w", err)
	}

	return &ong, nil
}
