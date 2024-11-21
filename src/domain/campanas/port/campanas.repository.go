package port

import (
	"context"
	"leal-backend/src/domain/campanas/entity"
)

type CampanaRepository interface {
	CreateCampana(ctx context.Context, campana *entity.Campana) error
	/* GetAllCampanas(comercioID int16, sucursalID int16) ([]*entity.Campana, error) */
}
