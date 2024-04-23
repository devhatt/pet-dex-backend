package dto

import (
	"pet-dex-backend/v2/entity"
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

func (dto *PetUpdatetDto) ToEntity() *entity.Pet {
	vaccines := make([]entity.Vaccines, len(dto.Vaccines))
	for i, v := range dto.Vaccines {
		vaccines[i] = entity.Vaccines{
			Name:      v.Name,
			Date:      v.Date,
			DoctorCRM: v.DoctorCRM,
		}
	}
	special_care := entity.SpecialCare{
		Needed:      dto.NeedSpecialCare.Needed,
		Description: dto.NeedSpecialCare.Description,
	}

	return &entity.Pet{
		Name:                dto.Name,
		Size:                dto.Size,
		Weight:              dto.Weight,
		AdoptionDate:        dto.AdoptionDate,
		Birthdate:           dto.Birthdate,
		Comorbidity:         dto.Comorbidity,
		Tags:                dto.Tags,
		Castrated:           dto.Castrated,
		AvailableToAdoption: dto.AvailableToAdoption,
		BreedID:             dto.BreedID,
		Vaccines:            vaccines,
		NeedSpecialCare:     special_care,
	}
}
