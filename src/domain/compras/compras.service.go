package compras

import (
	"context"
)

type CompraService struct {
	compraRepository ComprasRepository
}

func ConstructorCompraService(repository ComprasRepository) *CompraService {
	return &CompraService{
		compraRepository: repository,
	}
}

func (s *CompraService) CreateCompra(context context.Context, compra *Compras) error {
	return s.compraRepository.CreateCompra(context, compra)
}

func (s *CompraService) AddPuntosTipoBeneficio1(context context.Context, compra *Compras) error {
	return s.compraRepository.AddPuntosTipoBeneficio1(context, compra)
}

func (s *CompraService) AddPuntosTipoBeneficio2(context context.Context, compra *Compras) error {
	return s.compraRepository.AddPuntosTipoBeneficio2(context, compra)
}

func (s *CompraService) CampanaActiva(ctx context.Context, idCampana int16) (int16, error) {
	return s.compraRepository.CampanaActiva(ctx, idCampana)
}

func (s *CompraService) DeterminarQueTipoDeBeneficioEstaActivo(context context.Context, beneficioID int16, compra *Compras) error {
	return s.compraRepository.DeterminarQueTipoDeBeneficioEstaActivo(context, beneficioID, compra)
}
