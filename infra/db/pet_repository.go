package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"time"

	"github.com/google/uuid"
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

func (pr *PetRepository) Save(petToSave *entity.Pet) error {
	_, err := pr.dbconnection.NamedExec("INSERT INTO pets (name, weight, size, adoptionDate, birthdate, breedId, userId) VALUES (:name, :weight, :size, :adoptionDate, :birthdate, :breedId, :userId)", &petToSave)

	if err != nil {
		err = fmt.Errorf("error saving pet: %w", err)
		fmt.Println(err)
		return err
	}
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
		p.neededSpecialCare,
		p.descriptionSpecialCare,
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
		&pet.NeedSpecialCare.Needed,
		&pet.NeedSpecialCare.Description,
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

	query := "UPDATE pets SET"
	values := []interface{}{}

	if petToUpdate.Name != "" {
		query = query + " name =?,"
		values = append(values, petToUpdate.Name)
	}

	if petToUpdate.BreedID != uuid.Nil {
		query = query + " breedId =?,"
		values = append(values, petToUpdate.BreedID)
	}

	if petToUpdate.Size != "" {
		query = query + " size =?,"
		values = append(values, petToUpdate.Size)
	}

	if petToUpdate.Weight != 0 {
		query = query + " weight =?,"
		values = append(values, petToUpdate.Weight)
	}

	if !petToUpdate.AdoptionDate.IsZero() {
		query = query + " adoptionDate = ?,"
		values = append(values, petToUpdate.AdoptionDate)
	}

	if !petToUpdate.Birthdate.IsZero() {
		query = query + " birthdate = ?,"
		values = append(values, petToUpdate.Birthdate)
	}

	if petToUpdate.Comorbidity != "" {
		query = query + " comorbidity = ?,"
		values = append(values, petToUpdate.Comorbidity)
	}

	if petToUpdate.Tags != "" {
		query = query + " tags = ?,"
		values = append(values, petToUpdate.Tags)
	}

	if petToUpdate.Castrated != nil {
		query = query + " castrated = ?,"
		values = append(values, petToUpdate.Castrated)
	}

	if petToUpdate.AvailableToAdoption != nil {
		query = query + " availableToAdoption = ?,"
		values = append(values, petToUpdate.AvailableToAdoption)
	}

	if petToUpdate.UserID != uuid.Nil {
		query = query + " userId = ?,"
		values = append(values, petToUpdate.UserID)
	}

	if petToUpdate.NeedSpecialCare.Needed != nil {
		query = query + " needed = ?,"
		values = append(values, petToUpdate.NeedSpecialCare.Needed)
		query = query + " description = ?,"
		values = append(values, petToUpdate.NeedSpecialCare.Description)
	}

	n := len(query)
	query = query[:n-1] + " WHERE id =?"
	values = append(values, petID)

	_, err := pr.dbconnection.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error updating pet: %w \\n", err)
	}

	_, err = pr.dbconnection.Exec("DELETE FROM vaccines WHERE petId = ?", petID)
	if err != nil {
		return fmt.Errorf("error removing existing vaccines: %w", err)
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
			Needed:      &needed,
			Description: description,
		}

		pets = append(pets, &pet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over pet rows: %w", err)
	}

	return pets, nil
}

func (pr *PetRepository) ListAllByPage(page int) (pets []*entity.Pet, err error) {
	offset := (page - 1) * 12
	rows, err := pr.dbconnection.Query(`
	SELECT
		p.id,
		p.name,
		p.breedId,
		p.birthdate,
		p.availableToAdoption,
		b.name AS breed_name,
		pi.url AS pet_image_url
	FROM
		pets p
		JOIN breeds b ON p.breedId = b.id
		LEFT JOIN pets_image pi ON p.id = pi.petId
	LIMIT 12 
	OFFSET ?`, offset)

	if err != nil {
		return nil, fmt.Errorf("error retrieving pets: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var pet entity.Pet
		var birthdateStr string

		if err = rows.Scan(
			&pet.ID,
			&pet.Name,
			&pet.BreedID,
			&birthdateStr,
			&pet.AvailableToAdoption,
			&pet.BreedName,
			&pet.ImageUrl,
		); err != nil {
			return nil, fmt.Errorf("error scanning pet row: %w", err)
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
