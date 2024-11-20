package port

import "leal-backend/src/domain/campanas/entity"

type CampanaRepository interface {
	CreateCampana(campana *entity.Campana) error
	GetAllCampanas(comercioID int16, sucursalID int16) ([]*entity.Campana, error)
}
