package postgres

import (
	"log"

	"gorm.io/gorm"

	"wb_L0/pkg/models"
)

func autoMigrate(orm *gorm.DB) {
	log.Println("Auto-migration start...")

	// Try auto-migrate database tables with specified entities.
	if err := orm.AutoMigrate(models.OrderDTO{}); err != nil {
		log.Fatal("automigrate failed:", err)
	}

	log.Println("Auto-migration complete.")
}
