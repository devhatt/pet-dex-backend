package db

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"

	"github.com/jmoiron/sqlx"
)

var loggerUserRepository = config.GetLogger("user-repository")

type UserRepository struct {
	dbconnection *sqlx.DB
	logger       config.Logger
}

func NewUserRepository(dbconn *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		dbconnection: dbconn,
		logger:       *config.GetLogger("ong-repository"),
	}
}

func (ur *UserRepository) Delete(id uniqueEntityId.ID) error {
	return nil
}

func (ur *UserRepository) Save(user *entity.User) error {
	_, err := ur.dbconnection.NamedExec("INSERT INTO users (id, name, type, document, avatarUrl, email, phone, pass) VALUES (:id, :name, :type, :document, :avatarUrl, :email, :phone, :pass)", &user)

	if err != nil {
		ur.logger.Error("error saving user: ", err)
		return err
	}

	return nil
}

func (ur *UserRepository) SaveAddress(addr *entity.Address) error {
	_, err := ur.dbconnection.NamedExec("INSERT INTO addresses (id, userId, address, city, state, latitude, longitude) VALUES (:id, :userId, :address, :city, :state, :latitude, :longitude)", &addr)

	if err != nil {
		ur.logger.Error("error saving address: ", err)
		return err
	}

	return nil
}

func (ur *UserRepository) FindAddressByUserID(userID uniqueEntityId.ID) (*entity.Address, error) {
	row, err := ur.dbconnection.Query(`
	SELECT
		a.id,
		a.address,
		a.city,
		a.state,
		a.latitude,
		a.longitude,
	FROM
		addresses a
	WHERE
		a.userId = ?`,
		userID,
	)
	if err != nil {
		ur.logger.Error("error saving address: ", err)
		err = fmt.Errorf("error retrieving address %d: %w", userID, err)
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, errors.New("sql: no rows in result")
	}

	var address entity.Address

	err = row.Scan(
		&address.ID,
		&address.Address,
		&address.City,
		&address.State,
		&address.Latitude,
		&address.Longitude,
	)
	if err != nil {
		ur.logger.Error("error on user repository: ", err)
		err = fmt.Errorf("error scanning address %d: %w", userID, err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		ur.logger.Error("error on user repository: ", err)
		return nil, fmt.Errorf("error iterating over address rows: %w", err)
	}

	return &address, nil
}

func (ur *UserRepository) Update(userID uniqueEntityId.ID, userToUpdate entity.User) error {

	query := "UPDATE users SET"
	values := []interface{}{}

	if userToUpdate.Name != "" {
		query = query + " name =?,"
		values = append(values, userToUpdate.Name)
	}

	if userToUpdate.Document != "" {
		query = query + " document =?,"
		values = append(values, userToUpdate.Document)
	}

	if userToUpdate.AvatarURL != "" {
		query = query + " avatarUrl =?,"
		values = append(values, userToUpdate.AvatarURL)
	}

	if userToUpdate.Email != "" {
		query = query + " email =?,"
		values = append(values, userToUpdate.Email)
	}

	if userToUpdate.Phone != "" {
		query = query + " phone =?,"
		values = append(values, userToUpdate.Phone)
	}

	if userToUpdate.BirthDate != nil {
		query = query + " birthdate =?,"
		values = append(values, userToUpdate.BirthDate)
	}

	query = query + " updated_at =?,"
	values = append(values, time.Now())

	n := len(query)
	query = query[:n-1] + " WHERE id =?"
	values = append(values, userID)

	fmt.Printf("Query to update: %s", query)

	_, err := ur.dbconnection.Exec(query, values...)

	if err != nil {
		ur.logger.Error(fmt.Errorf("#UserRepository.Update error: %w", err))
		return fmt.Errorf("error on update user")
	}

	return nil
}

func (ur *UserRepository) FindByID(ID uniqueEntityId.ID) (*entity.User, error) {
	row, err := ur.dbconnection.Query(`
	SELECT
		u.id,
		u.name,
		u.birthdate,
		u.document,
		u.avatarUrl,
		u.email,
		u.phone
	FROM
		users u
	WHERE
		u.id = ?`,
		ID,
	)
	if err != nil {
		ur.logger.Error("error saving address: ", err)
		err = fmt.Errorf("error retrieving user %d: %w", ID, err)
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, errors.New("sql: no rows in result")
	}

	var user entity.User
	var birthdateStr string

	err = row.Scan(
		&user.ID,
		&user.Name,
		&birthdateStr,
		&user.Document,
		&user.AvatarURL,
		&user.Email,
		&user.Phone,
	)
	if err != nil {
		ur.logger.Error("error saving address: ", err)
		err = fmt.Errorf("error scanning user %d: %w", ID, err)
		return nil, err
	}

	if *user.BirthDate, err = time.Parse(config.StandardDateLayout, birthdateStr); err != nil {
		return nil, fmt.Errorf("error parsing adoptionDate: %w", err)
	}

	if err := row.Err(); err != nil {
		ur.logger.Error("error saving address: ", err)
		return nil, fmt.Errorf("error iterating over user rows: %w", err)
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) *entity.User {
	return &entity.User{}
}

func (ur *UserRepository) List() (users []entity.User, err error) {
	return nil, nil
}
