package dto

import "pet-dex-backend/v2/pkg/uniqueEntityId"

type BreedList struct {
	ID     uniqueEntityId.ID `json:"id" example:"0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"`
	Name   string            `json:"name" example:"Pastor Alemão"`
	ImgUrl string            `json:"img_url" example:"https://images.unsplash.com/photo-1530281700549-e82e7bf110d6?q=80&w=1888&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"`
}

func (breed *BreedList) Validate() bool {
	return (breed.Name != "" &&
		breed.ImgUrl != "")
}
