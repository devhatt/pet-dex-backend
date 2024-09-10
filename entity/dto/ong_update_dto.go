package dto

type Link struct {
	URL   string  `json:"url" example:"https://www.facebook.com/"`
	Text  string  `json:"text" example:"Facebook da Ong"`
}

type OngUpdateDto struct {
	Phone          string `json:"phone" db:"phone" example:"119596995887"`
	User           UserUpdateDto
	OpeningHours   string           `json:"openingHours" example:"08:00"`
	AdoptionPolicy string           `json:"adoptionPolicy" example:"não pode rato"`
	Links          []Link 					`json:"links"`
}

