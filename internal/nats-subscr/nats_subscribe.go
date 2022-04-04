package nats_subscr

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"

	json "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"wb_L0/internal/config"
	"wb_L0/internal/order/datastore/store"
	"wb_L0/pkg/datastore/cache"
	"wb_L0/pkg/models"
	"wb_L0/pkg/service"
)

type (
	NatsSubscribe interface {
		service.Service
	}

	natsSubscribeService struct {
		ctx context.Context
		wg  *sync.WaitGroup

		natsClient *nats.Conn
		connOpt    *nats.Options
		subs       sub

		repo  store.Storage
		cache cache.Cache
	}

	sub struct {
		order *nats.Subscription
	}
)

func NewNatsSubscribeService(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, orm *gorm.DB, cache cache.Cache) NatsSubscribe {
	return &natsSubscribeService{
		ctx: ctx,
		wg:  wg,
		connOpt: &nats.Options{
			ReconnectWait:  cfg.NatsSubscribeOptions.ReconnectWait,
			MaxReconnect:   cfg.NatsSubscribeOptions.MaxReconnect,
			AllowReconnect: cfg.NatsSubscribeOptions.AllowReconnect,
			Timeout:        cfg.NatsSubscribeOptions.Timeout,
		},

		cache: cache,
		repo:  store.NewOrderStorage(orm),
	}
}

func (ns *natsSubscribeService) Start() error {
	log.Println("Start NatsSubscribe service...")

	// Connect to a server
	var err error
	ns.natsClient, err = ns.connOpt.Connect()
	if err != nil {
		return errors.Wrap(err, "failed connect to nats server")
	}

	_, err = ns.natsClient.Subscribe("ORDERS", ns.orderHandler)
	if err != nil {
		return err
	}

	ns.createHandlerContext()

	log.Println("NatsSubscribe service start success.")

	return nil
}

func (ns *natsSubscribeService) Stop() error {
	log.Println("Stop NatsSubscribe Service...")
	ns.subs.order.Unsubscribe()
	ns.natsClient.Close()
	log.Println("NatsSubscribe service stopped.")

	return nil
}

func (ns *natsSubscribeService) createHandlerContext() {
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

func (ns *natsSubscribeService) orderHandler(msg *nats.Msg) {
	// TODO handle error
	_ = msg.InProgress()

	ord, err := ns.natsMsgDecode(msg)
	if err != nil {
		log.Println("MSG DECODE", err)
	}
	orderDTO := &models.OrderDTO{Uid: ord.OrderUid, Data: msg.Data}
	if err = ns.repo.Create(orderDTO); err != nil {
		// TODO handle error
		_ = msg.Nak()
	}

	ns.cache.Set([]byte(orderDTO.Uid), msg.Data)

	// TODO handle error
	_ = msg.Ack()
}

func (ns *natsSubscribeService) natsMsgDecode(msg *nats.Msg) (*models.Order, error) {
	order := &models.Order{}

	if err := json.Unmarshal(msg.Data, order); err != nil {
		log.Println(err)
		return nil, err
	}

	if !validate(order) {
		return nil, ErrOrderValidate(fmt.Errorf("order data is not valid"))
	}

	return order, nil

}

func validate(order *models.Order) bool {
	if order.OrderUid == "" {
		return false
	}

	return true
}
