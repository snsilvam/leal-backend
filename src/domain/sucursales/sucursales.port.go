package sucursales

import (
	"context"
)

type SucursalesRepository interface {
	GetAllSucursales(ctx context.Context) ([]*Sucursales, error)
}
