package entity

import (
	"time"

	"github.com/google/uuid"
)

type Pet struct {
	ID                  uuid.UUID `json:"id"`
	UserID              uuid.UUID `json:"user_id"`
	BreedID             uuid.UUID `json:"breed_id"`
	Name                string    `json:"name"`
	Size                string    `json:"size"`
	Weight              float64   `json:"weight"`
	AdoptionDate        time.Time `json:"adoption_date"`
	Birthdate           time.Time `json:"birthdate"`
	Comorbidity         string    `json:"comorbidity"`
	Tags                string    `json:"tags"`
	Castrated           bool      `json:"castrated"`
	AvailableToAdoption bool      `json:"available_to_adoption"`
	BreedName           string    `json:"breed_name"`
	ImageUrl            string    `json:"image_url,omitempty"`
}
