package comercioout

import (
	"context"
	"leal-backend/src/adapters/out/database"
	domain "leal-backend/src/domain/comercios/entity"
)

type ComercioPostgresRepository struct {
	database database.Database
}

func ConstructorComercioPostgresRepository(database database.Database) *ComercioPostgresRepository {
	return &ComercioPostgresRepository{
		database: database,
	}
}

func (r *ComercioPostgresRepository) GetAllComercios(ctx context.Context) ([]*domain.Comercio, error) {
	var allComercios []*domain.Comercio

	result := r.database.Db.Find(&allComercios)
	if result.Error != nil {
		return nil, result.Error
	}

	return allComercios, nil
}
