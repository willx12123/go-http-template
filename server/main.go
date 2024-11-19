package main

import (
	"server/internal/dal/db"
	"server/internal/pkg/config"
	"server/internal/pkg/logger"
	"server/internal/pkg/validator"
	"server/internal/srver"
)

func main() {
	config.Init()
	logger.Init()
	validator.Init()

	db.Init()

	srver.Init()
}
