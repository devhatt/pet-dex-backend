package entity

type Pet struct {
	Id int
	Name string
	Image string
	LocalizationOng string
	PetDetails string
	SocialMediaOng string
}
type PetDetails struct {
	Breed string
	Age int
	Size string
}