package order

import (
	"context"
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"

	"wb_L0/internal/config"
	"wb_L0/internal/order/datastore/store"
	"wb_L0/pkg/datastore/cache"
	"wb_L0/pkg/models"
	"wb_L0/pkg/service"
)

type (
	Service interface {
		service.Service
		Get(uid string) (*models.Order, error)
	}

	orderService struct {
		ctx context.Context
		wg  *sync.WaitGroup
		cfg *config.Config

		repo  store.Storage
		cache cache.Cache
	}
)

func NewOrderService(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, orm *gorm.DB, c cache.Cache) Service {
	return &orderService{
		ctx:   ctx,
		wg:    wg,
		cfg:   cfg,
		repo:  store.NewOrderStorage(orm),
		cache: c,
	}
}

func (ns *orderService) Get(uid string) (*models.Order, error) {
	dto := &models.OrderDTO{}
	dto.Data = ns.cache.Get(nil, []byte(uid))
	if dto.Data == nil {
		return nil, ErrOrderNotFound(fmt.Errorf("%v", uid))
	}

	ord, err := dto.Decode()
	if err != nil {
		return nil, ErrOrderDecode(err)
	}

	return ord, nil
}

func (ns *orderService) Start() error {
	log.Println("Start Order service...")

	if err := ns.recoveryFromDB(ns.cfg); err != nil {
		return err
	}

	ns.createHandlerContext()

	log.Println("Order service start success.")

	return nil
}

func (ns *orderService) Stop() error {
	log.Println("Stop Order Service...")

	log.Println("Order service stopped.")

	return nil
}

func (ns *orderService) createHandlerContext() {
	ns.wg.Add(1)
	go func() {
		for {
			<-ns.ctx.Done()
			_ = ns.Stop()
			ns.wg.Done()
			return
		}
	}()

}

func (ns *orderService) recoveryFromDB(cfg *config.Config) error {
	size := cfg.Cache.RecoverySize
	list, err := ns.repo.GetList(int(size))
	if err != nil {
		return err
	}
	for _, order := range list {
		ns.cache.Set([]byte(order.Uid), order.Data)
	}
	return nil
}
