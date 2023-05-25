package services

import (
	"labora-api/models"
)

type DBHandler interface {
	CreateItem(item models.Item) error
	GetItem(id int) (models.Item, error)
	UpdateItem(item models.Item) error
	DeleteItem(id int) error
}
