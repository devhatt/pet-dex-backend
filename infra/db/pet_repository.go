package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
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
	err = pr.dbconnection.QueryRow("SELECT id, name, localization_ong, pet_details, social_media_ong FROM pet WHERE id = ?", id).Scan(&petToRecive.Id, &petToRecive.Name, &petToRecive.LocalizationOng, &petToRecive.PetDetails, &petToRecive.SocialMediaOng)
	if err != nil && err != sql.ErrNoRows {
		err = fmt.Errorf("error finding pet %d: %w", id, err)
		fmt.Println(err)
		return nil, err
	}
	pet = &petToRecive
	return
}

func (pr *PetRepository) FindNoAuthPets() (*os.File, error) {
	// var allPets []entity.PetNoAuth
	res, err := os.Open("cats.json")
	if err != nil {
		fmt.Println()
		return nil, errors.New("Pet Crowler not found")
	}

	return res, nil
}
