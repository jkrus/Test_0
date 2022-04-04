package postgres

import (
	"context"
	"log"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"wb_L0/internal/config"
)

func Start(ctx context.Context, cfg *config.Config, wg *sync.WaitGroup) (*gorm.DB, error) {
	log.Println("Start connect to database...")
	orm, err := postgresConnect(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "database connect filed")
	}

	createContextHandler(ctx, orm, wg)
	autoMigrate(orm)

	log.Println("Connect to database success.")
	return orm, nil
}

func Close(orm *gorm.DB) error {
	if orm != nil {
		db, err := orm.DB()
		if err != nil {
			return err
		}

		return db.Close()
	}

	return nil
}

func createContextHandler(ctx context.Context, db *gorm.DB, wg *sync.WaitGroup) {
	cc, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go func() {
		for {
			<-cc.Done()
			log.Println("Stop database service...")
			if err := Close(db); err != nil {
				log.Println("Close database connection: ", err)
			} else {
				log.Println("Database Service stopped.")
			}
			cancel()
			wg.Done()
			return
		}
	}()

}
