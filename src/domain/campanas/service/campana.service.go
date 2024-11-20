package service

import (
	"leal-backend/src/domain/campanas/entity"
	"leal-backend/src/domain/campanas/port"
)

type CampanaService struct {
	campanaRepo port.CampanaRepository
}

func (s *CampanaService) CreateCampana(campana *entity.Campana) error {
	return s.campanaRepo.CreateCampana(campana)
}

func (s *CampanaService) GetAllCampanas(comercioID int16, sucursalID int16) ([]*entity.Campana, error) {
	return s.campanaRepo.GetAllCampanas(comercioID, sucursalID)
}
