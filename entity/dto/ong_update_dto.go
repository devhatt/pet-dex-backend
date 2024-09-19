package dto

type OngUpdateDto struct {
	Phone          string `json:"phone" db:"phone" example:"119596995887"`
	User           UserUpdateDto
	OpeningHours   string    `json:"openingHours" example:"08:00"`
	AdoptionPolicy string    `json:"adoptionPolicy" example:"não pode rato"`
	Links          []LinkDto `json:"links"`
}
