package store

import (
	"wb_L0/pkg/models"
)

type (
	Storage interface {
		Create(note *models.OrderDTO) error
		GetList(limit int) ([]models.OrderDTO, error)
	}
)
