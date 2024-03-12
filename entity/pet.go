package entity

import (
	"time"

	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type Pet struct {
	ID                  uniqueEntityId.ID `json:"id"`
	UserID              uniqueEntityId.ID `json:"user_id"`
	BreedID             uniqueEntityId.ID `json:"breed_id"`
	Name                string            `json:"name"`
	Size                string            `json:"size"`
	Weight              float64           `json:"weight"`
	AdoptionDate        time.Time         `json:"adoption_date"`
	Birthdate           time.Time         `json:"birthdate"`
	Comorbidity         string            `json:"comorbidity"`
	Tags                string            `json:"tags"`
	Castrated           bool              `json:"castrated"`
	AvailableToAdoption bool              `json:"available_to_adoption"`
	BreedName           string            `json:"breed_name"`
	ImageUrl            string            `json:"image_url"`
}
