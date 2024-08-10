package dto

import (
	"errors"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"regexp"
	"time"
)

var notEmptyRegex = regexp.MustCompile(`^\S+$`)

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

var sizeRegex = regexp.MustCompile(`^(small|medium|large|giant)$`)

type PetInsertDto struct {
	Name         string            `json:"name" example:"Thor"`
	UserID       uniqueEntityId.ID `json:"user_id" example:"fa1b8ae8-5351-11ef-8f02-0242ac130003"`
	BreedID      uniqueEntityId.ID `json:"breed_id" example:"0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"`
	AdoptionDate *time.Time        `json:"adoption_date" example:"2008-01-02T15:04:05Z"`
	Birthdate    *time.Time        `json:"birthdate" example:"2006-01-02T15:04:05Z"`
	Weight       float64           `json:"weight" example:"4.1"`
	Size         string            `json:"size" example:"medium"`
}

func (p *PetInsertDto) Validate() error {
	if !nameRegex.MatchString(p.Name) {
		return errors.New("name cannot be empty")
	}
	if len(p.Name) > 80 {
		return errors.New("name cannot exceed 80 characters")
	}
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	if !sizeRegex.MatchString(p.Size) {
		return errors.New("size can be only small, medium, large or giant")
	}
	if !notEmptyRegex.MatchString(p.UserID.String()) {
		return errors.New("UserID cannot be empty")
	}
	if !notEmptyRegex.MatchString(p.BreedID.String()) {
		return errors.New("BreedID cannot be empty")
	}
	if p.Weight < 0 {
		return errors.New("weight cannot be negative")
	}
	return nil
}
