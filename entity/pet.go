package entity

import "pet-dex-backend/v2/pkg/entity"

type Pet struct {
	ID                  entity.ID `json:"id"`
	Name                string    `json:"name"`
	BreedID             entity.ID `json:"breed_id"`
	Size                string    `json:"size"`
	Weight              float64   `json:"weight"`
	AdoptionDate        string    `json:"adoption_date"`
	Birthdate           string    `json:"birthdate"`
	Comorbidity         string    `json:"comorbidity"`
	Tags                string    `json:"tags"`
	Castrated           bool      `json:"castrated"`
	AvailableToAdoption bool      `json:"available_to_adoption"`
	UserID              entity.ID `json:"user_id"`
}
