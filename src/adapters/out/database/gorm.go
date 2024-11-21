package database

import (
	"leal-backend/src/adapters/out/database/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func ConstructorDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Comercio{})
	if err != nil {
		return nil, err
	}

	return &Database{
		Db: db,
	}, nil
}
