package sucursales

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockSucursalesRepository struct {
	mock.Mock
}

func (m *MockSucursalesRepository) GetAllSucursales(ctx context.Context, comercioId int16) ([]*Sucursales, error) {
	args := m.Called(ctx, comercioId)
	return args.Get(0).([]*Sucursales), args.Error(1)
}
