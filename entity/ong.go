package entity

type Ong struct {
	CNPJ         string      `json:"cnpj"`
	Email        string      `json:"email"`
	Localizacao  Location    `json:"localizacao"`
	Imagem       string      `json:"imagem"`
	RedesSociais SocialMedia `json:"redes_sociais"`
}

type Location struct {
	Endereco string `json:"endereco"`
	Cidade   string `json:"cidade"`
	Estado   string `json:"estado"`
}

type SocialMedia struct {
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
}
