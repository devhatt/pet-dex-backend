package entity

type Pet struct {
	Breed       string
	ID          string `json:"id"`
	Aniversario string `json:"aniversario"`
	DiaDoacao   string `json:"dia_doacao"`
}
