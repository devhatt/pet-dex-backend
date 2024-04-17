package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OngRepository struct {
	dbconnection *sqlx.DB
}

// FindByID implements interfaces.OngRepository.
func (pr *OngRepository) FindByID(ID uuid.UUID) (*entity.Ong, error) {
	row, err := pr.dbconnection.Query(`
        SELECT
            p.id,
            p.name,
            p.address,
            p.phone
        FROM
            ongs
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

	if err := row.Scan(
		&ong.ID,
		&ong.Name,
	); err != nil {
		return nil, fmt.Errorf("error scanning ONG: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over ONG rows: %w", err)
	}

	return &ong, nil
}

func NewOngRepository(dbconn *sqlx.DB) interfaces.OngRepository {
	return &OngRepository{
		dbconnection: dbconn,
	}
}
