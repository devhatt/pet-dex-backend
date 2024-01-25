package entity

type Ong struct {
	CNPJ        string      `json:"cnpj"`
	Email       string      `json:"email"`
	Location    Location    `json:"location"`
	Imagem      string      `json:"image"`
	SocialMedia SocialMedia `json:"socialmedia"`
	UserID      string      `json:"userid"`
}

type Location struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type SocialMedia struct {
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
}
