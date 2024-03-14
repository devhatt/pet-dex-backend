package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
	"strings"

	uniqueEntity "pet-dex-backend/v2/pkg/entity"
)

type BreedRepository struct {
	dbconnection *sql.DB
}

func NewBreedRepository(db *sql.DB) interfaces.BreedRepository {
	return &BreedRepository{
		dbconnection: db,
	}
}

func (breedRepository *BreedRepository) Save(entity.Breed) error {
	return nil
}

func (breedRepository *BreedRepository) FindById(uniqueEntity.ID) (breed *entity.Breed, err error) {

	return nil, nil
}

func (breedRepository *BreedRepository) List() (breeds []*dto.BreedList, err error) {
	rows, err := breedRepository.dbconnection.Query(`
		SELECT 
			id, 						
			name, 			
			imgUrl
		FROM breeds`)
	if err != nil {
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
			return nil, fmt.Errorf("error scanning breeds: %w", err)
		}
		breeds = append(breeds, &breed)
	}

	return breeds, nil
}

func (breedRepository *BreedRepository) Update(breedID string, updatePayload map[string]interface{}) error {
	query := "UPDATE breeds SET "
	values := []interface{}{}

	for key, value := range updatePayload {
		query += key + "=?, "
		values = append(values, value)
	}

	query = strings.TrimSuffix(query, ", ")
	query += " WHERE id=?"
	values = append(values, breedID)

	_, err := breedRepository.dbconnection.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error updating breed: %w \\n", err)
	}

	return nil
}
