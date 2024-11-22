package beneficiosdatabase

import (
	"context"
	"leal-backend/src/adapters/out/database"
	beneficios "leal-backend/src/domain/tipo-beneficios"
)

type BeneficiosPostgresRepository struct {
	database database.Database
}

func ConstructorBeneficiosPostgresRepository(database database.Database) *BeneficiosPostgresRepository {
	return &BeneficiosPostgresRepository{
		database: database,
	}
}

func (r *BeneficiosPostgresRepository) GetAllBeneficios(ctx context.Context) ([]*beneficios.TipoBeneficios, error) {
	var allBeneficios []*beneficios.TipoBeneficios

	result := r.database.Db.Find(&allBeneficios)
	if result.Error != nil {
		return nil, result.Error
	}

	return allBeneficios, nil
}
