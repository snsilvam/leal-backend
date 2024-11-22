package port

import (
	"context"
	"leal-backend/src/domain/campanas/entity"
)

type CampanaRepository interface {
	CreateCampana(ctx context.Context, campana *entity.Campana) error
	GetAllCampanasOfComercioAndSucursal(ctx context.Context, comercioID int16, sucursalID int16) ([]*entity.Campana, error)
}
