package database

import (
	"backend/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetDB() *sqlx.DB {
	connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		viper.GetString(config.DBUser),
		viper.GetString(config.DBPassword),
		viper.GetString(config.DBHost),
		viper.GetInt(config.DBPort),
		viper.GetString(config.DBName))
	//connString := fmt.Sprintf("user=%v host=%v port=%v dbname=%v password=%v sslmode=disable",
	//	viper.GetString(config.DBUser),
	//	viper.GetString(config.DBHost),
	//	viper.GetInt(config.DBPort),
	//	viper.GetString(config.DBName),
	//	viper.GetString(config.DBPassword),
	//)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to DB. Error: %v", err.Error()))
	}

	return db
}
