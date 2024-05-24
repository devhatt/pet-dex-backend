package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"

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
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error on saving ong")
		return err
	}

	return nil
}

func (or *OngRepository) List(limit, offset int, sortBy, order string) (ongs []*dto.OngListDto, err error) {
	query := fmt.Sprintf(`
	SELECT 
    legal_persons.id, 
    legal_persons.userId, 
    legal_persons.phone, 
    legal_persons.openingHours,
		legal_persons.links,
    users.name,
		addresses.address,
		addresses.city,
		addresses.state
	FROM 
    legal_persons
	INNER JOIN 
    users ON legal_persons.userId = users.id
	INNER JOIN
		addresses ON legal_persons.userId = addresses.userId
	ORDER BY 
    %s %s
	LIMIT ? OFFSET ?`, sortBy, order)
	rows, err := or.dbconnection.Queryx(query, limit, offset)
	if err != nil {
		logger.Error("error listing ongs", err)
		return nil, fmt.Errorf("error listing ongs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ong dto.OngListDto
		err := rows.StructScan(&ong)
		if err != nil {
			logger.Error("error scanning ongs", err)
			return nil, fmt.Errorf("error scanning ongs: %w", err)
		}
		ongs = append(ongs, &ong)
	}

	return ongs, nil
}
  
func (or *OngRepository) Update(id uniqueEntityId.ID, ongToUpdate entity.Ong) error {

	query := "UPDATE legal_persons SET"
	var values []interface{}

	if ongToUpdate.Phone != "" {
		query = query + " phone =?"
		values = append(values, ongToUpdate.Phone)
	}

	if ongToUpdate.OpeningHours != "" {
		query = query + " openingHours =?"
		values = append(values, ongToUpdate.OpeningHours)
	}

	if ongToUpdate.AdoptionPolicy != "" {
		query = query + " adoptionPolicy =?"
		values = append(values, ongToUpdate.AdoptionPolicy)
	}

	if string(*ongToUpdate.Links) != "" {
		query = query + " links =?"
		values = append(values, ongToUpdate.Links)
	}

	query = query + " updated_at =?,"
	values = append(values, time.Now())

	n := len(query)
	query = query[:n-1] + " WHERE id =?"
	values = append(values, id)

	fmt.Printf("Query to update: %s", query)

	_, err := or.dbconnection.Exec(query, values...)

	if err != nil {
		logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("error on updating ong")
		return err
	}

	return nil
}

func (or *OngRepository) FindById(id uniqueEntityId.ID) (*entity.Ong, error) {
	return nil, nil

}
