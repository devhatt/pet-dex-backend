package entity

import "pet-dex-backend/v2/pkg/uniqueEntityId"

type Breed struct {
	ID             uniqueEntityId.ID `json:"id"`
	Name           string            `json:"name"`
	Specie         string            `json:"specie"`
	Size           string            `json:"size"`
	Description    string            `json:"description"`
	Height         string            `json:"height"`
	Weight         string            `json:"weight"`
	PhysicalChar   string            `json:"physical_char"`
	Disposition    string            `json:"disposition"`
	IdealFor       string            `json:"ideal_for"`
	Fur            string            `json:"fur"`
	ImgUrl         string            `json:"img_url"`
	Weather        string            `json:"weather"`
	Dressage       string            `json:"dressage"`
	OrgID          string            `json:"org_id"`
	LifeExpectancy string            `json:"life_expectancy"`
}
