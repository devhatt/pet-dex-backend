package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"time"

	"github.com/jmoiron/sqlx"

	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type PetRepository struct {
	dbconnection *sqlx.DB
}

func NewPetRepository(dbconn *sqlx.DB) interfaces.PetRepository {
	return &PetRepository{
		dbconnection: dbconn,
	}
}

func (pr *PetRepository) Save(entity.Pet) error {
	return nil
}

func (pr *PetRepository) FindByID(ID uniqueEntityId.ID) (*entity.Pet, error) {
	row, err := pr.dbconnection.Query(`
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
		p.needed,
		p.description,
        b.name AS breed_name,
        pi.url AS pet_image_url
    FROM
        pets p
        JOIN breeds b ON p.breedId = b.id
        LEFT JOIN pets_image pi ON p.id = pi.petId
    WHERE
        p.id = ?`,
		ID,
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving pet %d: %w", ID, err)
	}
	defer row.Close()

	if !row.Next() {
		return nil, sql.ErrNoRows
	}

	var pet entity.Pet
	var adoptionDateStr string
	var birthdateStr string

	if err := row.Scan(
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
		return nil, fmt.Errorf("error scanning pet: %w", err)
	}

	if pet.AdoptionDate, err = time.Parse(config.StandardDateLayout, adoptionDateStr); err != nil {
		return nil, fmt.Errorf("error parsing adoptionDate: %w", err)
	}
	if pet.Birthdate, err = time.Parse(config.StandardDateLayout, birthdateStr); err != nil {
		return nil, fmt.Errorf("error parsing birthdate: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over pet rows: %w", err)
	}

	return &pet, nil
}
func (pr *PetRepository) Update(petID string, userID string, petToUpdate *entity.Pet) error {

	query := "UPDATE pets SET name=?, size=?, weight=?, adoptionDate=?, birthdate=?, comorbidity=?, tags=?, castrated=?, availableToAdoption=?, breedId=?, description=?, needed=? WHERE id=?"
	values := []interface{}{petToUpdate.Name, petToUpdate.Size, petToUpdate.Weight, petToUpdate.AdoptionDate, petToUpdate.Birthdate,
		petToUpdate.Comorbidity, petToUpdate.Tags, petToUpdate.Castrated, petToUpdate.AvailableToAdoption, petToUpdate.BreedID,
		petToUpdate.NeedSpecialCare.Description, petToUpdate.NeedSpecialCare.Needed, petID,
	}

	_, err := pr.dbconnection.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error updating pet: %w \\n", err)
	}

	_, errDel := pr.dbconnection.Exec("DELETE FROM vaccines WHERE petId = ?", petID)
	if err != nil {
		return fmt.Errorf("error removing existing vaccines: %w", errDel)
	}

	for _, vaccine := range petToUpdate.Vaccines {
		_, err := pr.dbconnection.Exec(
			"INSERT INTO vaccines (id, petId, name, date, doctorCRM) VALUES (?, ?, ?, ?, ?)",
			uniqueEntityId.NewID(), petID, vaccine.Name, vaccine.Date, vaccine.DoctorCRM,
		)
		if err != nil {
			return fmt.Errorf("error adding new vaccine: %w", err)
		}
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
		p.needed,
		p.description,
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
		var needed bool
		var description string

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
			&needed,
			&description,
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
		pet.NeedSpecialCare = entity.SpecialCare{
			Needed:      needed,
			Description: description,
		}

		pets = append(pets, &pet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over pet rows: %w", err)
	}

	return pets, nil
}
