package postgres

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
	"github.com/wobenshain/fullstack-developer-challenge/go/types"
	"os"
)

const createItem = `INSERT INTO item (name, picture, description) VALUES ($1, $2, $3)`
const createKVField = `INSERT INTO key_value_field ( item_id, name, value ) VALUES ($1, $2, $3)`
const deleteItem = `DELETE FROM item WHERE id = $1`
const deleteKVFields = `DELETE FROM key_value_field WHERE item_id = $1`
const getItem = "SELECT * FROM item WHERE id = $1"
const getItems = "SELECT * FROM item"
const getKVFields = "SELECT * FROM key_value_field WHERE item_id = $1"
const getAllKVFields = "SELECT * FROM key_value_field"
const updateItem = `
	UPDATE item
	SET
		name = $1,
		picture = $2,
		description = $3
	WHERE id = $4
`

type ItemRepo struct {
	connection string
}

func InitItemRepository() types.ItemRepository {
	host := os.Getenv("pg_host")
	port := os.Getenv("pg_port")
	user := os.Getenv("pg_user")
	password := os.Getenv("pg_password")
	dbname := os.Getenv("pg_dbname")
	i := ItemRepo{
		connection: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			password,
			dbname,
		),
	}
	db := sqlx.MustConnect("postgres", i.connection)
	defer db.Close()

	return &i
}

func (i *ItemRepo) GetAll() ([]models.Item, error) {
	db, err := sqlx.Connect("postgres", i.connection)
	if err != nil {
		return nil, errors.New(502, "Gateway")
	}
	defer db.Close()

	items := make([]Item, 0)
	err = db.Select(&items, getItems)
	if err != nil {
		return nil, errors.New(404, "Not Found")
	}

	kvFields := make([]KeyValueField, 0)
	err = db.Select(&kvFields, getAllKVFields)
	if err != nil {
		return nil, errors.New(404, "Not Found")
	}

	kvMap := map[int][]KeyValueField{}
	for _, kvField := range kvFields {
		kvMap[kvField.ItemID] = append(kvMap[kvField.ItemID], kvField)
	}

	returnItems := make([]models.Item, 0)
	for _, item := range items {
		returnItems = append(returnItems, convertToModel(item, kvMap[item.ID]))
	}

	return returnItems, nil
}

func (i *ItemRepo) Get(id int) (models.Item, error) {
	empty := models.Item{}
	db, err := sqlx.Connect("postgres", i.connection)
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}
	defer db.Close()

	item := make([]Item, 0)
	err = db.Select(&item, getItem, id)
	if err != nil || len(item) != 1 {
		return empty, errors.New(404, "Not Found")
	}

	kvFields := make([]KeyValueField, 0)
	err = db.Select(&kvFields, getKVFields, id)
	if err != nil {
		return empty, errors.New(404, "Not Found")
	}

	return convertToModel(item[0], kvFields), nil
}

func (i *ItemRepo) Create(item models.Item) (models.Item, error) {
	empty := models.Item{}
	db, err := sqlx.Connect("postgres", i.connection)
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}
	defer db.Close()

	tx := db.MustBegin()
	_ = tx.MustExec(createItem, item.Name, item.Picture, item.Description)
	row := tx.QueryRow("SELECT lastval()")
	err = row.Scan(&item.ID)
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}
	kvs, err := generateKeyValuePairs(tx, item)
	if err != nil {
		return empty, err
	}
	err = tx.Commit()
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}

	item.KeyValueFields = kvs
	return item, nil
}

func (i *ItemRepo) Update(item models.Item) (models.Item, error) {
	empty := models.Item{}
	db, err := sqlx.Connect("postgres", i.connection)
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}
	defer db.Close()

	tx := db.MustBegin()
	_ = tx.MustExec(updateItem, item.Name, item.Picture, item.Description, item.ID)
	_ = tx.MustExec(deleteKVFields, item.ID)
	kvs, err := generateKeyValuePairs(tx, item)
	if err != nil {
		return empty, err
	}
	err = tx.Commit()
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}

	item.KeyValueFields = kvs
	return item, nil
}

func (i *ItemRepo) Delete(id int) (models.Item, error) {
	empty := models.Item{}
	item, err := i.Get(id)
	if err != nil {
		return empty, errors.New(404, "Not Found")
	}

	db, err := sqlx.Connect("postgres", i.connection)
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}
	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec(deleteKVFields, item.ID)
	tx.MustExec(deleteItem, item.ID)
	err = tx.Commit()
	if err != nil {
		return empty, errors.New(502, "Gateway")
	}

	return item, nil
}

func generateKeyValuePairs(tx *sqlx.Tx, item models.Item) ([]*models.KeyValue, error) {
	kvs := make([]*models.KeyValue, 0)
	for _, kv := range item.KeyValueFields {
		if kv.Name == nil {
			return nil, errors.New(400, "Bad Request")
		}
		kvField := models.KeyValue{
			Name:   kv.Name,
			Value:  kv.Value,
		}
		tx.MustExec(createKVField, item.ID, kv.Name, kv.Value)
		row := tx.QueryRow("SELECT lastval()")
		err := row.Scan(&kvField.ID)
		if err != nil {
			return nil, errors.New(502, "Gateway")
		}
		kvs = append(kvs, &kvField)
	}
	return kvs, nil
}

func convertToModel(i Item, kvFields []KeyValueField) models.Item {
	kvs := make([]*models.KeyValue, 0)
	for i := range kvFields {
		kvs = append(kvs, &models.KeyValue{
			ID:    kvFields[i].ID,
			Name:  &kvFields[i].Name,
			Value: kvFields[i].Value,
		})
	}
	return models.Item{
		ID:             int64(i.ID),
		Name:           &i.Name,
		Picture:        &i.Picture,
		Description:    &i.Description,
		KeyValueFields: kvs,
	}
}
