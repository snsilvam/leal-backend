package service

import (
	"context"
	"leal-backend/src/domain/campanas/entity"
	"leal-backend/src/domain/campanas/port"
)

type CampanaService struct {
	campanaRepository port.CampanaRepository
}

func ConstructorCampanaService(repository port.CampanaRepository) *CampanaService {
	return &CampanaService{
		campanaRepository: repository,
	}
}

func (s *CampanaService) CreateCampana(ctx context.Context, campana *entity.Campana) error {
	return s.campanaRepository.CreateCampana(ctx, campana)
}

/* func (s *CampanaService) GetAllCampanas(comercioID int16, sucursalID int16) ([]*entity.Campana, error) {
	return s.campanaRepository.GetAllCampanas(comercioID, sucursalID)
} */
