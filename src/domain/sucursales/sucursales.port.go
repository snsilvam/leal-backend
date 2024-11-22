package sucursales

import (
	"context"
)

type SucursalesRepository interface {
	GetAllSucursales(ctx context.Context, comercioId int16) ([]*Sucursales, error)
}
