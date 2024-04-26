package dto

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type PetUpdatetDto struct {
	Name                string            `json:"name"`
	Size                string            `json:"size"`
	Weight              float64           `json:"weight"`
	AdoptionDate        time.Time         `json:"adoption_date"`
	Birthdate           time.Time         `json:"birthdate"`
	Comorbidity         string            `json:"comorbidity"`
	Tags                string            `json:"tags"`
	Castrated           *bool             `json:"castrated"`
	AvailableToAdoption *bool             `json:"available_to_adoption"`
	BreedID             uniqueEntityId.ID `json:"breed_id"`
	Vaccines            []VaccinesDto     `json:"vaccines"`
	NeedSpecialCare     SpecialCareDto    `json:"special_care"`
}

type VaccinesDto struct {
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	DoctorCRM string    `json:"doctor_crm"`
}

type SpecialCareDto struct {
	Needed      *bool  `json:"needed"`
	Description string `json:"description"`
}
