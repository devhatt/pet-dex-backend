package db

import (
	"fmt"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"

	"github.com/jmoiron/sqlx"
)

var logger = config.GetLogger("breed-repository")

type BreedRepository struct {
	dbconnection *sqlx.DB
}

func NewBreedRepository(dbconn *sqlx.DB) interfaces.BreedRepository {
	return &BreedRepository{
		dbconnection: dbconn,
	}
}

func (breedRepository *BreedRepository) List() (breeds []*dto.BreedList, err error) {
	rows, err := breedRepository.dbconnection.Query(`
		SELECT 
			id, 						
			name, 			
			imgUrl
		FROM breeds`)
	if err != nil {
		logger.Error("error listing breeds", err)
		return nil, fmt.Errorf("error listing breeds: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var breed dto.BreedList
		err := rows.Scan(
			&breed.ID,
			&breed.Name,
			&breed.ImgUrl,
		)
		if err != nil {
			logger.Error("error scanning breeds", err)
			return nil, fmt.Errorf("error scanning breeds: %w", err)
		}
		breeds = append(breeds, &breed)
	}

	return breeds, nil
}
