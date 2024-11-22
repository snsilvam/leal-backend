package campanaPostgresql

import (
	"context"
	"leal-backend/src/adapters/out/database"
	"leal-backend/src/domain/campanas/entity"
	"log"
)

type CampanasPostgresRepository struct {
	database database.Database
}

func ConstructorCampanasPostgresRepository(database database.Database) *CampanasPostgresRepository {
	return &CampanasPostgresRepository{
		database: database,
	}
}

func (r *CampanasPostgresRepository) CreateCampana(ctx context.Context, campana *entity.Campana) error {
	result := r.database.Db.Create(&campana)
	if result.Error != nil {
		log.Fatal("no se pudo insertar el registro de la campaña:", result.Error)
		return result.Error
	}

	return nil
}
