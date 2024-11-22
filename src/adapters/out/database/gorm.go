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

	err = db.AutoMigrate(&models.Comercio{}, &models.TipoBeneficios{}, &models.Sucursales{}, &models.Campana{}, &models.Usuario{}, &models.PuntosLeal{}, &models.Compras{}, &models.Ventas{})
	if err != nil {
		return nil, err
	}

	return &Database{
		Db: db,
	}, nil
}
