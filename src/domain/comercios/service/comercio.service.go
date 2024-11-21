package service

import (
	"context"
	domain "leal-backend/src/domain/comercios/entity"
	"leal-backend/src/domain/comercios/port"
)

type ComercioService struct {
	comercioRepository port.ComercioRepository
}

func ConstructorComercioService(repository port.ComercioRepository) *ComercioService {
	return &ComercioService{
		comercioRepository: repository,
	}
}

func (s *ComercioService) GetAllComercios(ctx context.Context) ([]*domain.Comercio, error) {
	return s.comercioRepository.GetAllComercios(ctx)
}
