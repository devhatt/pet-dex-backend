package entity

type Pet struct {
	Id              int
	Name            string
	Image           string
	LocalizationOng string
	SocialMediaOng  string
	PetDetails
}

type PetDetails struct {
	Age   int
	Size  string
	Breed string
}

type PetNoAuth struct {
	Id              int
	Breed           string `json:"nome"`
	Image           string `json:"imgUrl"`
	Description     string `json:"descricao"`
	Size            string `json:"tamanho"`
	Weight          string `json:"peso"`
	LifeExpectation string `json:"esperancaDeVida"`
}
