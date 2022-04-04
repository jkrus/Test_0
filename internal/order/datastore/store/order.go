package store

import (
	"gorm.io/gorm"

	"wb_L0/pkg/models"
)

type (
	db struct {
		orm *gorm.DB
	}
)

func NewOrderStorage(orm *gorm.DB) Storage {
	return &db{orm: orm}
}

func (d *db) Create(order *models.OrderDTO) error {
	return d.orm.Create(order).Error
}

func (d *db) GetList(limit int) ([]models.OrderDTO, error) {
	var s []models.OrderDTO
	err := d.orm.Model(models.OrderDTO{}).Limit(limit).Find(&s).Error
	return s, err
}
