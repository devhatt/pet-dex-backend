package entity

import (
	"time"

	"pet-dex-backend/v2/entity/dto"
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

func NewPet(userId, breedId uniqueEntityId.ID, size, name string, weight float64, adoptionDate, birthdate *time.Time) *Pet {
	petId := uniqueEntityId.NewID()

	return &Pet{
		ID:           petId,
		UserID:       userId,
		Name:         name,
		Weight:       weight,
		AdoptionDate: *adoptionDate,
		Birthdate:    *birthdate,
	}
}

func ToEntity(dto *dto.PetUpdatetDto) *Pet {
	vaccines := make([]Vaccines, len(dto.Vaccines))
	for i, v := range dto.Vaccines {
		vaccines[i] = Vaccines{
			Name:      v.Name,
			Date:      v.Date,
			DoctorCRM: v.DoctorCRM,
		}
	}
	special_care := SpecialCare{
		Needed:      dto.NeedSpecialCare.Needed,
		Description: dto.NeedSpecialCare.Description,
	}

	return &Pet{
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
