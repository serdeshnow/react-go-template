package main

import (
	"react-go-template/internal/delivery"
	"react-go-template/pkg/config"
	"react-go-template/pkg/database"
	"react-go-template/pkg/log"
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
