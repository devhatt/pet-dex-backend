package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"

	"github.com/jmoiron/sqlx"
)

var logger = config.GetLogger("user-repository")

type UserRepository struct {
	dbconnection *sqlx.DB
}

func NewUserRepository(dbconn *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		dbconnection: dbconn,
	}
}

func (ur *UserRepository) Delete(id uniqueEntityId.ID) error {
	return nil
}

func (ur *UserRepository) Save(user *entity.User) error {
	_, err := ur.dbconnection.NamedExec("INSERT INTO users (id, name, type, document, avatarUrl, email, phone, pass) VALUES (:id, :name, :type, :document, :avatarUrl, :email, :phone, :pass)", &user)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserRepository.Save error: %w", err))
		err = fmt.Errorf("error on saving user")
		return err
	}

	return nil
}

func (ur *UserRepository) SaveAddress(addr *entity.Address) error {
	_, err := ur.dbconnection.NamedExec("INSERT INTO addresses (id, userId, address, city, state, latitude, longitute) VALUES (:id, :userId, :address, :city, :state, :latitude, :longitute)", &addr)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserRepository.SaveAddress error: %w", err))
		err = fmt.Errorf("error on saving address")
		return err
	}

	return nil
}

func (ur *UserRepository) Update(userID uniqueEntityId.ID, user entity.User) error {
	_, err := ur.dbconnection.NamedExec("UPDATE users SET name=:name, document=:document, avatarURL=:avatarURL, email=:email, phone=:phone WHERE id=:id", &user)

	if err != nil {
		logger.Error(fmt.Errorf("#UserRepository.Update error: %w", err))
		return fmt.Errorf("error on update user, %w", err)
	}

	return nil
}

func (ur *UserRepository) FindById(id uniqueEntityId.ID) *entity.User {

	return &entity.User{}
}

func (ur *UserRepository) FindByEmail(email string) *entity.User {
	return &entity.User{}
}

func (ur *UserRepository) List() (users []entity.User, err error) {
	return nil, nil
}
