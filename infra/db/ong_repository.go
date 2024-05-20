package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"

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

func (or *OngRepository) List(limit, offset int, sortBy, order string) (ongs []*dto.OngList, err error) {
	query := fmt.Sprintf(`
	SELECT 
		id, 
		name, 
		city, 
		phone, 
		state, 
		openingHours, 
		adoptionPolicy, 
		links 
	FROM legal_persons 
	ORDER BY %s %s LIMIT $1 OFFSET $2`, sortBy, order)
	rows, err := or.dbconnection.Query(query, limit, offset)
	if err != nil {
		logger.Error("error listing ongs", err)
		return nil, fmt.Errorf("error listing ongs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ong dto.OngList
		err := rows.Scan(
			&ong.ID,
			&ong.Phone,
			&ong.OpeningHours,
			&ong.AdoptionPolicy,
			&ong.Links,
		)
		if err != nil {
			logger.Error("error scanning ongs", err)
			return nil, fmt.Errorf("error scanning ongs: %w", err)
		}
		ongs = append(ongs, &ong)
	}

	return ongs, nil
}
