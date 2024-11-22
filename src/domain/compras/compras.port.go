package compras

import (
	"context"
)

type ComprasRepository interface {
	CreateCompra(context context.Context, compra *Compras) error
	CampanaActiva(ctx context.Context, idCampana int16) (int16, error)
	DeterminarQueTipoDeBeneficioEstaActivo(context context.Context, beneficioID int16, compra *Compras) error
	AddPuntosTipoBeneficio1(context context.Context, compra *Compras) error
	AddPuntosTipoBeneficio2(context context.Context, compra *Compras) error
}
