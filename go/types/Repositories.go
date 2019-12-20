package types

import "github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"

type ItemRepository interface {
	GetAll() ([]models.Item, error)
	Get(id int) (models.Item, error)
	Create(models.Item) (models.Item, error)
	Update(models.Item) (models.Item, error)
	Delete(id int) (models.Item, error)
}

type Repositories struct {
	ItemRepository
}
