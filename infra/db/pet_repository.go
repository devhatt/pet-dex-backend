package db

import (
	"database/sql"
	"fmt"
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

func (pr *PetRepository) Find(id int) (pet *entity.Pet, err error) {
	 err = pr.dbconnection.QueryRow("SELECT id, name, image, localization_ong, pet_details, social_media_ong FROM pet WHERE id = ?", id).Scan(&pet.Id, &pet.Name, &pet.Image, &pet.Localization_ong, &pet.Pet_details, &pet.Social_media_ong)
	if err != nil && err != sql.ErrNoRows {
		err = fmt.Errorf("error finding pet %d: %w", id, err)
		fmt.Println(err)
		return nil, err
	} 
	return 
}
