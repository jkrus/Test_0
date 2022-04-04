package services

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"wb_L0/internal/config"
	api "wb_L0/internal/errors"
	"wb_L0/internal/order"
	"wb_L0/pkg/datastore/cache"
)

type (
	Services struct {
		Order order.Service
	}
)

func NewServices(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, orm *gorm.DB, c cache.Cache) (*Services, error) {
	// provide Order Service.
	orderService := order.NewOrderService(ctx, wg, cfg, orm, c)
	if err := orderService.Start(); err != nil {
		return nil, api.ErrStartOrderService(err)
	}

	return &Services{orderService}, nil

}
