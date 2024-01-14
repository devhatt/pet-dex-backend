package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	"strings"
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

// Incomplete method
func (pr *PetRepository) ListUserPets(userID int) (pets []*entity.Pet, err error) {
	// I am using * in the query because I don't know what needs to be fetched from the table
	rows, err := pr.dbconnection.Query("SELECT * FROM pet WHERE userID = ?", userID)
	if err != nil && err != sql.ErrNoRows {
		err = fmt.Errorf("error finding pets from user %d: %w", userID, err)
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(rows)
	return pets, err
}
