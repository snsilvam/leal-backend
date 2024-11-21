package port

import (
	"context"
	domain "leal-backend/src/domain/comercios/entity"
)

type ComercioRepository interface {
	GetAllComercios(ctx context.Context) ([]*domain.Comercio, error)
}
