package port

import (
	"context"
	domain "leal-backend/src/domain/comercios/entity"

	"github.com/stretchr/testify/mock"
)

type MockComercioRepository struct {
	mock.Mock
}

func (m *MockComercioRepository) GetAllComercios(ctx context.Context) ([]*domain.Comercio, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Comercio), args.Error(1)
}
