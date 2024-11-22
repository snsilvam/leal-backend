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

func (s *SucursalesService) GetAllSucursales(context context.Context) ([]*Sucursales, error) {
	return s.repositorySucursales.GetAllSucursales(context)
}
