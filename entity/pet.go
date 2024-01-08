package entity

type Pet struct {
	Id int
	Name string
	Image string
	Localization_ong []string
	Pet_details Pet_details
	Social_media_ong []string
}
type Pet_details struct {
	Breed string
	Age int
	Size string
}