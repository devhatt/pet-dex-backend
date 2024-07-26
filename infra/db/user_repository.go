package db

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"

	"github.com/jmoiron/sqlx"
)

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
	_, err := ur.dbconnection.Exec("UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?", id)

	if err != nil {
		ur.logger.Error("#UserRepository.Delete error: %w", err)
		return fmt.Errorf("error on delete user")
	}

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
	var address entity.Address

	err := ur.dbconnection.Get(&address,
		`SELECT
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
		ur.logger.Error("error retrieving address: ", err)
		err = fmt.Errorf("error retrieving address %d: %w", userID, err)
		return nil, err
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

	if userToUpdate.PushNotificationsEnabled != nil {
		query = query + " pushNotificationsEnabled =?,"
		values = append(values, userToUpdate.PushNotificationsEnabled)
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
	var user entity.User

	err := ur.dbconnection.Get(&user,
		`SELECT
		u.id,
		u.name,
		u.birthdate,
		u.document,
		u.avatarUrl,
		u.email,
		u.phone,
		u.pass
	FROM
		users u
	WHERE
		u.id = ?`,
		ID,
	)
	if err != nil {
		ur.logger.Error("error retrieving user: ", err)
		err = fmt.Errorf("error retrieving user %d: %w", ID, err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := ur.dbconnection.Get(&user,
		`SELECT
		u.id,
		u.name,
		u.email,
		u.pass
	FROM
		users u
	WHERE
		u.email = ?`,
		email,
	)
	if err != nil {
		ur.logger.Error("error retrieving user: ", err)
		err = fmt.Errorf("error retrieving user %s: %w", email, err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) ChangePassword(userId uniqueEntityId.ID, newPassword string) error {

	query := "UPDATE users SET pass = ?,"
	var values []interface{}

	values = append(values, newPassword)

	query = query + " updated_at =?,"
	values = append(values, time.Now())

	n := len(query)
	query = query[:n-1] + " WHERE id =?"
	values = append(values, userId)

	fmt.Printf("Query to update: %s", query)

	_, err := ur.dbconnection.Exec(query, values...)

	if err != nil {
		ur.logger.Error(fmt.Errorf("#UserRepository.ChangePassword error: %w", err))
		return fmt.Errorf("error on changing user password")
	}

	return nil
}


func (ur *UserRepository) List(input *dto.UserListInput) (output *dto.UserListOutput, err error) {
	var users []dto.UserList
	query := `SELECT 
				u.id, 
				u.name, 
				u.type, 
				u.document, 
				u.avatar_url, 
				u.email, 
				u.phone, 
				u.birthdate, 
				u.pushNotificationsEnabled, 
			  FROM users u`
	if input.Search != "" {
		query += " WHERE u.name LIKE ?"
	}

	offset := (input.Page - 1) * input.Limit
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", input.Limit, offset)

	rows, err := ur.dbconnection.Queryx(query, "%"+input.Search+"%")
	if err != nil {
		ur.logger.Error("error retrieving user list: ", err)
		return nil, errors.New("error retrieving user list")
	}
	defer rows.Close()

	for rows.Next() {
		var user dto.UserList
		if err := rows.StructScan(&user); err != nil {
			ur.logger.Error("error scanning user: ", err)
			return nil, errors.New("error retrieving users")
		}
		users = append(users, user)
	}

	var total int
	err = ur.dbconnection.Get(&total, `SELECT COUNT(*) FROM users WHERE name LIKE ?`, "%"+input.Search+"%")
	if err != nil {
		ur.logger.Error("error counting users: ", err)
		return nil, errors.New("error counting users")
	}

	return &dto.UserListOutput{
		Users: users,
		Total: total,
	}, nil
}
