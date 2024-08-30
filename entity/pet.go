package entity

import (
	"time"

	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type Pet struct {
	ID                  uniqueEntityId.ID `json:"id"`
	UserID              uniqueEntityId.ID `json:"user_id" db:"userId"`
	BreedID             uniqueEntityId.ID `json:"breed_id" db:"breedId"`
	Name                string            `json:"name"`
	Size                string            `json:"size"`
	Weight              float64           `json:"weight"`
	WeightMeasure       string            `json:"weight_measure"`
	AdoptionDate        time.Time         `json:"adoption_date" db:"adoptionDate"`
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
	Needed      *bool  `json:"neededSpecialCare"`
	Description string `json:"descriptionSpecialCare"`
}

func NewPet(userId, breedId uniqueEntityId.ID, size, name string, weight float64, adoptionDate, birthdate *time.Time) *Pet {
	petId := uniqueEntityId.NewID()

	return &Pet{
		ID:           petId,
		UserID:       userId,
		BreedID:      breedId,
		Size:         size,
		Name:         name,
		Weight:       weight,
		AdoptionDate: *adoptionDate,
		Birthdate:    *birthdate,
	}
}

func PetToEntity(dto *dto.PetUpdateDto) *Pet {
	vaccines := make([]Vaccines, len(dto.Vaccines))
	for i, v := range dto.Vaccines {
		vaccines[i] = Vaccines{
			Name:      v.Name,
			Date:      v.Date,
			DoctorCRM: v.DoctorCRM,
		}
	}
	specialCare := SpecialCare{
		Needed:      dto.NeedSpecialCare.Needed,
		Description: dto.NeedSpecialCare.Description,
	}

	return &Pet{
		Name:                dto.Name,
		Size:                dto.Size,
		Weight:              dto.Weight,
		WeightMeasure:       dto.WeightMeasure,
		AdoptionDate:        dto.AdoptionDate,
		Birthdate:           dto.Birthdate,
		Comorbidity:         dto.Comorbidity,
		Tags:                dto.Tags,
		Castrated:           dto.Castrated,
		AvailableToAdoption: dto.AvailableToAdoption,
		BreedID:             dto.BreedID,
		Vaccines:            vaccines,
		NeedSpecialCare:     specialCare,
	}
}
