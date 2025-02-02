package main

import (
	"react-fsd-template/internal/delivery"
	"react-fsd-template/pkg/config"
	"react-fsd-template/pkg/database"
	"react-fsd-template/pkg/log"
)

func main() {
	logger, infoFile, errorFile := log.InitLogger()
	defer infoFile.Close()
	defer errorFile.Close()

	config.InitConfig()
	logger.Info("config init success")

	db := database.GetDB()
	logger.Info("db init success")

	delivery.Start(db, logger)
}
