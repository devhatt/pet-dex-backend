package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type UserRepository interface {
	Save(user *entity.User) error
	Update(userID uniqueEntityId.ID, user entity.User) error
	Delete(id uniqueEntityId.ID) error
	FindByID(ID uniqueEntityId.ID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	List() ([]entity.User, error)
	AdressRepo
}
