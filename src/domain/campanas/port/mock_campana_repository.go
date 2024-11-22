package port

import (
	"context"
	"leal-backend/src/domain/campanas/entity"

	"github.com/stretchr/testify/mock"
)

type MockCampanaRepository struct {
	mock.Mock
}

func (m *MockCampanaRepository) CreateCampana(ctx context.Context, campana *entity.Campana) error {
	args := m.Called(ctx, campana)
	return args.Error(0)
}

func (m *MockCampanaRepository) GetAllCampanasOfComercioAndSucursal(ctx context.Context, comercioID int16, sucursalID int16) ([]*entity.Campana, error) {
	args := m.Called(ctx, comercioID, sucursalID)
	return args.Get(0).([]*entity.Campana), args.Error(1)
}
