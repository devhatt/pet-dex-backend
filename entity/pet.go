package entity

type Pet struct {
	Id              int        `json:"id"`
	Name            string     `json:"nome"`
	Image           string     `json:"image"`
	LocalizationOng []string   `json:"localizationOng"`
	PetDetails      PetDetails `json:"details"`
	SocialMediaOng  []string   `json:"socialmediaOng"`
}
type PetDetails struct {
	Breed string `json:"bree"`
	Age   int    `json:"age"`
	Size  string `json:"size"`
}