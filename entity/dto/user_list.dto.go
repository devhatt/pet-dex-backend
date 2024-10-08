package dto

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type UserListInput struct {
	Page   int
	Limit  int
	Search string
}

type UserList struct {
	ID                       uniqueEntityId.ID `json:"id" db:"id"`
	Name                     string            `json:"name" db:"name"`
	Type                     string            `json:"type" db:"type"`
	Document                 string            `json:"document" db:"document"`
	AvatarURL                string            `json:"avatar_url" db:"avatarUrl"`
	Email                    string            `json:"email" db:"email"`
	Phone                    string            `json:"phone" db:"phone"`
	Birthdate                *time.Time        `json:"birthdate"`
	PushNotificationsEnabled *bool             `json:"pushNotificationsEnabled" db:"pushNotificationsEnabled"`
}

type UserListOutput struct {
	Users []UserList `json:"users"`
	Total int        `json:"total"`
}

func NewUserListInput(page, limit int, search string) *UserListInput {
	return &UserListInput{
		Page:   page,
		Limit:  limit,
		Search: search,
	}
}
