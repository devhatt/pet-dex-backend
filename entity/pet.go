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
	Castrated           *bool             `json:"castrated"`
	AvailableToAdoption *bool             `json:"available_to_adoption"`
	BreedName           string            `json:"breed_name"`
	ImageUrl            string            `json:"image_url"`
	Vaccines            []Vaccines        `json:"vaccines"`
	NeedSpecialCare     SpecialCare       `json:"special_care"`
}

type Vaccines struct {
	ID        uniqueEntityId.ID `json:"id"`
	PetID     uniqueEntityId.ID `json:"pet_id"`
	Name      string            `json:"name"`
	Date      time.Time         `json:"date"`
	DoctorCRM string            `json:"doctor_crm"`
}
type SpecialCare struct {
	Needed      *bool  `json:"needed"`
	Description string `json:"description"`
}

type PetDetails struct {
	Breed string
	Age   int
	Size  string
}
