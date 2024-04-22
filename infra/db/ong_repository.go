package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"

	"github.com/jmoiron/sqlx"
)

type OngRepository struct {
	dbconnection *sqlx.DB
	logger       config.Logger
}

func NewOngRepository(db *sqlx.DB) interfaces.OngRepository {
	return &OngRepository{
		dbconnection: db,
		logger:       *config.GetLogger("ong-repository"),
	}
}

func (or *OngRepository) Save(ong *entity.Ong) error {
	_, err := or.dbconnection.NamedExec("INSERT INTO legal_persons (id, userId, phone, links, openingHours, adoptionPolicy) VALUES (:id, :userId, :phone, :links, :openingHours, :adoptionPolicy)", &ong)

	if err != nil {
		logger.Error("error on Save: ", err)
		err = fmt.Errorf("error on saving ong")
		return err
	}

	return nil
}

func (or *OngRepository) SaveUser(user *entity.User) error {
	_, err := or.dbconnection.NamedExec("INSERT INTO users (id, name, type, document, avatarUrl, email, phone, pass) VALUES (:id, :name, :type, :document, :avatarUrl, :email, :phone, :pass)", &user)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngRepository.SaveUser error: %w", err))
		err = fmt.Errorf("error on saving user")
		return err
	}

	return nil
}

func (or *OngRepository) SaveAddress(addr *entity.Address) error {
	_, err := or.dbconnection.NamedExec("INSERT INTO addresses (id, userId, address, city, state, latitude, longitude) VALUES (:id, :userId, :address, :city, :state, :latitude, :longitude)", &addr)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngRepository.SaveAddress error: %w", err))
		err = fmt.Errorf("error on saving address")
		return err
	}

	return nil
}
