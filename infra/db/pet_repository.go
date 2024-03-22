package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"strings"
	"time"

	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type PetRepository struct {
	dbconnection *sql.DB
}

func NewPetRepository(db *sql.DB) interfaces.PetRepository {
	return &PetRepository{
		dbconnection: db,
	}
}

func (pr *PetRepository) Save(entity.Pet) error {
	return nil
}

func (pr *PetRepository) FindById(id int) (pet *entity.Pet, err error) {
	//var petToRecive entity.Pet
	//err = pr.dbconnection.QueryRow("SELECT id, name, localization_ong, pet_details, social_media_ong FROM pet WHERE id = ?", id).Scan(&petToRecive.Id, &petToRecive.Name, &petToRecive.LocalizationOng, &petToRecive.PetDetails, &petToRecive.SocialMediaOng)
	//if err != nil && err != sql.ErrNoRows {
	//	err = fmt.Errorf("error finding pet %d: %w", id, err)
	//	fmt.Println(err)
	//	return nil, err
	//}
	//pet = &petToRecive
	return
}
func (pr *PetRepository) Update(petID string, userID string, updatePayload map[string]interface{}) error {
	query := "UPDATE pets SET "
	values := []interface{}{}

	for key, value := range updatePayload {
		query += key + "=?, "
		values = append(values, value)
	}

	query = strings.TrimSuffix(query, ", ")
	query += " WHERE id=? AND userId=?"
	values = append(values, petID)
	values = append(values, userID)

	_, err := pr.dbconnection.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error updating pet: %w \\n", err)
	}

	return nil
}

func (pr *PetRepository) ListByUser(userID uniqueEntityId.ID) (pets []*entity.Pet, err error) {
	rows, err := pr.dbconnection.Query(`
		SELECT
		p.id,
		p.name,
		p.breedId,
		p.size,
		p.weight,
		p.adoptionDate,
		p.birthdate,
		p.comorbidity,
		p.tags,
		p.castrated,
		p.availableToAdoption,
		p.userId,
		b.name AS breed_name,
		pi.url AS pet_image_url
	FROM
		pets p
		JOIN breeds b ON p.breedId = b.id
		LEFT JOIN pets_image pi ON p.id = pi.petId
	WHERE
		p.userId = ?`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving pets for user %d: %w", userID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var pet entity.Pet
		var adoptionDateStr string
		var birthdateStr string

		if err = rows.Scan(
			&pet.ID,
			&pet.Name,
			&pet.BreedID,
			&pet.Size,
			&pet.Weight,
			&adoptionDateStr,
			&birthdateStr,
			&pet.Comorbidity,
			&pet.Tags,
			&pet.Castrated,
			&pet.AvailableToAdoption,
			&pet.UserID,
			&pet.BreedName,
			&pet.ImageUrl,
		); err != nil {
			return nil, fmt.Errorf("error scanning pet row: %w", err)
		}

		if pet.AdoptionDate, err = time.Parse(config.StandardDateLayout, adoptionDateStr); err != nil {
			return nil, fmt.Errorf("error parsing adoptionDate: %w", err)
		}
		if pet.Birthdate, err = time.Parse(config.StandardDateLayout, birthdateStr); err != nil {
			return nil, fmt.Errorf("error parsing birthdate: %w", err)
		}

		pets = append(pets, &pet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over pet rows: %w", err)
	}

	return pets, nil
}
