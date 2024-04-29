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
	Name         string            `json:"name"`
	UserID       uniqueEntityId.ID `json:"user_id"`
	BreedID      uniqueEntityId.ID `json:"breed_id"`
	AdoptionDate *time.Time        `json:"adoption_date"`
	Birthdate    *time.Time        `json:"birthdate"`
	Weight       float64           `json:"weight"`
	Size         string            `json:"size"`
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
