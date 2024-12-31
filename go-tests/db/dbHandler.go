package db

import (
	"fmt"

	"go-tests/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandler struct {
}

func MigrateDb() error {
	conn, err := InitConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer CloseConnection(conn)

	return err
}

func InitConnection() (*gorm.DB, error) {
	appConfig := config.GetConfig()
	return gorm.Open(postgres.Open(appConfig.DbConnectionString), &gorm.Config{})
}

func CloseConnection(conn *gorm.DB) error {
	dbConn, err := conn.DB()
	if err != nil {
		fmt.Println("CloseConnection (db package) error - " + err.Error())
	}
	defer dbConn.Close()

	return err
}
