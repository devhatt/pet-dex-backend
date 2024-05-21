package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"

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

func (breedRepository *BreedRepository) FindByID(ID uniqueEntityId.ID) (*entity.Breed, error) {
	row, err := breedRepository.dbconnection.Query(`
		SELECT
			id,
			name,
			specie,
			size,
			description,
			height,
			weight,
			physicalChar,
			disposition,
			idealFor,
			fur,
			imgUrl,
			weather,
			dressage,
			lifeExpectancy
		FROM breeds
		WHERE
			id = ?;`, 
		ID,
	)

	if err != nil {
		return nil, fmt.Errorf("error retrieving breed %d: %w", ID, err)
	}
	defer row.Close()

	if !row.Next() {
		return nil, sql.ErrNoRows
	}

	var breed entity.Breed;
	err = row.Scan(&breed.ID, &breed.Name, &breed.Specie, &breed.Size, &breed.Description, &breed.Height, &breed.Weight, &breed.PhysicalChar, &breed.Disposition, &breed.IdealFor, &breed.Fur, &breed.ImgUrl, &breed.Weather, &breed.Dressage, &breed.LifeExpectancy)
	if err!= nil {
		return nil, fmt.Errorf("error scanning row into breed: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over breed rows: %w", err)
	}

	return &breed, nil
}