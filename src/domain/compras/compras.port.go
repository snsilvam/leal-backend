package compras

import (
	"context"
)

type ComprasRepository interface {
	CreateCompra(context context.Context, compra *Compras) error
	AddPuntosTipoBeneficio1(context context.Context, compra *Compras) error
	AddPuntosTipoBeneficio2(context context.Context, compra *Compras) error
	ComercioConCampanaActiva(context context.Context, compra *Compras) (bool error)
}
