package interfaces

import "pet-dex-backend/v2/entity"

type Emailrepository interface {
	SendConfirmationEmail(user *entity.User) error
	SendNotificationEmail(message string, recipient string) error
}
