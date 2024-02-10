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
	var petToRecive entity.Pet
	// err = pr.dbconnection.QueryRow("SELECT * FROM pets WHERE id = ?", id).Scan(&petToRecive.Id, &petToRecive.Name, &petToRecive.LocalizationOng, &petToRecive.PetDetails, &petToRecive.SocialMediaOng)
	if err != nil && err != sql.ErrNoRows {
		err = fmt.Errorf("error finding pet %d: %w", id, err)
		fmt.Println(err)
		return nil, err
	}
	pet = &petToRecive
	return
}

func (pr *PetRepository) Update(userID string, petID string, updatePayload map[string]interface{}) error {
	query := "UPDATE petdex.pets SET "
	values := []interface{}{}

	for key, value := range updatePayload {
		query += key + " = ?, "
		values = append(values, value)
	}

	query = strings.TrimSuffix(query, ", ")

	query += " WHERE id = ? and userId = ?"
	values = append(values, petID)
	values = append(values, userID)

	result, err := pr.dbconnection.Exec(query, values...)
	if err != nil {
		fmt.Printf("[err] Error updating pet %s: %v \n", petID, err)
		return fmt.Errorf("Error updating pet %s: %w \n", petID, err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Erro when try get rows affected %s: %v \n", petID, err)
		return fmt.Errorf("Erro when try get rows affected %s: %w \n", petID, err)
	}

	if affectedRows == 0 {
		fmt.Printf("No pets were updated %s: %v \n", petID, err)
		return fmt.Errorf("No pets were updated %s: %w \n", petID, err)
	}

	return nil
}
