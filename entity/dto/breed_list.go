package dto

import (
	uniqueEntity "pet-dex-backend/v2/pkg/entity"
)

type BreedList struct {
	ID     uniqueEntity.ID `json:"id"`
	Name   string          `json:"name"`
	ImgUrl string          `json:"img_url"`
}

func (breed *BreedList) Validate() bool {
	return (breed.Name != "" &&
		breed.ImgUrl != "")
}
