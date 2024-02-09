package entity

type Pet struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	BreedID             string  `json:"breed_id"`
	Size                string  `json:"size"`
	Weight              float64 `json:"weight"`
	AdoptionDate        string  `json:"adoption_date"`
	Birthdate           string  `json:"birthdate"`
	Comorbidity         string  `json:"comorbidity"`
	Tags                string  `json:"tags"`
	Castrated           bool    `json:"castrated"`
	AvailableToAdoption bool    `json:"available_to_adoption"`
	UserID              string  `json:"user_id"`
}
