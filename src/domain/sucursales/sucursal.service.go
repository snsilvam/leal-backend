package sucursales

import "context"

type SucursalesService struct {
	repositorySucursales SucursalesRepository
}

func ConstructorSucursalesService(repository SucursalesRepository) *SucursalesService {
	return &SucursalesService{
		repositorySucursales: repository,
	}
}

func (s *SucursalesService) GetAllSucursales(context context.Context, comercioId int16) ([]*Sucursales, error) {
	return s.repositorySucursales.GetAllSucursales(context, comercioId)
}
