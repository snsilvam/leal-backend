package sucursalesdatabase

import (
	"context"
	"leal-backend/src/adapters/out/database"
	"leal-backend/src/domain/sucursales"
)

type SucursalesPostgresRepository struct {
	database database.Database
}

func ConstructorSucursalesPostgresRepository(database database.Database) *SucursalesPostgresRepository {
	return &SucursalesPostgresRepository{
		database: database,
	}
}

func (r *SucursalesPostgresRepository) GetAllSucursales(context context.Context, comercioId int16) ([]*sucursales.Sucursales, error) {
	var allSucursales []*sucursales.Sucursales

	if err := r.database.Db.Where("comercio_id = ?", comercioId).Find(&allSucursales).Error; err != nil {
		return nil, err
	}

	return allSucursales, nil
}
