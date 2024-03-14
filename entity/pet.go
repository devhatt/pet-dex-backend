package entity

import (
	"time"

	uniqueEntity "pet-dex-backend/v2/pkg/entity"
)

type Pet struct {
	ID                  uniqueEntity.ID `json:"id"`
	UserID              uniqueEntity.ID `json:"user_id"`
	BreedID             uniqueEntity.ID `json:"breed_id"`
	Name                string          `json:"name"`
	Size                string          `json:"size"`
	Weight              float64         `json:"weight"`
	WeightMeasure       string          `json:"weight_measure"`
	AdoptionDate        time.Time       `json:"adoption_date"`
	Birthdate           time.Time       `json:"birthdate"`
	Comorbidity         string          `json:"comorbidity"`
	Tags                string          `json:"tags"`
	Castrated           bool            `json:"castrated"`
	AvailableToAdoption bool            `json:"available_to_adoption"`
	BreedName           string          `json:"breed_name"`
	ImageUrl            string          `json:"image_url"`
}
