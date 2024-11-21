package campanaPostgresql

import (
	"context"
	"fmt"
	"leal-backend/src/adapters/out/database"
	"leal-backend/src/domain/campanas/entity"
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
	fmt.Println("capa out - Creando la siguiente campaña. . . ", campana)
	return nil
}
