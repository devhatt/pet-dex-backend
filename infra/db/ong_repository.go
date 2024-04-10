package db

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OngRepository struct {
	dbconnection *sqlx.DB
}

// Delete implements interfaces.UserRepository.
func (o *OngRepository) Delete(id uuid.UUID) error {
	panic("unimplemented")
}

// FindByEmail implements interfaces.UserRepository.
func (o *OngRepository) FindByEmail(email string) *entity.User {
	panic("unimplemented")
}

// FindById implements interfaces.UserRepository.
func (o *OngRepository) FindById(id uuid.UUID) *entity.User {
	panic("unimplemented")
}

// List implements interfaces.UserRepository.
func (o *OngRepository) List() ([]entity.User, error) {
	panic("unimplemented")
}

// Save implements interfaces.UserRepository.
func (o *OngRepository) Save(user *entity.User) error {
	panic("unimplemented")
}

// SaveAddress implements interfaces.UserRepository.
func (o *OngRepository) SaveAddress(addr *entity.Address) error {
	panic("unimplemented")
}

// Update implements interfaces.UserRepository.
func (o *OngRepository) Update(userID uuid.UUID, user entity.User) error {
	panic("unimplemented")
}

func NewOngRepository(dbconn *sqlx.DB) interfaces.UserRepository {
	return &OngRepository{
		dbconnection: dbconn,
	}
}
