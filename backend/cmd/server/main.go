package main

import (
	"fmt"
	"log"

	"xiang/backend/internal/config"
	"xiang/backend/internal/model"
	"xiang/backend/internal/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.BuildDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect database failed: %v", err)
	}

	if err := model.AutoMigrate(db); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	r := router.New(db, cfg)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	if err := r.Run(addr); err != nil {
		log.Fatalf("run server failed: %v", err)
	}
}
