package dto

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type PetUpdateDto struct {
	Name                string            `json:"name" example:"Spike"`
	Size                string            `json:"size" example:"small"`
	Weight              float64           `json:"weight" example:"4.8"`
	WeightMeasure       string            `json:"weight_measure" example:"kg"`
	AdoptionDate        time.Time         `json:"adoption_date" example:"2008-01-02T00:00:00Z"`
	Birthdate           time.Time         `json:"birthdate" example:"2006-01-02T00:00:00Z"`
	Comorbidity         string            `json:"comorbidity" example:"asma"`
	Tags                string            `json:"tags" example:"Dog"`
	Castrated           *bool             `json:"castrated" example:"true"`
	AvailableToAdoption *bool             `json:"available_to_adoption" example:"true"`
	BreedID             uniqueEntityId.ID `json:"breed_id" example:"0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"`
	Vaccines            []VaccinesDto     `json:"vaccines"`
	NeedSpecialCare     SpecialCareDto    `json:"special_care"`
}

type VaccinesDto struct {
	Name      string    `json:"name" example:"PetVax"`
	Date      time.Time `json:"date" example:"2007-01-02T00:00:00Z"`
	DoctorCRM string    `json:"doctor_crm" example:"000000"`
}

type SpecialCareDto struct {
	Needed      *bool  `json:"neededSpecialCare" example:"true"`
	Description string `json:"descriptionSpecialCare" example:"obesity"`
}
