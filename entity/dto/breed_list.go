package dto

import "pet-dex-backend/v2/pkg/uniqueEntityId"

type BreedList struct {
	ID     uniqueEntityId.ID `json:"id"`
	Name   string            `json:"name"`
	ImgUrl string            `json:"img_url"`
}

func (breed *BreedList) Validate() bool {
	return (breed.Name != "" &&
		breed.ImgUrl != "")
}
